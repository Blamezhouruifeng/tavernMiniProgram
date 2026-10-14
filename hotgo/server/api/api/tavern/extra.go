package tavern

import (
	"github.com/gogf/gf/v2/frame/g"

	"hotgo/internal/model/input/tavernin"
)

/* ---------------- 追加接口：手机号绑定 / 退款 / 员工代存与转台 ---------------- */

type MemberBindPhoneReq struct {
	g.Meta `path:"/member/bindPhone" method:"post" tags:"酒馆-会员" summary:"绑定手机号（存酒凭证前置）"`
	tavernin.BindPhoneInp
}

type MemberBindPhoneRes struct {
	*tavernin.MemberInfoModel
}

type OrderRefundApplyReq struct {
	g.Meta `path:"/order/refund/apply" method:"post" tags:"酒馆-订单" summary:"申请退款（按明细行）"`
	tavernin.RefundApplyInp
}

type OrderRefundApplyRes struct {
	*tavernin.OrderModel
}

type OrderRefundAuditReq struct {
	g.Meta `path:"/order/refund/audit" method:"post" tags:"酒馆-订单" summary:"退款审核（员工）"`
	tavernin.RefundAuditInp
}

type OrderRefundAuditRes struct {
	*tavernin.OrderModel
}

type StaffDepositRegisterReq struct {
	g.Meta `path:"/staff/bottle/deposit" method:"post" tags:"酒馆-员工" summary:"代存登记（自带酒水）"`
	tavernin.DepositRegisterInp
}

type StaffDepositRegisterRes struct {
	*tavernin.BottleModel
}

type SessionCallReq struct {
	g.Meta `path:"/session/call" method:"post" tags:"酒馆-桌台" summary:"呼叫服务员/催单"`
	tavernin.SessionCallInp
}

type SessionCallRes struct{}

type StaffHandleCallReq struct {
	g.Meta `path:"/staff/call/handle" method:"post" tags:"酒馆-员工" summary:"响应呼叫"`
	tavernin.SessionCallInp
}

type StaffHandleCallRes struct{}

type StaffMoveTableReq struct {
	g.Meta `path:"/staff/table/move" method:"post" tags:"酒馆-员工" summary:"转台/并台"`
	tavernin.MoveTableInp
}

type StaffMoveTableRes struct{}
