package tavern

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"

	"hotgo/internal/dao"
	"hotgo/internal/model/input/tavernin"
	"hotgo/internal/service"
)

// CallMark 呼叫标记（写入会话备注，员工端看板置顶提示）
const CallMark = "【呼叫】"

// Call 呼叫服务员 / 催单
func (s *sTavernSession) Call(ctx context.Context, in *tavernin.SessionCallInp) (err error) {
	if _, err = service.TavernSession().CurrentSessionId(ctx, in.SessionId); err != nil {
		return err
	}

	text := CallMark + "呼叫服务员"
	if in.Type == "urge" {
		text = CallMark + "催单"
	}
	if in.Remark != "" {
		text += "：" + in.Remark
	}
	if in.OrderNo != "" {
		text += "(" + in.OrderNo + ")"
	}

	cols := dao.TavernTableSession.Columns()
	_, err = dao.TavernTableSession.Ctx(ctx).WherePri(in.SessionId).Data(g.Map{
		cols.Remark:    text + " @" + gtime.Now().Format("H:i"),
		cols.UpdatedAt: gtime.Now(),
	}).Update()
	return
}

// HandleCall 响应呼叫（员工确认后清除标记）
func (s *sTavernStaff) HandleCall(ctx context.Context, in *tavernin.SessionCallInp) (err error) {
	if _, err = s.requireStaff(ctx); err != nil {
		return err
	}
	session, err := getSessionById(ctx, in.SessionId)
	if err != nil {
		return err
	}
	if session.IsEmpty() {
		return gerror.New("桌台会话不存在")
	}

	cols := dao.TavernTableSession.Columns()
	remark := session[cols.Remark].String()
	if !gstr.Contains(remark, CallMark) {
		return gerror.New("该桌没有待响应的呼叫")
	}
	_, err = dao.TavernTableSession.Ctx(ctx).WherePri(in.SessionId).Data(g.Map{
		cols.Remark:    gstr.Replace(remark, CallMark, "【已响应】"),
		cols.UpdatedAt: gtime.Now(),
	}).Update()
	return
}
