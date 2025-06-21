// server/main.go
package service

import (
	"google.golang.org/grpc"
	"jank.com/jank_blog/internal/plugin_rpc"
	pb "jank.com/jank_blog/pkg/proto/plugin_market"
	"log"
	"net"
)

const (
	managerPort = ":50050"
)

// StartPluginServer 启动插件管理器 gRPC 服务器
func StartPluginServer() {
	lis, err := net.Listen("tcp", managerPort)
	if err != nil {
		log.Fatalf("无法监听端口 %s: %v", managerPort, err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPluginManagerServiceServer(grpcServer, plugin_rpc.NewPluginManagerServer())

	log.Printf("插件管理器服务正在监听 %s", managerPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("启动 gRPC 服务器失败: %v", err)
	}
}
