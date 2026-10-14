package tavern

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"

	"hotgo/internal/dao"
	"hotgo/internal/model/input/tavernin"
	"hotgo/internal/service"
)

// sTavernAdmin 后台管理实现（HotGo 后台，登录态由 AdminAuth 保证）
type sTavernAdmin struct{}

func init() {
	service.RegisterTavernAdmin(NewTavernAdmin())
}

func NewTavernAdmin() *sTavernAdmin {
	return &sTavernAdmin{}
}

/* ---------------- 门店配置 ---------------- */

func (s *sTavernAdmin) StoreView(ctx context.Context) (res *tavernin.AdminStoreViewModel, err error) {
	cols := dao.TavernStore.Columns()
	row, err := dao.TavernStore.Ctx(ctx).OrderDesc(cols.Id).One()
	if err != nil {
		return
	}
	res = &tavernin.AdminStoreViewModel{
		BusinessOpen:     "18:00",
		BusinessClose:    "02:00",
		LastOrderGap:     30,
		PayFirst:         true,
		PayExpireMinutes: 5,
	}
	if row.IsEmpty() {
		return
	}
	res.Id = row[cols.Id].Int64()
	res.Name = row[cols.Name].String()
	res.Brief = row[cols.Brief].String()
	res.Address = row[cols.Address].String()
	res.Phone = row[cols.Phone].String()
	res.BusinessOpen = row[cols.BusinessOpen].String()
	res.BusinessClose = row[cols.BusinessClose].String()
	res.LastOrderGap = row[cols.LastOrderGap].Int()
	res.PayFirst = row[cols.PayFirst].Int() == 1
	res.PayExpireMinutes = row[cols.PayExpireMinutes].Int()
	return
}

func (s *sTavernAdmin) StoreEdit(ctx context.Context, in *tavernin.AdminStoreEditInp) (err error) {
	cols := dao.TavernStore.Columns()
	payFirst := 0
	if in.PayFirst {
		payFirst = 1
	}
	data := g.Map{
		cols.Name:             in.Name,
		cols.Brief:            in.Brief,
		cols.Address:          in.Address,
		cols.Phone:            in.Phone,
		cols.BusinessOpen:     in.BusinessOpen,
		cols.BusinessClose:    in.BusinessClose,
		cols.LastOrderGap:     in.LastOrderGap,
		cols.PayFirst:         payFirst,
		cols.PayExpireMinutes: in.PayExpireMinutes,
		cols.UpdatedAt:        gtime.Now(),
	}

	if in.Id > 0 {
		_, err = dao.TavernStore.Ctx(ctx).WherePri(in.Id).Data(data).Update()
		return err
	}
	_, err = dao.TavernStore.Ctx(ctx).Data(data).Insert()
	return err
}

/* ---------------- 区域 ---------------- */

func (s *sTavernAdmin) AreaList(ctx context.Context, in *tavernin.AdminAreaListInp) (list []*tavernin.AdminAreaModel, totalCount int, err error) {
	cols := dao.TavernArea.Columns()
	m := dao.TavernArea.Ctx(ctx)
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
	list = make([]*tavernin.AdminAreaModel, 0, len(rows))
	for _, row := range rows {
		list = append(list, &tavernin.AdminAreaModel{
			Id:        row[cols.Id].Int64(),
			StoreId:   row[cols.StoreId].Int64(),
			Name:      row[cols.Name].String(),
			Status:    row[cols.Status].Int(),
			Sort:      row[cols.Sort].Int(),
			Remark:    row[cols.Remark].String(),
			CreatedAt: row[cols.CreatedAt].GTime(),
			UpdatedAt: row[cols.UpdatedAt].GTime(),
		})
	}
	return
}

