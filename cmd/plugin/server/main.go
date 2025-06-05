package main

import (
	"jank.com/jank_blog/internal/rpc"
	"log"
)

func main() {
	server := rpc.NewServer()

	log.Println("启动插件市场RPC服务器...")
	rpcPort := "8080" // TODO 写入配置
	if err := server.Start(rpcPort); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
