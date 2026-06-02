#!/bin/bash

# MCP 项目打包脚本
# 支持跨平台编译为 Windows exe

set -e

APP_NAME="mcp-server"
BUILD_DIR="build"

# 颜色输出
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}  MCP Server 打包脚本${NC}"
echo -e "${GREEN}========================================${NC}"

# 清理旧构建
rm -rf ${BUILD_DIR}
mkdir -p ${BUILD_DIR}

# 编译当前系统版本
echo -e "${YELLOW}[1/4] 编译当前系统版本...${NC}"
go build -ldflags="-s -w" -o ${BUILD_DIR}/${APP_NAME} .
echo -e "${GREEN}  -> ${BUILD_DIR}/${APP_NAME}${NC}"

# 编译 Windows amd64 exe
echo -e "${YELLOW}[2/4] 编译 Windows amd64...${NC}"
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o ${BUILD_DIR}/${APP_NAME}-windows-amd64.exe .
echo -e "${GREEN}  -> ${BUILD_DIR}/${APP_NAME}-windows-amd64.exe${NC}"

# 编译 Linux amd64
echo -e "${YELLOW}[3/4] 编译 Linux amd64...${NC}"
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ${BUILD_DIR}/${APP_NAME}-linux-amd64 .
echo -e "${GREEN}  -> ${BUILD_DIR}/${APP_NAME}-linux-amd64${NC}"

# 编译 macOS arm64 (Apple Silicon)
echo -e "${YELLOW}[4/4] 编译 macOS arm64...${NC}"
GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o ${BUILD_DIR}/${APP_NAME}-darwin-arm64 .
echo -e "${GREEN}  -> ${BUILD_DIR}/${APP_NAME}-darwin-arm64${NC}"

echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}  打包完成！${NC}"
echo -e "${GREEN}  输出目录: ${BUILD_DIR}/${NC}"
echo -e "${GREEN}========================================${NC}"
ls -lh ${BUILD_DIR}/
