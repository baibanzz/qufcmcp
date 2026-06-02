package tools

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// UpdateReferralArgs 修改推荐关系参数
type UpdateReferralArgs struct {
	UserId          string `json:"user_id"`
	NewReferrerId   string `json:"new_referrer_id"`
	FundingPassword string `json:"funding_password"`
}

func (m *MCP) updateReferral(tool *server.MCPServer) {
	updateTool := mcp.NewTool("update_referral",
		mcp.WithDescription("修改推荐关系（需要资金密码验证）"),
		mcp.WithString("user_id",
			mcp.Required(),
			mcp.Description("用户ID"),
		),
		mcp.WithString("new_referrer_id",
			mcp.Required(),
			mcp.Description("新推荐人ID"),
		),
		mcp.WithString("funding_password",
			mcp.Required(),
			mcp.Description("资金密码"),
		),
	)
	tool.AddTool(updateTool, mcp.NewTypedToolHandler(func(ctx context.Context, request mcp.CallToolRequest, args UpdateReferralArgs) (*mcp.CallToolResult, error) {
		data := map[string]any{
			"user_id":          args.UserId,
			"new_referrer_id":  args.NewReferrerId,
			"funding_password": args.FundingPassword,
		}
		result, err := m.h.UpdateReferral(data)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("修改推荐关系失败: %v", err)), nil
		}
		return mcp.NewToolResultText(result), nil
	}))
}
