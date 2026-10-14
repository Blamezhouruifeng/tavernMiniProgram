package tavern

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"hotgo/internal/dao"
	"hotgo/internal/model/input/tavernin"
	"hotgo/internal/service"
)

type sTavernCart struct{}

func init() {
	service.RegisterTavernCart(NewTavernCart())
}

func NewTavernCart() *sTavernCart {
	return &sTavernCart{}
}

// loadCartItems 会话购物车（含商品名 / 规格名 / 会员名，全桌可见）
func loadCartItems(ctx context.Context, sessionId int64) (res []*tavernin.CartItemModel, err error) {
	cols := dao.TavernCartItem.Columns()
	rows, err := dao.TavernCartItem.Ctx(ctx).
		Where(cols.SessionId, sessionId).
		WhereNull(cols.DeletedAt).
		OrderAsc(cols.Id).
		All()
	if err != nil {
		return
	}

	goodsCols := dao.TavernGoods.Columns()
	goodsRows, _ := dao.TavernGoods.Ctx(ctx).Fields(goodsCols.Id, goodsCols.Name).All()
	goodsName := map[int64]string{}
	for _, r := range goodsRows {
		goodsName[r[goodsCols.Id].Int64()] = r[goodsCols.Name].String()
	}

	skuCols := dao.TavernGoodsSku.Columns()
	skuRows, _ := dao.TavernGoodsSku.Ctx(ctx).Fields(skuCols.Id, skuCols.Name, skuCols.Price).All()
	skuName := map[int64]string{}
	skuPrice := map[int64]float64{}
	for _, r := range skuRows {
		skuName[r[skuCols.Id].Int64()] = r[skuCols.Name].String()
		skuPrice[r[skuCols.Id].Int64()] = r[skuCols.Price].Float64()
	}

	memberCols := dao.TavernSessionMember.Columns()
	memberRows, _ := dao.TavernSessionMember.Ctx(ctx).
		Where(memberCols.SessionId, sessionId).
		All()
	memberName := map[int64]string{}
	for _, r := range memberRows {
		memberName[r[memberCols.MemberId].Int64()] = r[memberCols.Nickname].String()
	}

	res = make([]*tavernin.CartItemModel, 0, len(rows))
	for _, row := range rows {
		goodsId := row[cols.GoodsId].Int64()
		skuId := row[cols.SkuId].Int64()
		memberId := row[cols.MemberId].Int64()
		res = append(res, &tavernin.CartItemModel{
			Id:         row[cols.Id].Int64(),
			SessionId:  row[cols.SessionId].Int64(),
			MemberId:   memberId,
			MemberName: memberName[memberId],
			GoodsId:    goodsId,
			SkuId:      skuId,
			GoodsName:  goodsName[goodsId],
			SkuName:    skuName[skuId],
			Price:      skuPrice[skuId],
			Qty:        row[cols.Qty].Int(),
			Remark:     row[cols.Remark].String(),
		})
	}
	return
}

