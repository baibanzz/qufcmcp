package tools

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// LoginArgs 定义登录工具的参数
type LoginArgs struct {
	Account  string  `json:"account"`
	Password string  `json:"password"`
	MFACode  *string `json:"mfa_code,omitempty"`
}

func (m *MCP) baseLogin(tool *server.MCPServer) {
	loginTool := mcp.NewTool("login",
		mcp.WithDescription("登录平台"),
		mcp.WithString("account",
			mcp.Required(),
			mcp.Description("账号"),
		),
		mcp.WithString("password",
			mcp.Required(),
			mcp.Description("密码"),
		),
		mcp.WithString("mfa_code",
			mcp.Description("MFA验证码（如果启用MFA则必填）"),
		),
	)
	tool.AddTool(loginTool, mcp.NewTypedToolHandler(func(ctx context.Context, request mcp.CallToolRequest, args LoginArgs) (*mcp.CallToolResult, error) {
		// 获取 MFA code，如果没有则为空字符串
		mfaCode := ""
		if args.MFACode != nil {
			mfaCode = *args.MFACode
		}

		// 调用登录接口
		err := m.h.BaseLogin(args.Account, args.Password, mfaCode)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("登录失败: %v", err)), nil
		}

		return mcp.NewToolResultText(), nil
	}))
}
