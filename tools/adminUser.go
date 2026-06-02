package tools

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// FindPageAdminUserArgs 管理员分页查询参数
type FindPageAdminUserArgs struct {
	Current      string `json:"current,omitempty"`
	PageSize     string `json:"pageSize,omitempty"`
	UserName     string `json:"user_name,omitempty"`
	UserNick     string `json:"user_nick,omitempty"`
	AccountType  string `json:"account_type,omitempty"`
	Mobile       string `json:"mobile,omitempty"`
	DepartmentId string `json:"department_id,omitempty"`
	RoleIds      string `json:"role_ids,omitempty"`
	Status       string `json:"status,omitempty"`
	Remarks      string `json:"remarks,omitempty"`
	MerchantCode string `json:"merchant_code,omitempty"`
	CreateBy     string `json:"create_by,omitempty"`
	UpdateBy     string `json:"update_by,omitempty"`
}

func (m *MCP) findPageAdminUser(tool *server.MCPServer) {
	userTool := mcp.NewTool("find_page_admin_user",
		mcp.WithDescription("管理员分页查询"),
		mcp.WithString("current", mcp.Description("页码，默认1")),
		mcp.WithString("pageSize", mcp.Description("每页条数，默认10")),
		mcp.WithString("user_name", mcp.Description("用户名")),
		mcp.WithString("user_nick", mcp.Description("用户昵称")),
		mcp.WithString("account_type", mcp.Description("账号类型：1超级管理员 2商户管理员 3普通账号")),
		mcp.WithString("mobile", mcp.Description("手机号")),
		mcp.WithString("department_id", mcp.Description("部门ID")),
		mcp.WithString("role_ids", mcp.Description("角色ID集合")),
		mcp.WithString("status", mcp.Description("状态：1启用 0停用")),
		mcp.WithString("remarks", mcp.Description("备注")),
		mcp.WithString("merchant_code", mcp.Description("商户编码")),
		mcp.WithString("create_by", mcp.Description("创建人")),
		mcp.WithString("update_by", mcp.Description("修改人")),
	)
	tool.AddTool(userTool, mcp.NewTypedToolHandler(func(ctx context.Context, request mcp.CallToolRequest, args FindPageAdminUserArgs) (*mcp.CallToolResult, error) {
		params := make(map[string]string)
		if args.Current != "" {
			params["current"] = args.Current
		}
		if args.PageSize != "" {
			params["pageSize"] = args.PageSize
		}
		if args.UserName != "" {
			params["user_name"] = args.UserName
		}
		if args.UserNick != "" {
			params["user_nick"] = args.UserNick
		}
		if args.AccountType != "" {
			params["account_type"] = args.AccountType
		}
		if args.Mobile != "" {
			params["mobile"] = args.Mobile
		}
		if args.DepartmentId != "" {
			params["department_id"] = args.DepartmentId
		}
		if args.RoleIds != "" {
			params["role_ids"] = args.RoleIds
		}
		if args.Status != "" {
			params["status"] = args.Status
		}
		if args.Remarks != "" {
			params["remarks"] = args.Remarks
		}
		if args.MerchantCode != "" {
			params["merchant_code"] = args.MerchantCode
		}
		if args.CreateBy != "" {
			params["create_by"] = args.CreateBy
		}
		if args.UpdateBy != "" {
			params["update_by"] = args.UpdateBy
		}
		result, err := m.h.FindPageAdminUser(params)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("查询失败: %v", err)), nil
		}
		return mcp.NewToolResultText(result), nil
	}))
}
