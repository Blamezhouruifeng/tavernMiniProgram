package tavern

import (
	"context"

	apiTavern "hotgo/api/api/tavern"
	"hotgo/internal/service"
)

func (c *cMember) MemberBindPhone(ctx context.Context, req *apiTavern.MemberBindPhoneReq) (res *apiTavern.MemberBindPhoneRes, err error) {
	info, err := service.TavernMember().BindPhone(ctx, &req.BindPhoneInp)
	if err != nil {
		return
	}
	res = &apiTavern.MemberBindPhoneRes{MemberInfoModel: info}
	return
}

func (c *cOrder) OrderRefundApply(ctx context.Context, req *apiTavern.OrderRefundApplyReq) (res *apiTavern.OrderRefundApplyRes, err error) {
	order, err := service.TavernOrder().ApplyRefund(ctx, &req.RefundApplyInp)
	if err != nil {
		return
	}
	res = &apiTavern.OrderRefundApplyRes{OrderModel: order}
	return
}

func (c *cOrder) OrderRefundAudit(ctx context.Context, req *apiTavern.OrderRefundAuditReq) (res *apiTavern.OrderRefundAuditRes, err error) {
	order, err := service.TavernOrder().AuditRefund(ctx, &req.RefundAuditInp)
	if err != nil {
		return
	}
	res = &apiTavern.OrderRefundAuditRes{OrderModel: order}
	return
}

func (c *cStaff) StaffDepositRegister(ctx context.Context, req *apiTavern.StaffDepositRegisterReq) (res *apiTavern.StaffDepositRegisterRes, err error) {
	bottle, err := service.TavernStaff().DepositRegister(ctx, &req.DepositRegisterInp)
	if err != nil {
		return
	}
	res = &apiTavern.StaffDepositRegisterRes{BottleModel: bottle}
	return
}

func (c *cSession) SessionCall(ctx context.Context, req *apiTavern.SessionCallReq) (res *apiTavern.SessionCallRes, err error) {
	err = service.TavernSession().Call(ctx, &req.SessionCallInp)
	return
}

func (c *cStaff) StaffHandleCall(ctx context.Context, req *apiTavern.StaffHandleCallReq) (res *apiTavern.StaffHandleCallRes, err error) {
	err = service.TavernStaff().HandleCall(ctx, &req.SessionCallInp)
	return
}

func (c *cStaff) StaffMoveTable(ctx context.Context, req *apiTavern.StaffMoveTableReq) (res *apiTavern.StaffMoveTableRes, err error) {
	err = service.TavernStaff().MoveTable(ctx, &req.MoveTableInp)
	return
}
