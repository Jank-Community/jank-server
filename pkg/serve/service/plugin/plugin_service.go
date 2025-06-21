package service

import (
	"context"
	"fmt"
	"jank.com/jank_blog/internal/plugin_rpc"
	"log"
	"time"
)

func CallPlugin(ctx context.Context, host *plugin_rpc.PluginHost) {
	{
		ticker := time.NewTicker(10 * time.Second) // 每10秒尝试调用
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				if len(host.LoadedPlugins) > 0 {
					// 获取第一个连接的插件ID（仅作示例，实际应有更智能的选择逻辑）
					var firstPluginID string
					for id := range host.LoadedPlugins {
						firstPluginID = id
						break
					}
					if firstPluginID != "" {
						payload := map[string]interface{}{
							"input": fmt.Sprintf("Hello from Host! Time: %s", time.Now().Format(time.RFC3339)),
						}
						result, err := host.CallPluginExecute(ctx, firstPluginID, "processData", payload)
						if err != nil {
							log.Printf("调用插件 %s Execute 失败: %v", firstPluginID, err)
						} else {
							log.Printf("调用插件 %s Execute 成功，结果: %s", firstPluginID, result)
						}

						// 尝试调用流式接口
						//streamMessages := []string{
						//	"Stream Msg 1",
						//	"Stream Msg 2",
						//	"Stream Msg 3",
						//}
						//err = host.CallPluginStreamData(ctx, firstPluginID, streamMessages)
						//if err != nil {
						//	log.Printf("调用插件 %s StreamData 失败: %v", firstPluginID, err)
						//}
					}
				}
			}
		}
	}
}
