package main

import (
	"clear/config"
	"clear/global"
	"clear/tools"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/mark3labs/mcp-go/server"
	"gopkg.in/yaml.v3"
)

func main() {
	global.Config = config.Load()
	if global.Config.URL == "" {
		file, err := ioutil.ReadFile("config.yaml")
		if err != nil {
			log.Fatal(err)
		}
		err = yaml.Unmarshal(file, &global.Config)
		if err != nil {
			log.Fatal(err)
		}
	}
	t := tools.New().MCPServer
	//if err := server.ServeStdio(t); err != nil {
	//	fmt.Printf("Server error: %v\n", err)
	//}
	log.Printf("SSE endpoint: http://127.0.0.1:28080/sse")
	if err := server.NewSSEServer(t).Start(":28080"); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
