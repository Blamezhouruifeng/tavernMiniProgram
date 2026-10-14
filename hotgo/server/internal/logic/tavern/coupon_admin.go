package tavern

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"

	"hotgo/internal/dao"
	"hotgo/internal/model/input/tavernin"
	"hotgo/internal/service"
)

// CouponList 优惠券模板列表（含发放/使用统计）
func (s *sTavernAdmin) CouponList(ctx context.Context, in *tavernin.AdminCouponListInp) (list []*tavernin.AdminCouponModel, totalCount int, err error) {
	cols := dao.TavernCoupon.Columns()
	m := dao.TavernCoupon.Ctx(ctx)
	if in.Name != "" {
		m = m.WhereLike(cols.Name, "%"+in.Name+"%")
	}
	totalCount, err = m.Count()
	if err != nil {
		return
	}
	rows, err := m.Page(in.Page, in.PerPage).OrderAsc(cols.Sort).OrderAsc(cols.Id).All()
	if err != nil {
		return
	}

	memberCols := dao.TavernCouponMember.Columns()
	list = make([]*tavernin.AdminCouponModel, 0, len(rows))
	for _, row := range rows {
		id := row[cols.Id].Int64()
		granted, _ := dao.TavernCouponMember.Ctx(ctx).Where(memberCols.CouponId, id).Count()
		used, _ := dao.TavernCouponMember.Ctx(ctx).
			Where(memberCols.CouponId, id).
			Where(memberCols.Status, 2).
			Count()
		list = append(list, &tavernin.AdminCouponModel{
			Id:             id,
			Name:           row[cols.Name].String(),
			Type:           row[cols.Type].Int(),
			Amount:         row[cols.Amount].Float64(),
			MinAmount:      row[cols.MinAmount].Float64(),
			Discount:       row[cols.Discount].Float64(),
			ValidDays:      row[cols.ValidDays].Int(),
			TotalLimit:     row[cols.TotalLimit].Int(),
			PerMemberLimit: row[cols.PerMemberLimit].Int(),
			IsNewcomer:     row[cols.IsNewcomer].Int() == 1,
			Status:         row[cols.Status].Int(),
			Sort:           row[cols.Sort].Int(),
			Remark:         row[cols.Remark].String(),
			GrantedCount:   granted,
			UsedCount:      used,
			CreatedAt:      row[cols.CreatedAt].GTime(),
		})
	}
	return
}

// CouponEdit 保存优惠券模板
func (s *sTavernAdmin) CouponEdit(ctx context.Context, in *tavernin.AdminCouponEditInp) (err error) {
	cols := dao.TavernCoupon.Columns()
	isNewcomer := 0
	if in.IsNewcomer {
		isNewcomer = 1
	}
	data := g.Map{
		cols.Name:           in.Name,
		cols.Type:           in.Type,
		cols.Amount:         in.Amount,
		cols.MinAmount:      in.MinAmount,
		cols.Discount:       in.Discount,
		cols.ValidDays:      in.ValidDays,
		cols.TotalLimit:     in.TotalLimit,
		cols.PerMemberLimit: in.PerMemberLimit,
		cols.IsNewcomer:     isNewcomer,
		cols.Status:         in.Status,
		cols.Sort:           in.Sort,
		cols.Remark:         in.Remark,
		cols.UpdatedAt:      gtime.Now(),
	}
	if in.Id > 0 {
		_, err = dao.TavernCoupon.Ctx(ctx).WherePri(in.Id).Data(data).Update()
		return err
	}
	data[cols.StoreId] = 1
	_, err = dao.TavernCoupon.Ctx(ctx).Data(data).Insert()
	return err
}

// CouponGrant 发放优惠券：支持按会员ID、按手机号、或全体会员
func (s *sTavernAdmin) CouponGrant(ctx context.Context, in *tavernin.AdminCouponGrantInp) (count int, err error) {
	memberIds := in.MemberIds

	if in.Phone != "" {
		memberCols := dao.TavernMember.Columns()
		row, e := dao.TavernMember.Ctx(ctx).Where(memberCols.Phone, in.Phone).One()
		if e != nil {
			return 0, e
		}
		if row.IsEmpty() {
			return 0, gerror.New("未找到该手机号对应的会员")
		}
		memberIds = append(memberIds, row[memberCols.Id].Int64())
	}

	if in.All {
		memberCols := dao.TavernMember.Columns()
		rows, e := dao.TavernMember.Ctx(ctx).Fields(memberCols.Id).Where(memberCols.Status, 1).Limit(2000).All()
		if e != nil {
			return 0, e
		}
		memberIds = make([]int64, 0, len(rows))
		for _, row := range rows {
			memberIds = append(memberIds, row[memberCols.Id].Int64())
		}
	}

	if len(memberIds) == 0 {
		return 0, gerror.New("请选择要发放的会员或手机号")
	}

	for _, memberId := range memberIds {
		if _, e := service.TavernCoupon().Grant(ctx, in.CouponId, memberId, 3); e != nil {
			g.Log().Warningf(ctx, "发放优惠券失败 memberId:%d couponId:%d err:%+v", memberId, in.CouponId, e)
			continue
		}
		count++
	}
	if count == 0 {
		return 0, gerror.New("发放失败：可能已达每人限领或发放总量上限")
	}

	g.Log().Infof(ctx, "酒馆优惠券发放完成 couponId:%d count:%d", in.CouponId, count)
	return count, nil
}

// 兼容未使用参数
var _ = gconv.String
