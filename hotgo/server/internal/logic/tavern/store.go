package tavern

import (
	"context"

	"hotgo/internal/dao"
	"hotgo/internal/model/input/tavernin"
	"hotgo/internal/service"
)

type sTavernStore struct{}

func init() {
	service.RegisterTavernStore(NewTavernStore())
}

func NewTavernStore() *sTavernStore {
	return &sTavernStore{}
}

// Info 门店信息（含先付制与打烊截止配置）
func (s *sTavernStore) Info(ctx context.Context) (res *tavernin.StoreInfoModel, err error) {
	cols := dao.TavernStore.Columns()
	one, err := dao.TavernStore.Ctx(ctx).Where(cols.Status, 1).OrderDesc(cols.Id).One()
	if err != nil {
		return
	}

	res = &tavernin.StoreInfoModel{
		Name:             "晚风民谣小酒馆",
		BusinessOpen:     "18:00",
		BusinessClose:    "02:00",
		LastOrderGap:     30,
		PayFirst:         true,
		PayExpireMinutes: 5,
	}
	if one.IsEmpty() {
		return
	}

	var row struct {
		Id               int64
		Name             string
		Address          string
		Phone            string
		BusinessOpen     string
		BusinessClose    string
		LastOrderGap     int
		PayFirst         int
		PayExpireMinutes int
	}
	if err = one.Struct(&row); err != nil {
		return
	}

	res.Id = row.Id
	res.Name = row.Name
	res.Address = row.Address
	res.Phone = row.Phone
	res.PayFirst = row.PayFirst == 1
	if row.BusinessOpen != "" {
		res.BusinessOpen = row.BusinessOpen
	}
	if row.BusinessClose != "" {
		res.BusinessClose = row.BusinessClose
	}
	if row.LastOrderGap > 0 {
		res.LastOrderGap = row.LastOrderGap
	}
	if row.PayExpireMinutes > 0 {
		res.PayExpireMinutes = row.PayExpireMinutes
	}
	return
}
