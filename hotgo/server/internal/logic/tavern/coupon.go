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
	"hotgo/internal/service"
)

// sTavernCoupon 优惠券（新人礼 / 活动发放 / 定向发放）
type sTavernCoupon struct{}

func init() {
	service.RegisterTavernCoupon(NewTavernCoupon())
}

func NewTavernCoupon() *sTavernCoupon {
	return &sTavernCoupon{}
}

// MyList 我的优惠券（可按状态与订单金额筛选可用）
func (s *sTavernCoupon) MyList(ctx context.Context, in *tavernin.CouponListInp) (res []*tavernin.CouponModel, err error) {
	memberId := service.TavernMember().CurrentId(ctx)
	if memberId == 0 {
		return nil, gerror.New("请先登录")
	}

	cols := dao.TavernCouponMember.Columns()
	m := dao.TavernCouponMember.Ctx(ctx).Where(cols.MemberId, memberId)
	if in.Status > 0 {
		m = m.Where(cols.Status, in.Status)
	}
	rows, err := m.OrderAsc(cols.Status).OrderAsc(cols.ExpireAt).Limit(100).All()
	if err != nil {
		return nil, err
	}

	couponIds := make([]int64, 0, len(rows))
	for _, row := range rows {
		couponIds = append(couponIds, row[cols.CouponId].Int64())
	}
	templates := map[int64]gdb.Record{}
	if len(couponIds) > 0 {
		tplCols := dao.TavernCoupon.Columns()
		tplRows, e := dao.TavernCoupon.Ctx(ctx).WhereIn(tplCols.Id, couponIds).All()
		if e != nil {
			return nil, e
		}
		for _, t := range tplRows {
			templates[t[tplCols.Id].Int64()] = t
		}
	}

	res = make([]*tavernin.CouponModel, 0, len(rows))
	for _, row := range rows {
		couponId := row[cols.CouponId].Int64()
		tpl := templates[couponId]

		model := &tavernin.CouponModel{
			Id:       row[cols.Id].Int64(),
			CouponId: couponId,
			Code:     row[cols.Code].String(),
			Source:   row[cols.Source].Int(),
			Status:   row[cols.Status].Int(),
			ExpireAt: row[cols.ExpireAt].GTime(),
		}
		if tpl != nil {
			tplCols := dao.TavernCoupon.Columns()
			model.Name = tpl[tplCols.Name].String()
			model.Type = tpl[tplCols.Type].Int()
			model.Amount = tpl[tplCols.Amount].Float64()
			model.MinAmount = tpl[tplCols.MinAmount].Float64()
			model.Discount = tpl[tplCols.Discount].Float64()
		}

		// 可用性判定
		model.Usable, model.UnusableReason = judgeCoupon(model, row[cols.Status].Int(), row[cols.ExpireAt].GTime(), in.Amount)
		res = append(res, model)
	}
	return
}

func judgeCoupon(model *tavernin.CouponModel, status int, expireAt *gtime.Time, amount float64) (usable bool, reason string) {
	if status == 2 {
		return false, "已使用"
	}
	if status == 3 || (expireAt != nil && expireAt.Before(gtime.Now())) {
		return false, "已过期"
	}
	if amount > 0 && amount < model.MinAmount {
		return false, "未满 " + gconv.String(model.MinAmount) + " 元"
	}
	return true, ""
}

// CalcDiscount 试算优惠金额（不落库）
func (s *sTavernCoupon) CalcDiscount(ctx context.Context, memberId, couponMemberId int64, amount float64) (discount float64, err error) {
	row, model, err := s.loadMemberCoupon(ctx, memberId, couponMemberId)
	if err != nil {
		return 0, err
	}
	_ = row

	if usable, reason := judgeCoupon(model, model.Status, model.ExpireAt, amount); !usable {
		return 0, gerror.New("优惠券不可用：" + reason)
	}

	switch model.Type {
	case 2: // 折扣
		if model.Discount <= 0 || model.Discount >= 1 {
			return 0, gerror.New("折扣券配置不正确")
		}
		discount = round2(amount * (1 - model.Discount))
	default: // 满减
		discount = model.Amount
	}
	if discount > amount {
		discount = amount
	}
	return discount, nil
}

// UseCoupon 标记券已使用（支付成功时调用）
func (s *sTavernCoupon) UseCoupon(ctx context.Context, memberId, couponMemberId, orderId int64, discount float64) (err error) {
	if couponMemberId == 0 {
		return nil
	}
	cols := dao.TavernCouponMember.Columns()
	_, err = dao.TavernCouponMember.Ctx(ctx).
		WherePri(couponMemberId).
		Where(cols.MemberId, memberId).
		Data(g.Map{
			cols.Status:    2,
			cols.OrderId:   orderId,
			cols.Amount:    discount,
			cols.UsedAt:    gtime.Now(),
			cols.UpdatedAt: gtime.Now(),
		}).Update()
	return err
}

// ReturnCoupon 订单退款时退回优惠券
func (s *sTavernCoupon) ReturnCoupon(ctx context.Context, orderId int64) (err error) {
	cols := dao.TavernCouponMember.Columns()
	_, err = dao.TavernCouponMember.Ctx(ctx).
		Where(cols.OrderId, orderId).
		Where(cols.Status, 2).
		Data(g.Map{
			cols.Status:    1,
			cols.OrderId:   0,
			cols.Amount:    0,
			cols.UsedAt:    nil,
			cols.UpdatedAt: gtime.Now(),
		}).Update()
	return err
}

