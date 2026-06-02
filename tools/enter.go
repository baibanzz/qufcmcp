package tools

import (
	"clear/api"
	"clear/global"

	"github.com/mark3labs/mcp-go/server"
)

type MCP struct {
	*server.MCPServer
	h *api.HTTP
}

func New() *MCP {
	// Create MCP server
	s := server.NewMCPServer(
		"Apifox Import",
		"1.0.0",
		server.WithToolCapabilities(false),
	)
	mcp := &MCP{
		MCPServer: s,
		h:         api.New(global.Config.URL, global.Config.Token),
	}
	mcp.router()
	return mcp
}

func (m *MCP) router() {
	//登录工具
	m.baseLogin(m.MCPServer)
	//用户管理
	m.userRepeat(m.MCPServer)
	m.findPageUser(m.MCPServer)
	m.saveRepeatUser(m.MCPServer)
	//推荐关系
	m.updateReferral(m.MCPServer)

}