func (s *sTavernAdmin) AreaEdit(ctx context.Context, in *tavernin.AdminAreaEditInp) (err error) {
	cols := dao.TavernArea.Columns()
	data := g.Map{
		cols.Name:      in.Name,
		cols.Status:    in.Status,
		cols.Sort:      in.Sort,
		cols.Remark:    in.Remark,
		cols.UpdatedAt: gtime.Now(),
	}
	if in.Id > 0 {
		_, err = dao.TavernArea.Ctx(ctx).WherePri(in.Id).Data(data).Update()
		return err
	}
	data[cols.StoreId] = 1
	_, err = dao.TavernArea.Ctx(ctx).Data(data).Insert()
	return err
}

func (s *sTavernAdmin) AreaDelete(ctx context.Context, in *tavernin.AdminDeleteInp) (err error) {
	cols := dao.TavernArea.Columns()
	_, err = dao.TavernArea.Ctx(ctx).WhereIn(cols.Id, in.Ids).Delete()
	return err
}

/* ---------------- 桌位 ---------------- */

func (s *sTavernAdmin) TableList(ctx context.Context, in *tavernin.AdminTableListInp) (list []*tavernin.AdminTableModel, totalCount int, err error) {
	cols := dao.TavernTable.Columns()
	m := dao.TavernTable.Ctx(ctx)
	if in.Name != "" {
		m = m.WhereLike(cols.Name, "%"+in.Name+"%")
	}
	if in.AreaId > 0 {
		m = m.Where(cols.AreaId, in.AreaId)
	}
	totalCount, err = m.Count()
	if err != nil {
		return
	}
	rows, err := m.Page(in.Page, in.PerPage).OrderAsc(cols.Sort).OrderAsc(cols.Id).All()
	if err != nil {
		return
	}

	areaCols := dao.TavernArea.Columns()
	areaRows, _ := dao.TavernArea.Ctx(ctx).Fields(areaCols.Id, areaCols.Name).All()
	areaName := map[int64]string{}
	for _, r := range areaRows {
		areaName[r[areaCols.Id].Int64()] = r[areaCols.Name].String()
	}

	list = make([]*tavernin.AdminTableModel, 0, len(rows))
	for _, row := range rows {
		areaId := row[cols.AreaId].Int64()
		list = append(list, &tavernin.AdminTableModel{
			Id:         row[cols.Id].Int64(),
			AreaId:     areaId,
			AreaName:   areaName[areaId],
			Name:       row[cols.Name].String(),
			Capacity:   row[cols.Capacity].Int(),
			Scene:      row[cols.QrcodeScene].String(),
			SeatStatus: row[cols.SeatStatus].Int(),
			Status:     row[cols.Status].Int(),
			Sort:       row[cols.Sort].Int(),
			CreatedAt:  row[cols.CreatedAt].GTime(),
		})
	}
	return
}

func (s *sTavernAdmin) TableEdit(ctx context.Context, in *tavernin.AdminTableEditInp) (err error) {
	cols := dao.TavernTable.Columns()

	// 桌号唯一
	exist, err := dao.TavernTable.Ctx(ctx).
		Where(cols.Name, in.Name).
		WhereNot(cols.Id, in.Id).
		One()
	if err != nil {
		return err
	}
	if !exist.IsEmpty() {
		return gerror.New("桌号已存在")
	}

	data := g.Map{
		cols.AreaId:      in.AreaId,
		cols.Name:        in.Name,
		cols.Capacity:    in.Capacity,
		cols.Status:      in.Status,
		cols.Sort:        in.Sort,
		cols.QrcodeScene: "t=" + in.Name,
		cols.UpdatedAt:   gtime.Now(),
	}
	if in.Id > 0 {
		_, err = dao.TavernTable.Ctx(ctx).WherePri(in.Id).Data(data).Update()
		return err
	}
	data[cols.StoreId] = 1
	data[cols.SeatStatus] = 1
	_, err = dao.TavernTable.Ctx(ctx).Data(data).Insert()
	return err
}