// GrantNewcomer 新人礼：首次建档时自动发放
func (s *sTavernCoupon) GrantNewcomer(ctx context.Context, memberId int64) (err error) {
	cols := dao.TavernCoupon.Columns()
	rows, err := dao.TavernCoupon.Ctx(ctx).
		Where(cols.IsNewcomer, 1).
		Where(cols.Status, 1).
		All()
	if err != nil {
		return err
	}
	for _, row := range rows {
		if _, err = s.Grant(ctx, row[cols.Id].Int64(), memberId, 1); err != nil {
			return
		}
	}
	return nil
}

// Grant 发放优惠券（source：1新人礼 2活动 3后台定向）
func (s *sTavernCoupon) Grant(ctx context.Context, couponId, memberId int64, source int) (res *tavernin.CouponModel, err error) {
	tplCols := dao.TavernCoupon.Columns()
	tpl, err := dao.TavernCoupon.Ctx(ctx).WherePri(couponId).One()
	if err != nil {
		return nil, err
	}
	if tpl.IsEmpty() || tpl[tplCols.Status].Int() != 1 {
		return nil, gerror.New("优惠券不存在或已停用")
	}

	cols := dao.TavernCouponMember.Columns()

	// 每人限领
	limit := tpl[tplCols.PerMemberLimit].Int()
	if limit > 0 {
		count, e := dao.TavernCouponMember.Ctx(ctx).
			Where(cols.CouponId, couponId).
			Where(cols.MemberId, memberId).
			Count()
		if e != nil {
			return nil, e
		}
		if count >= limit {
			return nil, gerror.New("该优惠券每位会员限领 " + gconv.String(limit) + " 张")
		}
	}

	// 发放总量
	totalLimit := tpl[tplCols.TotalLimit].Int()
	if totalLimit >= 0 {
		count, e := dao.TavernCouponMember.Ctx(ctx).Where(cols.CouponId, couponId).Count()
		if e != nil {
			return nil, e
		}
		if count >= totalLimit {
			return nil, gerror.New("该优惠券已发放完")
		}
	}

	validDays := tpl[tplCols.ValidDays].Int()
	if validDays <= 0 {
		validDays = 30
	}

	id, err := dao.TavernCouponMember.Ctx(ctx).Data(g.Map{
		cols.StoreId:   1,
		cols.CouponId:  couponId,
		cols.MemberId:  memberId,
		cols.Code:      "CP" + gtime.Now().Format("YmdHis") + gconv.String(gtime.TimestampMilli()%1000),
		cols.Source:    source,
		cols.Status:    1,
		cols.ExpireAt:  gtime.New(gtime.Now().Unix() + int64(validDays*24*3600)),
		cols.CreatedAt: gtime.Now(),
		cols.UpdatedAt: gtime.Now(),
	}).InsertAndGetId()
	if err != nil {
		return nil, gerror.Wrap(err, "发放优惠券失败")
	}

	row, err := dao.TavernCouponMember.Ctx(ctx).WherePri(id).One()
	if err != nil {
		return nil, err
	}

	res = &tavernin.CouponModel{
		Id:        row[cols.Id].Int64(),
		CouponId:  couponId,
		Name:      tpl[tplCols.Name].String(),
		Type:      tpl[tplCols.Type].Int(),
		Amount:    tpl[tplCols.Amount].Float64(),
		MinAmount: tpl[tplCols.MinAmount].Float64(),
		Discount:  tpl[tplCols.Discount].Float64(),
		Code:      row[cols.Code].String(),
		Source:    source,
		Status:    1,
		ExpireAt:  row[cols.ExpireAt].GTime(),
		Usable:    true,
	}
	return
}

// loadMemberCoupon 读取会员券与模板
func (s *sTavernCoupon) loadMemberCoupon(ctx context.Context, memberId, couponMemberId int64) (row gdb.Record, model *tavernin.CouponModel, err error) {
	cols := dao.TavernCouponMember.Columns()
	row, err = dao.TavernCouponMember.Ctx(ctx).WherePri(couponMemberId).One()
	if err != nil {
		return
	}
	if row.IsEmpty() || row[cols.MemberId].Int64() != memberId {
		return nil, nil, gerror.New("优惠券不存在")
	}

	tplCols := dao.TavernCoupon.Columns()
	tpl, e := dao.TavernCoupon.Ctx(ctx).WherePri(row[cols.CouponId].Int64()).One()
	if e != nil {
		return nil, nil, e
	}
	if tpl.IsEmpty() {
		return nil, nil, gerror.New("优惠券模板不存在")
	}

	model = &tavernin.CouponModel{
		Id:        row[cols.Id].Int64(),
		CouponId:  row[cols.CouponId].Int64(),
		Name:      tpl[tplCols.Name].String(),
		Type:      tpl[tplCols.Type].Int(),
		Amount:    tpl[tplCols.Amount].Float64(),
		MinAmount: tpl[tplCols.MinAmount].Float64(),
		Discount:  tpl[tplCols.Discount].Float64(),
		Code:      row[cols.Code].String(),
		Status:    row[cols.Status].Int(),
		ExpireAt:  row[cols.ExpireAt].GTime(),
	}
	return
}
