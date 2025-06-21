// plugin/example_plugin/main.go
package main

import (
	"context"
	"example_plugin/internal/proto"
	"example_plugin/internal/service"
	"flag"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const (
	managerAddress = "127.0.0.1:50050"
)

func main() {
	pluginID := flag.String("id", uuid.New().String(), "插件唯一ID")
	pluginName := flag.String("name", "", "插件名称")
	pluginVersion := flag.String("version", "1.0.0", "插件版本")
	pluginAuthor := flag.String("author", "ixue", "插件作者")
	pluginDescription := flag.String("desc", "A simple example plugin written in Go", "插件描述")
	pluginPort := flag.String("port", "", "插件监听端口，例如: :50051")
	flag.Parse()

	if *pluginName == "" {
		log.Fatalf("必须提供插件名称 -name 参数")
	}
	if *pluginPort == "" {
		log.Fatalf("必须提供插件监听端口 -port 参数")
	}

	log.Printf("插件 %s 正在启动...", *pluginName)

	// 1. 启动插件自身的 gRPC 服务
	lis, err := net.Listen("tcp", ":"+*pluginPort)
	if err != nil {
		log.Fatalf("无法监听端口 %s: %v", *pluginPort, err)
	}
	grpcServer := grpc.NewServer()
	pluginServer := service.NewExamplePluginServer(*pluginID)
	plugin_market.RegisterIPluginServiceServer(grpcServer, pluginServer)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Printf("插件 %s (gRPC) 正在监听 :%s", *pluginID, *pluginPort)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("启动插件 gRPC 服务器失败: %v", err)
		}
	}()

	// 2. 连接到插件管理器并注册
	managerConn, err := grpc.Dial(managerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("无法连接到插件管理器: %v", err)
	}
	defer managerConn.Close()
	managerClient := plugin_market.NewPluginManagerServiceClient(managerConn)

	pluginInfo := &plugin_market.PluginInfo{
		Id:          *pluginID,
		Name:        *pluginName,
		Version:     *pluginVersion,
		Author:      *pluginAuthor,
		Description: *pluginDescription,
		Address:     "localhost:" + *pluginPort, // 注意这里是插件自身监听的地址
		Status:      plugin_market.PluginStatus_REGISTERED,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	registerResp, err := managerClient.RegisterPlugin(ctx, &plugin_market.RegisterRequest{PluginInfo: pluginInfo})
	if err != nil {
		log.Fatalf("向管理器注册插件失败: %v", err)
	}
	if !registerResp.Success {
		log.Fatalf("管理器拒绝注册: %s", registerResp.Message)
	}
	log.Printf("插件 %s 成功注册到管理器", *pluginID)

	// 3. 定期发送心跳
	go func() {
		ticker := time.NewTicker(20 * time.Second) // 略小于管理器超时时间
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return // 上下文取消，停止心跳
			case <-ticker.C:
				heartbeatResp, err := managerClient.SendHeartbeat(ctx, &plugin_market.HeartbeatRequest{
					PluginId:      *pluginID,
					CurrentStatus: plugin_market.PluginStatus_RUNNING, // 插件明确表示自己正在运行
				})
				if err != nil {
					log.Printf("发送心跳失败: %v", err)
					// 如果连续失败多次，可以考虑自动注销或退出
				} else if !heartbeatResp.Success {
					log.Printf("管理器拒绝心跳: %v", err)
				}
			}
		}
	}()

	// 监听终止信号，优雅关闭
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	fmt.Printf("插件 %s 正在关闭...\n", *pluginID)

	// 停止心跳 goroutine
	cancel()

	// 向管理器发送注销请求
	unregisterResp, err := managerClient.UnregisterPlugin(context.Background(), &plugin_market.UnregisterRequest{PluginId: *pluginID})
	if err != nil {
		log.Printf("向管理器注销插件失败: %v", err)
	} else if !unregisterResp.Success {
		log.Printf("管理器拒绝注销: %s", unregisterResp.Message)
	} else {
		log.Printf("插件 %s 成功向管理器注销", *pluginID)
	}

	// 停止插件自身的 gRPC 服务
	grpcServer.GracefulStop()
	wg.Wait() // 等待 gRPC 服务 goroutine 退出
	log.Printf("插件 %s 已停止。", *pluginID)
}
