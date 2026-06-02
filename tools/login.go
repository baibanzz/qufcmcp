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
type TokenArgs struct {
	Token string `json:"token"`
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
		token, err := m.h.BaseLogin(args.Account, args.Password, mfaCode)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("登录失败: %v", err)), nil
		}

		return mcp.NewToolResultText(token), nil
	}))
}

func (m *MCP) setToken(tool *server.MCPServer) {
	loginTool := mcp.NewTool("set_token",
		mcp.WithDescription("设置Token"),
		mcp.WithString("token",
			mcp.Required(),
			mcp.Description("直接设置Token"),
		),
	)
	tool.AddTool(loginTool, mcp.NewTypedToolHandler(func(ctx context.Context, request mcp.CallToolRequest, args TokenArgs) (*mcp.CallToolResult, error) {
		// 获取 MFA code，如果没有则为空字符串
		if args.Token != "" {
			return nil, nil
		}

		// 调用登录接口
		m.h.SetToken(args.Token)
		return nil, nil
	}))
}