func (s *sTavernAdmin) TableDelete(ctx context.Context, in *tavernin.AdminDeleteInp) (err error) {
	cols := dao.TavernTable.Columns()

	// 有进行中会话的桌位不允许删除
	sessionCols := dao.TavernTableSession.Columns()
	inUse, err := dao.TavernTableSession.Ctx(ctx).
		WhereIn(sessionCols.TableId, in.Ids).
		Where(sessionCols.Status, 1).
		Count()
	if err != nil {
		return err
	}
	if inUse > 0 {
		return gerror.New("所选桌位存在进行中的用餐会话，请先清台")
	}

	_, err = dao.TavernTable.Ctx(ctx).WhereIn(cols.Id, in.Ids).Delete()
	return err
}

/* ---------------- 分类 ---------------- */

func (s *sTavernAdmin) CategoryList(ctx context.Context, in *tavernin.AdminCategoryListInp) (list []*tavernin.AdminCategoryModel, totalCount int, err error) {
	cols := dao.TavernCategory.Columns()
	m := dao.TavernCategory.Ctx(ctx)
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
	list = make([]*tavernin.AdminCategoryModel, 0, len(rows))
	for _, row := range rows {
		list = append(list, &tavernin.AdminCategoryModel{
			Id:        row[cols.Id].Int64(),
			Name:      row[cols.Name].String(),
			Icon:      row[cols.Icon].String(),
			TimeSlots: row[cols.TimeSlots].String(),
			Status:    row[cols.Status].Int(),
			Sort:      row[cols.Sort].Int(),
			CreatedAt: row[cols.CreatedAt].GTime(),
		})
	}
	return
}

func (s *sTavernAdmin) CategoryEdit(ctx context.Context, in *tavernin.AdminCategoryEditInp) (err error) {
	cols := dao.TavernCategory.Columns()
	data := g.Map{
		cols.Name:      in.Name,
		cols.Icon:      in.Icon,
		cols.TimeSlots: in.TimeSlots,
		cols.Status:    in.Status,
		cols.Sort:      in.Sort,
		cols.UpdatedAt: gtime.Now(),
	}
	if in.Id > 0 {
		_, err = dao.TavernCategory.Ctx(ctx).WherePri(in.Id).Data(data).Update()
		return err
	}
	data[cols.StoreId] = 1
	_, err = dao.TavernCategory.Ctx(ctx).Data(data).Insert()
	return err
}

func (s *sTavernAdmin) CategoryDelete(ctx context.Context, in *tavernin.AdminDeleteInp) (err error) {
	goodsCols := dao.TavernGoods.Columns()
	count, err := dao.TavernGoods.Ctx(ctx).WhereIn(goodsCols.CategoryId, in.Ids).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New("该分类下仍有商品，请先移除商品")
	}
	cols := dao.TavernCategory.Columns()
	_, err = dao.TavernCategory.Ctx(ctx).WhereIn(cols.Id, in.Ids).Delete()
	return err
}

/* ---------------- 商品（含规格子表） ---------------- */

func (s *sTavernAdmin) GoodsList(ctx context.Context, in *tavernin.AdminGoodsListInp) (list []*tavernin.AdminGoodsModel, totalCount int, err error) {
	cols := dao.TavernGoods.Columns()
	m := dao.TavernGoods.Ctx(ctx)
	if in.Name != "" {
		m = m.WhereLike(cols.Name, "%"+in.Name+"%")
	}
	if in.CategoryId > 0 {
		m = m.Where(cols.CategoryId, in.CategoryId)
	}
	if in.Status > 0 {
		m = m.Where(cols.Status, in.Status)
	}
	totalCount, err = m.Count()
	if err != nil {
		return
	}
	rows, err := m.Page(in.Page, in.PerPage).OrderAsc(cols.Sort).OrderAsc(cols.Id).All()
	if err != nil {
		return
	}

	catCols := dao.TavernCategory.Columns()
	catRows, _ := dao.TavernCategory.Ctx(ctx).Fields(catCols.Id, catCols.Name).All()
	catName := map[int64]string{}
	for _, r := range catRows {
		catName[r[catCols.Id].Int64()] = r[catCols.Name].String()
	}

	list = make([]*tavernin.AdminGoodsModel, 0, len(rows))
	for _, row := range rows {
		list = append(list, toAdminGoodsModel(row, catName[row[cols.CategoryId].Int64()], nil))
	}
	return
}