// Add 加入购物车（每人一个篮子，同一行重复加购则累加数量）
func (s *sTavernCart) Add(ctx context.Context, in *tavernin.CartAddInp) (err error) {
	memberId := service.TavernMember().CurrentId(ctx)
	if memberId == 0 {
		return gerror.New("请先登录")
	}
	if _, err = service.TavernSession().CurrentSessionId(ctx, in.SessionId); err != nil {
		return err
	}

	goodsCols := dao.TavernGoods.Columns()
	goods, err := dao.TavernGoods.Ctx(ctx).WherePri(in.GoodsId).One()
	if err != nil {
		return err
	}
	if goods.IsEmpty() {
		return gerror.New("商品不存在")
	}
	if goods[goodsCols.Status].Int() != 1 {
		return gerror.New("该商品已下架")
	}
	if goods[goodsCols.IsCleared].Int() == 1 {
		return gerror.New(goods[goodsCols.Name].String() + " 已沽清")
	}

	// 规格校验
	skuCols := dao.TavernGoodsSku.Columns()
	sku, err := dao.TavernGoodsSku.Ctx(ctx).
		WherePri(in.SkuId).
		Where(skuCols.Status, 1).
		One()
	if err != nil {
		return err
	}
	if sku.IsEmpty() || sku[skuCols.GoodsId].Int64() != in.GoodsId {
		return gerror.New("规格不存在")
	}

	qty := in.Qty
	if qty <= 0 {
		qty = 1
	}

	cols := dao.TavernCartItem.Columns()
	exist, err := dao.TavernCartItem.Ctx(ctx).
		Where(cols.SessionId, in.SessionId).
		Where(cols.MemberId, memberId).
		Where(cols.GoodsId, in.GoodsId).
		Where(cols.SkuId, in.SkuId).
		Where(cols.Remark, in.Remark).
		WhereNull(cols.DeletedAt).
		One()
	if err != nil {
		return err
	}

	if !exist.IsEmpty() {
		_, err = dao.TavernCartItem.Ctx(ctx).
			WherePri(exist[cols.Id].Int64()).
			Data(g.Map{
				cols.Qty:       exist[cols.Qty].Int() + qty,
				cols.UpdatedAt: gtime.Now(),
			}).Update()
		return err
	}

	_, err = dao.TavernCartItem.Ctx(ctx).Data(g.Map{
		cols.StoreId:   1,
		cols.SessionId: in.SessionId,
		cols.MemberId:  memberId,
		cols.GoodsId:   in.GoodsId,
		cols.SkuId:     in.SkuId,
		cols.Qty:       qty,
		cols.Remark:    in.Remark,
		cols.CreatedAt: gtime.Now(),
		cols.UpdatedAt: gtime.Now(),
	}).Insert()
	return err
}

// Update 修改自己篮子里的行（归属不可越权）
func (s *sTavernCart) Update(ctx context.Context, in *tavernin.CartUpdateInp) (err error) {
	memberId := service.TavernMember().CurrentId(ctx)
	if memberId == 0 {
		return gerror.New("请先登录")
	}

	cols := dao.TavernCartItem.Columns()
	row, err := dao.TavernCartItem.Ctx(ctx).WherePri(in.Id).One()
	if err != nil {
		return err
	}
	if row.IsEmpty() || row[cols.DeletedAt].GTime() != nil {
		return gerror.New("购物车行不存在")
	}
	if row[cols.MemberId].Int64() != memberId {
		return gerror.New("只能修改自己添加的商品")
	}

	qty := in.Qty
	if qty <= 0 {
		qty = 1
	}

	_, err = dao.TavernCartItem.Ctx(ctx).WherePri(in.Id).Data(g.Map{
		cols.Qty:       qty,
		cols.Remark:    in.Remark,
		cols.UpdatedAt: gtime.Now(),
	}).Update()
	return err
}

// Remove 删除自己篮子里的行
func (s *sTavernCart) Remove(ctx context.Context, in *tavernin.CartRemoveInp) (err error) {
	memberId := service.TavernMember().CurrentId(ctx)
	if memberId == 0 {
		return gerror.New("请先登录")
	}

	cols := dao.TavernCartItem.Columns()
	row, err := dao.TavernCartItem.Ctx(ctx).WherePri(in.Id).One()
	if err != nil {
		return err
	}
	if row.IsEmpty() || row[cols.DeletedAt].GTime() != nil {
		return gerror.New("购物车行不存在")
	}
	if row[cols.MemberId].Int64() != memberId {
		return gerror.New("只能删除自己添加的商品")
	}

	_, err = dao.TavernCartItem.Ctx(ctx).WherePri(in.Id).Data(g.Map{
		cols.DeletedAt: gtime.Now(),
		cols.UpdatedAt: gtime.Now(),
	}).Update()
	return err
}
