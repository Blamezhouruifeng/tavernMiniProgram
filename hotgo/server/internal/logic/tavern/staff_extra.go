package tavern

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"

	"hotgo/internal/dao"
	"hotgo/internal/model/input/tavernin"
)

// DepositRegister 代存登记（自带酒水）：实物入库 + 生成瓶签 + 占用储位
func (s *sTavernStaff) DepositRegister(ctx context.Context, in *tavernin.DepositRegisterInp) (res *tavernin.BottleModel, err error) {
	operatorId, err := s.requireStaff(ctx)
	if err != nil {
		return nil, err
	}

	// 定位会员：优先 memberId，其次手机号
	memberCols := dao.TavernMember.Columns()
	var memberRow gdb.Record
	if in.MemberId > 0 {
		memberRow, err = dao.TavernMember.Ctx(ctx).WherePri(in.MemberId).One()
	} else if in.Phone != "" {
		memberRow, err = dao.TavernMember.Ctx(ctx).Where(memberCols.Phone, in.Phone).One()
	} else {
		return nil, gerror.New("请提供会员ID或手机号")
	}
	if err != nil {
		return nil, err
	}
	if memberRow.IsEmpty() {
		return nil, gerror.New("未找到该会员，请先让客人扫码登录并绑定手机号")
	}

	memberId := memberRow[memberCols.Id].Int64()
	nickname := memberRow[memberCols.Nickname].String()
	phone := memberRow[memberCols.Phone].String()
	phoneTail := ""
	if len(phone) >= 4 {
		phoneTail = phone[len(phone)-4:]
	}

	// 储位：指定则校验，未指定则取第一个空闲
	slotCols := dao.TavernStorageSlot.Columns()
	slotCode := in.SlotCode
	var slotId int64
	if slotCode != "" {
		slot, e := dao.TavernStorageSlot.Ctx(ctx).Where(slotCols.Code, slotCode).One()
		if e != nil {
			return nil, e
		}
		if slot.IsEmpty() {
			return nil, gerror.New("储位不存在")
		}
		slotId = slot[slotCols.Id].Int64()
	} else {
		slot, e := dao.TavernStorageSlot.Ctx(ctx).
			Where(slotCols.Status, 1).
			OrderAsc(slotCols.Id).
			One()
		if e == nil && !slot.IsEmpty() {
			slotId = slot[slotCols.Id].Int64()
			slotCode = slot[slotCols.Code].String()
		} else {
			slotCode = "待分配"
		}
	}

	remain := in.RemainRatio
	if remain <= 0 {
		remain = 100
	}

	cols := dao.TavernBottleDeposit.Columns()
	id, err := dao.TavernBottleDeposit.Ctx(ctx).Data(g.Map{
		cols.StoreId:     1,
		cols.DepositNo:   "DY" + gtime.Now().Format("YmdHis"),
		cols.LabelCode:   "LK" + gconv.String(gtime.TimestampMilli()),
		cols.MemberId:    memberId,
		cols.MemberName:  nickname,
		cols.PhoneTail:   phoneTail,
		cols.SourceType:  1,
		cols.GoodsName:   in.GoodsName,
		cols.Spec:        in.Spec,
		cols.Degree:      in.Degree,
		cols.TotalRatio:  remain,
		cols.RemainRatio: remain,
		cols.SlotId:      slotId,
		cols.SlotCode:    slotCode,
		cols.Photo:       in.Photo,
		cols.Status:      1,
		cols.DepositAt:   gtime.Now(),
		cols.ExpireAt:    gtime.New(gtime.Now().Unix() + int64(180*24*3600)),
		cols.Remark:      in.Remark,
		cols.CreatedBy:   operatorId,
		cols.CreatedAt:   gtime.Now(),
		cols.UpdatedAt:   gtime.Now(),
	}).InsertAndGetId()
	if err != nil {
		return nil, gerror.Wrap(err, "代存登记失败")
	}

	if slotId > 0 {
		_, _ = dao.TavernStorageSlot.Ctx(ctx).WherePri(slotId).Data(g.Map{
			slotCols.IsUsed:    1,
			slotCols.UpdatedAt: gtime.Now(),
		}).Update()
	}

	row, err := dao.TavernBottleDeposit.Ctx(ctx).WherePri(id).One()
	if err != nil {
		return nil, err
	}
	return toBottleModel(row), nil
}

// MoveTable 转台/并台：会话与订单归属迁移到新桌
func (s *sTavernStaff) MoveTable(ctx context.Context, in *tavernin.MoveTableInp) (err error) {
	if _, err = s.requireStaff(ctx); err != nil {
		return err
	}

	sessionCols := dao.TavernTableSession.Columns()
	session, err := getSessionById(ctx, in.SessionId)
	if err != nil {
		return err
	}
	if session.IsEmpty() {
		return gerror.New("桌台会话不存在")
	}
	if session[sessionCols.Status].Int() != 1 {
		return gerror.New("该会话已结束，无法转台")
	}

	table, err := getTableByName(ctx, in.ToTableNo)
	if err != nil {
		return err
	}
	if table.IsEmpty() {
		return gerror.New("目标桌号不存在")
	}
	toTableId := table[dao.TavernTable.Columns().Id].Int64()
	fromTableId := session[sessionCols.TableId].Int64()
	if fromTableId == toTableId {
		return gerror.New("目标桌与当前桌相同")
	}

	// 目标桌若有进行中会话，视为并台：将目标会话关闭并提示
	target, _ := getOpenSession(ctx, toTableId)
	if !target.IsEmpty() {
		return gerror.New("目标桌正在使用中，请先清台或选择其他桌位")
	}

	if _, err = dao.TavernTableSession.Ctx(ctx).WherePri(in.SessionId).Data(g.Map{
		sessionCols.TableId:   toTableId,
		sessionCols.UpdatedAt: gtime.Now(),
	}).Update(); err != nil {
		return gerror.Wrap(err, "转台失败")
	}

	// 订单快照同步桌号
	orderCols := dao.TavernOrder.Columns()
	_, _ = dao.TavernOrder.Ctx(ctx).
		Where(orderCols.SessionId, in.SessionId).
		Data(g.Map{
			orderCols.TableId:   toTableId,
			orderCols.TableName: in.ToTableNo,
			orderCols.UpdatedAt: gtime.Now(),
		}).Update()

	tableCols := dao.TavernTable.Columns()
	_, _ = dao.TavernTable.Ctx(ctx).WherePri(fromTableId).Data(g.Map{
		tableCols.SeatStatus: 1,
		tableCols.UpdatedAt:  gtime.Now(),
	}).Update()
	_, _ = dao.TavernTable.Ctx(ctx).WherePri(toTableId).Data(g.Map{
		tableCols.SeatStatus: 2,
		tableCols.UpdatedAt:  gtime.Now(),
	}).Update()

	g.Log().Infof(ctx, "tavern move table sessionId:%d from:%d to:%d", in.SessionId, fromTableId, toTableId)
	return
}
