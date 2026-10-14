package tavernin

/* ============ 追加：绑定手机号 / 退款 / 员工代存与转台 ============ */

type BindPhoneInp struct {
	Phone string `json:"phone" v:"required|phone#手机号不能为空|手机号格式不正确"`
}

type RefundApplyInp struct {
	OrderNo string  `json:"orderNo" v:"required#订单号不能为空"`
	ItemIds []int64 `json:"itemIds"`
	Reason  string  `json:"reason" v:"required#请填写退款原因"`
}

type RefundAuditInp struct {
	OrderNo string `json:"orderNo" v:"required#订单号不能为空"`
	Approve bool   `json:"approve"`
	Reason  string `json:"reason"`
}

type DepositRegisterInp struct {
	MemberId    int64   `json:"memberId"`
	Phone       string  `json:"phone"`
	GoodsName   string  `json:"goodsName" v:"required#请填写酒品名称"`
	Spec        string  `json:"spec"`
	Degree      string  `json:"degree"`
	RemainRatio float64 `json:"remainRatio"`
	SlotCode    string  `json:"slotCode"`
	Photo       string  `json:"photo"`
	Remark      string  `json:"remark"`
}

type SessionCallInp struct {
	SessionId int64  `json:"sessionId" v:"required#会话ID不能为空"`
	Type      string `json:"type"`
	OrderNo   string `json:"orderNo"`
	Remark    string `json:"remark"`
}

type MoveTableInp struct {
	SessionId int64  `json:"sessionId" v:"required#会话ID不能为空"`
	ToTableNo string `json:"toTableNo" v:"required#目标桌号不能为空"`
}
