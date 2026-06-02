package tools

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// UserRepeatArgs 用户查重查询参数
type UserRepeatArgs struct {
	Current        string `json:"current,omitempty"`
	PageSize       string `json:"pageSize,omitempty"`
	UserId         string `json:"user_id,omitempty"`
	UserName       string `json:"user_name,omitempty"`
	Tel            string `json:"tel,omitempty"`
	Email          string `json:"email,omitempty"`
	Qq             string `json:"qq,omitempty"`
	Wx             string `json:"wx,omitempty"`
	Address        string `json:"address,omitempty"`
	RegisterIp     string `json:"register_ip,omitempty"`
	LastLoginIp    string `json:"last_login_ip,omitempty"`
	RegistVisitorId string `json:"regist_visitor_id,omitempty"`
	MerchantCode   string `json:"merchant_code,omitempty"`
	UserType       string `json:"userType,omitempty"`
}

// SaveRepeatUserArgs 保存用户信息参数
type SaveRepeatUserArgs struct {
	UserId           string  `json:"user_id"`
	MerchantCode     *string `json:"merchant_code,omitempty"`
	LoginStatus      *string `json:"login_status,omitempty"`
	IsWithdraw       *string `json:"is_withdraw,omitempty"`
	IsBusinessLock   *string `json:"is_business_lock,omitempty"`
	IsFutures        *string `json:"is_futures,omitempty"`
	WithdrawStatus   *string `json:"withdraw_status,omitempty"`
	WithdrawMoney    *string `json:"withdraw_money,omitempty"`
	WithdrawFlowLimit *string `json:"withdraw_flow_limit,omitempty"`
	Remark           *string `json:"remark,omitempty"`
	Level            *string `json:"level,omitempty"`
	InviterId        *string `json:"inviter_id,omitempty"`
}

func (m *MCP) userRepeat(tool *server.MCPServer) {
	userTool := mcp.NewTool("user_repeat",
		mcp.WithDescription("用户查重查询（分页查询用户信息）"),
		mcp.WithString("current", mcp.Description("页码，默认1")),
		mcp.WithString("pageSize", mcp.Description("每页条数，默认10")),
		mcp.WithString("user_id", mcp.Description("用户ID")),
		mcp.WithString("user_name", mcp.Description("用户名")),
		mcp.WithString("tel", mcp.Description("手机号")),
		mcp.WithString("email", mcp.Description("邮箱")),
		mcp.WithString("qq", mcp.Description("QQ")),
		mcp.WithString("wx", mcp.Description("微信")),
		mcp.WithString("address", mcp.Description("地址")),
		mcp.WithString("register_ip", mcp.Description("注册IP")),
		mcp.WithString("last_login_ip", mcp.Description("最后登录IP")),
		mcp.WithString("regist_visitor_id", mcp.Description("注册访客ID")),
		mcp.WithString("merchant_code", mcp.Description("商户编码")),
		mcp.WithString("userType", mcp.Description("用户类型，默认-1")),
	)
	tool.AddTool(userTool, mcp.NewTypedToolHandler(func(ctx context.Context, request mcp.CallToolRequest, args UserRepeatArgs) (*mcp.CallToolResult, error) {
		params := make(map[string]string)
		if args.Current != "" {
			params["current"] = args.Current
		}
		if args.PageSize != "" {
			params["pageSize"] = args.PageSize
		}
		if args.UserId != "" {
			params["user_id"] = args.UserId
		}
		if args.UserName != "" {
			params["user_name"] = args.UserName
		}
		if args.Tel != "" {
			params["tel"] = args.Tel
		}
		if args.Email != "" {
			params["email"] = args.Email
		}
		if args.Qq != "" {
			params["qq"] = args.Qq
		}
		if args.Wx != "" {
			params["wx"] = args.Wx
		}
		if args.Address != "" {
			params["address"] = args.Address
		}
		if args.RegisterIp != "" {
			params["register_ip"] = args.RegisterIp
		}
		if args.LastLoginIp != "" {
			params["last_login_ip"] = args.LastLoginIp
		}
		if args.RegistVisitorId != "" {
			params["regist_visitor_id"] = args.RegistVisitorId
		}
		if args.MerchantCode != "" {
			params["merchant_code"] = args.MerchantCode
		}
		if args.UserType != "" {
			params["userType"] = args.UserType
		}
		result, err := m.h.PageUserRepeat(params)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("查询失败: %v", err)), nil
		}
		return mcp.NewToolResultText(result), nil
	}))
}

func (m *MCP) saveRepeatUser(tool *server.MCPServer) {
	saveTool := mcp.NewTool("save_repeat_user",
		mcp.WithDescription("保存/更新用户信息（用户信息维护更新）"),
		mcp.WithString("user_id",
			mcp.Required(),
			mcp.Description("用户ID"),
		),
		mcp.WithString("merchant_code", mcp.Description("商户编码")),
		mcp.WithString("login_status", mcp.Description("登录权限：0正常 1限制登录")),
		mcp.WithString("is_withdraw", mcp.Description("提现权限：1正常 2限制提现")),
		mcp.WithString("is_business_lock", mcp.Description("业务锁定：0正常 1锁定")),
		mcp.WithString("is_futures", mcp.Description("用户交割合约：0正常 1禁止")),
		mcp.WithString("withdraw_status", mcp.Description("提现状态：0默认 1交易量不达标也可提现")),
		mcp.WithString("withdraw_money", mcp.Description("可提现金额限制")),
		mcp.WithString("withdraw_flow_limit", mcp.Description("提现流水限制")),
		mcp.WithString("remark", mcp.Description("备注")),
		mcp.WithString("level", mcp.Description("等级")),
		mcp.WithString("inviter_id", mcp.Description("推荐人ID")),
	)
	tool.AddTool(saveTool, mcp.NewTypedToolHandler(func(ctx context.Context, request mcp.CallToolRequest, args SaveRepeatUserArgs) (*mcp.CallToolResult, error) {
		data := make(map[string]any)
		data["user_id"] = args.UserId
		if args.MerchantCode != nil {
			data["merchant_code"] = *args.MerchantCode
		}
		if args.LoginStatus != nil {
			data["login_status"] = *args.LoginStatus
		}
		if args.IsWithdraw != nil {
			data["is_withdraw"] = *args.IsWithdraw
		}
		if args.IsBusinessLock != nil {
			data["is_business_lock"] = *args.IsBusinessLock
		}
		if args.IsFutures != nil {
			data["is_futures"] = *args.IsFutures
		}
		if args.WithdrawStatus != nil {
			data["withdraw_status"] = *args.WithdrawStatus
		}
		if args.WithdrawMoney != nil {
			data["withdraw_money"] = *args.WithdrawMoney
		}
		if args.WithdrawFlowLimit != nil {
			data["withdraw_flow_limit"] = *args.WithdrawFlowLimit
		}
		if args.Remark != nil {
			data["remark"] = *args.Remark
		}
		if args.Level != nil {
			data["level"] = *args.Level
		}
		if args.InviterId != nil {
			data["inviter_id"] = *args.InviterId
		}
		result, err := m.h.SaveRepeatUser(data)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("保存失败: %v", err)), nil
		}
		return mcp.NewToolResultText(result), nil
	}))
}
