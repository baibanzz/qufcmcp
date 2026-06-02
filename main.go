package main

import (
	"clear/config"
	"clear/global"
	"clear/tools"
	"fmt"

	"github.com/mark3labs/mcp-go/server"
)

func main() {
	global.Config = config.Load()
	if err := server.ServeStdio(tools.New().MCPServer); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