func (s *sTavernAdmin) GoodsView(ctx context.Context, in *tavernin.AdminGoodsViewInp) (res *tavernin.AdminGoodsModel, err error) {
	cols := dao.TavernGoods.Columns()
	row, err := dao.TavernGoods.Ctx(ctx).WherePri(in.Id).One()
	if err != nil {
		return
	}
	if row.IsEmpty() {
		return nil, gerror.New("商品不存在")
	}

	categoryName := ""
	catCols := dao.TavernCategory.Columns()
	if cat, _ := dao.TavernCategory.Ctx(ctx).WherePri(row[cols.CategoryId].Int64()).One(); !cat.IsEmpty() {
		categoryName = cat[catCols.Name].String()
	}

	skuCols := dao.TavernGoodsSku.Columns()
	skuRows, _ := dao.TavernGoodsSku.Ctx(ctx).Where(skuCols.GoodsId, in.Id).OrderAsc(skuCols.Sort).All()
	skus := make([]*tavernin.AdminGoodsSkuModel, 0, len(skuRows))
	for _, r := range skuRows {
		skus = append(skus, &tavernin.AdminGoodsSkuModel{
			Id:            r[skuCols.Id].Int64(),
			GoodsId:       r[skuCols.GoodsId].Int64(),
			Name:          r[skuCols.Name].String(),
			Price:         r[skuCols.Price].Float64(),
			Stock:         r[skuCols.Stock].Int(),
			UnitCups:      r[skuCols.UnitCups].Float64(),
			IsDepositUnit: r[skuCols.IsDepositUnit].Int() == 1,
			Status:        r[skuCols.Status].Int(),
			Sort:          r[skuCols.Sort].Int(),
		})
	}

	res = toAdminGoodsModel(row, categoryName, skus)
	return
}

func (s *sTavernAdmin) GoodsEdit(ctx context.Context, in *tavernin.AdminGoodsEditInp) (err error) {
	cols := dao.TavernGoods.Columns()
	isWine, onlyDineIn, isCleared := 0, 0, 0
	if in.IsWine {
		isWine = 1
	}
	if in.OnlyDineIn {
		onlyDineIn = 1
	}
	if in.IsCleared {
		isCleared = 1
	}
	data := g.Map{
		cols.CategoryId:  in.CategoryId,
		cols.Name:        in.Name,
		cols.Cover:       in.Cover,
		cols.Brief:       in.Brief,
		cols.Tags:        in.Tags,
		cols.IsWine:      isWine,
		cols.OnlyDineIn:  onlyDineIn,
		cols.IsCleared:   isCleared,
		cols.MakeMinutes: in.MakeMinutes,
		cols.Status:      in.Status,
		cols.Sort:        in.Sort,
		cols.UpdatedAt:   gtime.Now(),
	}

	// 展示价取最低规格价
	var minPrice float64
	for _, sku := range in.Skus {
		if minPrice == 0 || sku.Price < minPrice {
			minPrice = sku.Price
		}
	}
	data[cols.Price] = minPrice

	var goodsId int64
	if in.Id > 0 {
		goodsId = in.Id
		if _, err = dao.TavernGoods.Ctx(ctx).WherePri(in.Id).Data(data).Update(); err != nil {
			return err
		}
	} else {
		data[cols.StoreId] = 1
		goodsId, err = dao.TavernGoods.Ctx(ctx).Data(data).InsertAndGetId()
		if err != nil {
			return gerror.Wrap(err, "保存商品失败")
		}
	}

	// 规格：按前端提交的 id 增量维护（有 id 更新，无 id 新增，未出现的删除）
	skuCols := dao.TavernGoodsSku.Columns()
	keepIds := make([]int64, 0, len(in.Skus))
	for i, sku := range in.Skus {
		depositUnit := 0
		if sku.IsDepositUnit {
			depositUnit = 1
		}
		stock := sku.Stock
		if stock == 0 {
			stock = -1 // 0 视为不限量
		}
		unitCups := sku.UnitCups
		if unitCups <= 0 {
			unitCups = 1
		}
		skuData := g.Map{
			skuCols.GoodsId:       goodsId,
			skuCols.Name:          sku.Name,
			skuCols.Price:         sku.Price,
			skuCols.Stock:         stock,
			skuCols.UnitCups:      unitCups,
			skuCols.IsDepositUnit: depositUnit,
			skuCols.Status:        1,
			skuCols.Sort:          i + 1,
			skuCols.UpdatedAt:     gtime.Now(),
		}
		if sku.Id > 0 {
			keepIds = append(keepIds, sku.Id)
			if _, err = dao.TavernGoodsSku.Ctx(ctx).WherePri(sku.Id).Data(skuData).Update(); err != nil {
				return err
			}
			continue
		}
		skuData[skuCols.StoreId] = 1
		id, e := dao.TavernGoodsSku.Ctx(ctx).Data(skuData).InsertAndGetId()
		if e != nil {
			return gerror.Wrap(e, "保存规格失败")
		}
		keepIds = append(keepIds, id)
	}

	del := dao.TavernGoodsSku.Ctx(ctx).Where(skuCols.GoodsId, goodsId)
	if len(keepIds) > 0 {
		del = del.WhereNotIn(skuCols.Id, keepIds)
	}
	if _, err = del.Delete(); err != nil {
		return err
	}
	return
}

