package plugin_rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	pb "jank.com/jank_blog/pkg/proto/plugin_market"
	"log"
	"time"
)

const (
	managerAddress = ":50050"
)

// PluginHost 插件宿主结构体
type PluginHost struct {
	managerClient pb.PluginManagerServiceClient
	// 存储已加载的插件客户端连接
	LoadedPlugins map[string]pb.IPluginServiceClient
}

func NewPluginHost() (*PluginHost, error) {
	// 连接到插件管理器
	conn, err := grpc.Dial(managerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("无法连接到插件管理器: %v", err)
	}
	managerClient := pb.NewPluginManagerServiceClient(conn)

	return &PluginHost{
		managerClient: managerClient,
		LoadedPlugins: make(map[string]pb.IPluginServiceClient),
	}, nil
}

// ListAndConnectPlugins 列出并连接到所有可用插件
func (h *PluginHost) ListAndConnectPlugins(ctx context.Context) error {
	resp, err := h.managerClient.ListPlugins(ctx, &pb.ListPluginsRequest{StatusFilter: pb.PluginStatus_RUNNING})
	if err != nil {
		return fmt.Errorf("列出插件失败: %v", err)
	}

	if len(resp.Plugins) == 0 {
		fmt.Println("没有找到任何运行中的插件。")
		return nil
	}

	fmt.Println("发现以下运行中的插件:")
	for _, info := range resp.Plugins {
		fmt.Printf("  - ID: %s, Name: %s, Version: %s, Address: %s\n", info.Id, info.Name, info.Version, info.Address)
		// 尝试连接到插件
		if _, ok := h.LoadedPlugins[info.Id]; ok {
			fmt.Printf("    插件 %s 已经连接。\n", info.Name)
			continue
		}

		pluginConn, err := grpc.Dial(info.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Printf("    无法连接到插件 %s (%s): %v", info.Name, info.Address, err)
			continue
		}
		h.LoadedPlugins[info.Id] = pb.NewIPluginServiceClient(pluginConn)
		fmt.Printf("    成功连接到插件 %s。\n", info.Name)
	}
	return nil
}

// CallPluginExecute 调用插件的 Execute 方法
func (h *PluginHost) CallPluginExecute(ctx context.Context, pluginID, method string, payload map[string]interface{}) (string, error) {
	client, ok := h.LoadedPlugins[pluginID]
	if !ok {
		return "", fmt.Errorf("插件 %s 未加载或连接", pluginID)
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("序列化 payload 失败: %v", err)
	}

	req := &pb.ExecuteRequest{
		Method:  method,
		Payload: string(payloadBytes),
	}

	resp, err := client.Execute(ctx, req)
	if err != nil {
		return "", fmt.Errorf("调用插件 %s 的 %s 方法失败: %v", pluginID, method, err)
	}

	if !resp.Success {
		return "", fmt.Errorf("插件 %s 执行失败: %s", pluginID, resp.ErrorMessage)
	}

	return resp.Result, nil
}

// CallPluginStreamData 调用插件的 StreamData 方法进行双向流式通信
func (h *PluginHost) CallPluginStreamData(ctx context.Context, pluginID string, messages []string) error {
	client, ok := h.LoadedPlugins[pluginID]
	if !ok {
		return fmt.Errorf("插件 %s 未加载或连接", pluginID)
	}

	stream, err := client.StreamData(ctx)
	if err != nil {
		return fmt.Errorf("无法建立插件 %s 的流式连接: %v", pluginID, err)
	}
	defer stream.CloseSend()

	waitc := make(chan struct{})

	// 接收协程
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// 读取完毕
				close(waitc)
				return
			}
			if err != nil {
				log.Printf("从插件 %s 接收流数据失败: %v", pluginID, err)
				close(waitc)
				return
			}
			fmt.Printf("Host收到来自插件 %s 的流响应: [Seq:%d] %s\n", pluginID, in.SequenceNum, in.ResponseData)
		}
	}()

	// 发送数据
	for i, msg := range messages {
		req := &pb.StreamDataRequest{
			Data:        msg,
			SequenceNum: int32(i + 1),
		}
		fmt.Printf("Host发送流数据到插件 %s: [Seq:%d] %s\n", pluginID, req.SequenceNum, req.Data)
		if err := stream.Send(req); err != nil {
			return fmt.Errorf("向插件 %s 发送流数据失败: %v", pluginID, err)
		}
		time.Sleep(500 * time.Millisecond) // 模拟间隔
	}
	<-waitc // 等待接收协程完成
	return nil
}
