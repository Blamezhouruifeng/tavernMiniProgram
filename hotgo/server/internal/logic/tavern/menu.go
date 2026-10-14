package tavern

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"

	"hotgo/internal/dao"
	"hotgo/internal/model/input/tavernin"
	"hotgo/internal/service"
)

type sTavernMenu struct{}

func init() {
	service.RegisterTavernMenu(NewTavernMenu())
}

func NewTavernMenu() *sTavernMenu {
	return &sTavernMenu{}
}

// List 菜单（分类 + 商品 + 规格）
// 分时段：分类上的 time_slots 由前端展示时过滤，真实业务可在此按当前时间过滤
func (s *sTavernMenu) List(ctx context.Context) (res []*tavernin.MenuCategoryModel, err error) {
	catCols := dao.TavernCategory.Columns()
	categories, err := dao.TavernCategory.Ctx(ctx).
		Where(catCols.Status, 1).
		OrderAsc(catCols.Sort).
		OrderAsc(catCols.Id).
		All()
	if err != nil {
		return
	}

	goodsCols := dao.TavernGoods.Columns()
	goodsList, err := dao.TavernGoods.Ctx(ctx).
		Where(goodsCols.Status, 1).
		OrderAsc(goodsCols.Sort).
		OrderAsc(goodsCols.Id).
		All()
	if err != nil {
		return
	}

	skuCols := dao.TavernGoodsSku.Columns()
	skuList, err := dao.TavernGoodsSku.Ctx(ctx).
		Where(skuCols.Status, 1).
		OrderAsc(skuCols.Sort).
		OrderAsc(skuCols.Id).
		All()
	if err != nil {
		return
	}

	skusByGoods := map[int64][]*tavernin.MenuSkuModel{}
	for _, row := range skuList {
		m := &tavernin.MenuSkuModel{
			Id:            row[skuCols.Id].Int64(),
			GoodsId:       row[skuCols.GoodsId].Int64(),
			Name:          row[skuCols.Name].String(),
			Price:         row[skuCols.Price].Float64(),
			Stock:         row[skuCols.Stock].Int(),
			UnitCups:      row[skuCols.UnitCups].Float64(),
			IsDepositUnit: row[skuCols.IsDepositUnit].Int() == 1,
		}
		skusByGoods[m.GoodsId] = append(skusByGoods[m.GoodsId], m)
	}

	goodsByCategory := map[int64][]*tavernin.MenuGoodsModel{}
	for _, row := range goodsList {
		id := row[goodsCols.Id].Int64()
		m := &tavernin.MenuGoodsModel{
			Id:          id,
			CategoryId:  row[goodsCols.CategoryId].Int64(),
			Name:        row[goodsCols.Name].String(),
			Cover:       row[goodsCols.Cover].String(),
			Brief:       row[goodsCols.Brief].String(),
			Price:       row[goodsCols.Price].Float64(),
			IsWine:      row[goodsCols.IsWine].Int() == 1,
			OnlyDineIn:  row[goodsCols.OnlyDineIn].Int() == 1,
			IsCleared:   row[goodsCols.IsCleared].Int() == 1,
			MakeMinutes: row[goodsCols.MakeMinutes].Int(),
			Status:      row[goodsCols.Status].Int(),
			Skus:        skusByGoods[id],
		}
		if tags := row[goodsCols.Tags].String(); tags != "" {
			m.Tags = gstr.SplitAndTrim(tags, ",")
		}
		if m.Skus == nil {
			m.Skus = []*tavernin.MenuSkuModel{}
		}
		goodsByCategory[m.CategoryId] = append(goodsByCategory[m.CategoryId], m)
	}

	res = make([]*tavernin.MenuCategoryModel, 0, len(categories))
	for _, row := range categories {
		id := row[catCols.Id].Int64()
		m := &tavernin.MenuCategoryModel{
			Id:        id,
			Name:      row[catCols.Name].String(),
			Icon:      row[catCols.Icon].String(),
			TimeSlots: row[catCols.TimeSlots].String(),
			Goods:     goodsByCategory[id],
		}
		if m.Goods == nil {
			m.Goods = []*tavernin.MenuGoodsModel{}
		}
		res = append(res, m)
	}

	g.Log().Debug(ctx, "tavern menu loaded", gconv.String(len(res)))
	return
}