func (s *sTavernAdmin) GoodsDelete(ctx context.Context, in *tavernin.AdminDeleteInp) (err error) {
	cols := dao.TavernGoods.Columns()
	if _, err = dao.TavernGoodsSku.Ctx(ctx).WhereIn(dao.TavernGoodsSku.Columns().GoodsId, in.Ids).Delete(); err != nil {
		return err
	}
	_, err = dao.TavernGoods.Ctx(ctx).WhereIn(cols.Id, in.Ids).Delete()
	return err
}

func (s *sTavernAdmin) GoodsStatus(ctx context.Context, in *tavernin.AdminGoodsStatusInp) (err error) {
	cols := dao.TavernGoods.Columns()
	_, err = dao.TavernGoods.Ctx(ctx).WherePri(in.Id).Data(g.Map{
		cols.Status:    in.Status,
		cols.UpdatedAt: gtime.Now(),
	}).Update()
	return err
}

func toAdminGoodsModel(row gdb.Record, categoryName string, skus []*tavernin.AdminGoodsSkuModel) *tavernin.AdminGoodsModel {
	cols := dao.TavernGoods.Columns()
	if skus == nil {
		skus = []*tavernin.AdminGoodsSkuModel{}
	}
	return &tavernin.AdminGoodsModel{
		Id:           row[cols.Id].Int64(),
		CategoryId:   row[cols.CategoryId].Int64(),
		CategoryName: categoryName,
		Name:         row[cols.Name].String(),
		Cover:        row[cols.Cover].String(),
		Brief:        row[cols.Brief].String(),
		Price:        row[cols.Price].Float64(),
		Tags:         row[cols.Tags].String(),
		IsWine:       row[cols.IsWine].Int() == 1,
		OnlyDineIn:   row[cols.OnlyDineIn].Int() == 1,
		IsCleared:    row[cols.IsCleared].Int() == 1,
		MakeMinutes:  row[cols.MakeMinutes].Int(),
		Status:       row[cols.Status].Int(),
		Sort:         row[cols.Sort].Int(),
		CreatedAt:    row[cols.CreatedAt].GTime(),
		Skus:         skus,
	}
}

/* ---------------- 订单查询 ---------------- */

