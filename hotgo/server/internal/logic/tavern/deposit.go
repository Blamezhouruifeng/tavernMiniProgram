package tavern

import (
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"

	"hotgo/internal/dao"
	"hotgo/internal/library/cache"
	"hotgo/internal/model/input/tavernin"
	"hotgo/internal/service"
)

type sTavernDeposit struct{}

func init() {
	service.RegisterTavernDeposit(NewTavernDeposit())
}

func NewTavernDeposit() *sTavernDeposit {
	return &sTavernDeposit{}
}

// takeCodeCacheKey 取酒码缓存键
func takeCodeCacheKey(code string) string {
	return fmt.Sprintf("tavern:takecode:%s", code)
}

// List 我的存酒
func (s *sTavernDeposit) List(ctx context.Context, in *tavernin.BottleListInp) (res []*tavernin.BottleModel, err error) {
	memberId := service.TavernMember().CurrentId(ctx)
	if memberId == 0 {
		return nil, gerror.New("请先登录")
	}

	cols := dao.TavernBottleDeposit.Columns()
	m := dao.TavernBottleDeposit.Ctx(ctx)
	if !in.All {
		m = m.Where(cols.MemberId, memberId)
	}
	rows, err := m.OrderDesc(cols.Id).Limit(100).All()
	if err != nil {
		return nil, err
	}

	res = make([]*tavernin.BottleModel, 0, len(rows))
	for _, row := range rows {
		res = append(res, toBottleModel(row))
	}
	return
}

// toBottleModel 存酒记录 → 出参模型
func toBottleModel(row gdb.Record) *tavernin.BottleModel {
	cols := dao.TavernBottleDeposit.Columns()
	return &tavernin.BottleModel{
		Id:          row[cols.Id].Int64(),
		DepositNo:   row[cols.DepositNo].String(),
		LabelCode:   row[cols.LabelCode].String(),
		MemberId:    row[cols.MemberId].Int64(),
		MemberName:  row[cols.MemberName].String(),
		PhoneTail:   row[cols.PhoneTail].String(),
		SourceType:  row[cols.SourceType].Int(),
		GoodsName:   row[cols.GoodsName].String(),
		Spec:        row[cols.Spec].String(),
		Degree:      row[cols.Degree].String(),
		RemainRatio: row[cols.RemainRatio].Float64(),
		SlotCode:    row[cols.SlotCode].String(),
		Photo:       row[cols.Photo].String(),
		Status:      row[cols.Status].Int(),
		DepositAt:   row[cols.DepositAt].GTime(),
		ExpireAt:    row[cols.ExpireAt].GTime(),
	}
}

// TakeCode 生成动态取酒码（30 秒有效，防截图冒领）
func (s *sTavernDeposit) TakeCode(ctx context.Context, in *tavernin.TakeCodeInp) (res *tavernin.TakeCodeModel, err error) {
	memberId := service.TavernMember().CurrentId(ctx)
	if memberId == 0 {
		return nil, gerror.New("请先登录")
	}

	code := gconv.String(gtime.TimestampNano() % 1000000)
	if len(code) < 6 {
		code = gconv.String(100000 + gtime.TimestampNano()%900000)
	}

	payload := g.Map{
		"memberId":   memberId,
		"depositIds": in.DepositIds,
	}
	if err = cache.Instance().Set(ctx, takeCodeCacheKey(code), payload, 30*time.Second); err != nil {
		return nil, gerror.Wrap(err, "生成取酒码失败")
	}

	res = &tavernin.TakeCodeModel{
		Code:          code,
		ExpireSeconds: 30,
		DepositIds:    in.DepositIds,
	}
	return
}

// getTakeCodePayload 读取取酒码（供员工核销使用，包内共享）
func getTakeCodePayload(ctx context.Context, code string) (memberId int64, depositIds []int64, err error) {
	val, err := cache.Instance().Get(ctx, takeCodeCacheKey(code))
	if err != nil {
		return
	}
	if val.IsNil() {
		return 0, nil, gerror.New("取酒码无效或已过期")
	}
	m := val.Map()
	memberId = gconv.Int64(m["memberId"])
	if ids, ok := m["depositIds"]; ok {
		depositIds = gconv.Int64s(ids)
	}
	return
}

// DepositFromOrder 退酒入存：支付成功订单中的整瓶酒可直接转为存酒
func (s *sTavernDeposit) DepositFromOrder(ctx context.Context, in *tavernin.DepositFromOrderInp) (res *tavernin.BottleModel, err error) {
	memberId := service.TavernMember().CurrentId(ctx)
	if memberId == 0 {
		return nil, gerror.New("请先登录")
	}

	orderCols := dao.TavernOrder.Columns()
	order, err := dao.TavernOrder.Ctx(ctx).WherePri(in.OrderId).One()
	if err != nil {
		return nil, err
	}
	if order.IsEmpty() {
		return nil, gerror.New("订单不存在")
	}
	if order[orderCols.PayStatus].Int() != 1 {
		return nil, gerror.New("仅支付成功的订单可转存")
	}
	if order[orderCols.MemberId].Int64() != memberId {
		return nil, gerror.New("只能寄存自己下单的酒品")
	}

	itemCols := dao.TavernOrderItem.Columns()
	item, err := dao.TavernOrderItem.Ctx(ctx).
		Where(itemCols.OrderId, in.OrderId).
		WhereIn(itemCols.Id, in.ItemIds).
		One()
	if err != nil {
		return nil, err
	}
	if item.IsEmpty() {
		return nil, gerror.New("请选择要寄存的酒品")
	}
	if item[itemCols.IsWine].Int() != 1 {
		return nil, gerror.New("只有酒水可以寄存")
	}

	profile, err := service.TavernMember().Profile(ctx)
	if err != nil {
		return nil, err
	}
	phoneTail := ""
	if len(profile.Phone) >= 4 {
		phoneTail = profile.Phone[len(profile.Phone)-4:]
	}

	cols := dao.TavernBottleDeposit.Columns()
	depositNo := "DY" + gtime.Now().Format("YmdHis")
	labelCode := "LK" + gconv.String(gtime.TimestampMilli())
	expireAt := gtime.New(gtime.Now().Unix() + int64(180*24*3600))

	id, err := dao.TavernBottleDeposit.Ctx(ctx).Data(g.Map{
		cols.StoreId:       1,
		cols.DepositNo:     depositNo,
		cols.LabelCode:     labelCode,
		cols.MemberId:      memberId,
		cols.MemberName:    profile.Nickname,
		cols.PhoneTail:     phoneTail,
		cols.SourceType:    3,
		cols.SourceOrderId: in.OrderId,
		cols.SessionId:     order[orderCols.SessionId].Int64(),
		cols.GoodsName:     item[itemCols.GoodsName].String(),
		cols.Spec:          item[itemCols.SkuName].String(),
		cols.TotalRatio:    100,
		cols.RemainRatio:   100,
		cols.SlotCode:      "待分配",
		cols.Status:        1,
		cols.DepositAt:     gtime.Now(),
		cols.ExpireAt:      expireAt,
		cols.CreatedBy:     memberId,
		cols.CreatedAt:     gtime.Now(),
		cols.UpdatedAt:     gtime.Now(),
	}).InsertAndGetId()
	if err != nil {
		return nil, gerror.Wrap(err, "寄存失败")
	}

	row, err := dao.TavernBottleDeposit.Ctx(ctx).WherePri(id).One()
	if err != nil {
		return nil, err
	}
	return toBottleModel(row), nil
}
