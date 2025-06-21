// host/main.go
package service

import (
	"context"
	"fmt"
	"jank.com/jank_blog/internal/plugin_rpc"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// StartPluginHost 启动插件宿主服务
func StartPluginHost() {
	host, err := plugin_rpc.NewPluginHost()
	if err != nil {
		log.Fatalf("初始化插件宿主失败: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 定期列出并连接插件
	go func() {
		ticker := time.NewTicker(5 * time.Second) // 每5秒检查一次
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				err := host.ListAndConnectPlugins(ctx)
				if err != nil {
					log.Printf("周期性连接插件失败: %v", err)
				}
			}
		}
	}()

	// 每隔一段时间调用一个插件的 Execute 方法
	//go func() {
	//	CallPlugin(ctx, host) // 要确保 CallPlugin 内部也监听 ctx.Done()
	//}()

	// 监听终止信号
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	// 等待信号
	<-sigChan
	fmt.Println("宿主服务正在关闭...")
}