func (s *sTavernAdmin) OrderList(ctx context.Context, in *tavernin.AdminOrderListInp) (list []*tavernin.AdminOrderModel, totalCount int, err error) {
	cols := dao.TavernOrder.Columns()
	m := dao.TavernOrder.Ctx(ctx)
	if in.OrderNo != "" {
		m = m.WhereLike(cols.OrderNo, "%"+in.OrderNo+"%")
	}
	if in.TableName != "" {
		m = m.Where(cols.TableName, in.TableName)
	}
	if in.PayStatus > 0 {
		m = m.Where(cols.PayStatus, in.PayStatus)
	}
	if len(in.DateRange) == 2 && in.DateRange[0] != nil && in.DateRange[1] != nil {
		m = m.WhereBetween(cols.CreatedAt, in.DateRange[0], in.DateRange[1])
	}
	totalCount, err = m.Count()
	if err != nil {
		return
	}
	rows, err := m.Page(in.Page, in.PerPage).OrderDesc(cols.Id).All()
	if err != nil {
		return
	}

	list = make([]*tavernin.AdminOrderModel, 0, len(rows))
	for _, row := range rows {
		list = append(list, &tavernin.AdminOrderModel{
			Id:           row[cols.Id].Int64(),
			OrderNo:      row[cols.OrderNo].String(),
			TableName:    row[cols.TableName].String(),
			OrderSeq:     row[cols.OrderSeq].Int(),
			MemberName:   row[cols.NicknameSnapshot].String(),
			PayStatus:    row[cols.PayStatus].Int(),
			MakingStatus: row[cols.MakingStatus].Int(),
			RefundStatus: row[cols.RefundStatus].Int(),
			PayAmount:    row[cols.PayAmount].Float64(),
			RefundAmount: row[cols.RefundAmount].Float64(),
			PayChannel:   row[cols.PayChannel].String(),
			PayTime:      row[cols.PayTime].GTime(),
			CreatedAt:    row[cols.CreatedAt].GTime(),
		})
	}
	return
}

/* ---------------- 存酒台账 ---------------- */

func (s *sTavernAdmin) BottleList(ctx context.Context, in *tavernin.AdminBottleListInp) (list []*tavernin.BottleModel, totalCount int, err error) {
	cols := dao.TavernBottleDeposit.Columns()
	m := dao.TavernBottleDeposit.Ctx(ctx)
	if in.Keyword != "" {
		kw := "%" + in.Keyword + "%"
		m = m.Where(
			"(member_name LIKE ? OR goods_name LIKE ? OR label_code LIKE ? OR phone_tail LIKE ?)",
			kw, kw, kw, kw,
		)
	}
	if in.Status > 0 {
		m = m.Where(cols.Status, in.Status)
	}
	if in.SourceType > 0 {
		m = m.Where(cols.SourceType, in.SourceType)
	}
	totalCount, err = m.Count()
	if err != nil {
		return
	}
	rows, err := m.Page(in.Page, in.PerPage).OrderDesc(cols.Id).All()
	if err != nil {
		return
	}
	list = make([]*tavernin.BottleModel, 0, len(rows))
	for _, row := range rows {
		list = append(list, toBottleModel(row))
	}
	return
}

func (s *sTavernAdmin) BottleUpdate(ctx context.Context, in *tavernin.AdminBottleUpdateInp) (err error) {
	cols := dao.TavernBottleDeposit.Columns()
	row, err := dao.TavernBottleDeposit.Ctx(ctx).WherePri(in.Id).One()
	if err != nil {
		return err
	}
	if row.IsEmpty() {
		return gerror.New("存酒记录不存在")
	}

	data := g.Map{cols.UpdatedAt: gtime.Now()}
	if in.Status > 0 {
		data[cols.Status] = in.Status
	}
	if in.RemainRatio >= 0 && in.RemainRatio <= 100 && in.RemainRatio != row[cols.RemainRatio].Float64() {
		data[cols.RemainRatio] = in.RemainRatio
	}
	if in.SlotCode != "" && in.SlotCode != row[cols.SlotCode].String() {
		data[cols.SlotCode] = in.SlotCode
	}
	if in.ExpireAt != "" {
		data[cols.ExpireAt] = gtime.NewFromStr(in.ExpireAt)
	}
	if in.Remark != "" {
		data[cols.Remark] = gstr.Trim(in.Remark)
	}

	_, err = dao.TavernBottleDeposit.Ctx(ctx).WherePri(in.Id).Data(data).Update()
	return err
}
