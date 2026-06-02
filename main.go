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
	t := tools.New().MCPServer
	if err := server.ServeStdio(t); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
	//sseServer := server.NewSSEServer(t)
	//log.Printf("SSE endpoint: http://127.0.0.1:28080/sse")
	//if err := sseServer.Start(":28080"); err != nil {
	//	fmt.Printf("Server error: %v\n", err)
	//}
}
