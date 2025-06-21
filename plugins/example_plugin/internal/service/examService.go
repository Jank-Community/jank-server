package service

import (
	"context"
	"encoding/json"
	"example_plugin/internal/proto"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"time"
)

// ExamplePluginServer 实现了 IPluginService
type ExamplePluginServer struct {
	plugin_market.UnimplementedIPluginServiceServer
	pluginID string
	// 宿主反向调用服务的客户端（可选，如果插件需要调用宿主功能）
	hostClient plugin_market.HostServiceClient
}

func NewExamplePluginServer(id string) *ExamplePluginServer {
	return &ExamplePluginServer{
		pluginID: id,
	}
}

// Execute 插件执行逻辑
func (s *ExamplePluginServer) Execute(ctx context.Context, req *plugin_market.ExecuteRequest) (*plugin_market.ExecuteResponse, error) {
	log.Printf("插件 %s 收到 Execute 请求: Method=%s, Payload=%s", s.pluginID, req.Method, req.Payload)

	switch req.Method {
	case "processData":
		// 模拟处理数据
		var data map[string]interface{}
		err := json.Unmarshal([]byte(req.Payload), &data)
		if err != nil {
			return &plugin_market.ExecuteResponse{Success: false, ErrorMessage: fmt.Sprintf("无效的 payload: %v", err)}, nil
		}
		input, ok := data["input"].(string)
		if !ok {
			return &plugin_market.ExecuteResponse{Success: false, ErrorMessage: "payload 中缺少 'input' 字段或类型不正确"}, nil
		}
		result := fmt.Sprintf("Plugin %s processed: '%s' at %s", s.pluginID, input, time.Now().Format(time.RFC3339))
		return &plugin_market.ExecuteResponse{Success: true, Result: result}, nil
	case "greet":
		return &plugin_market.ExecuteResponse{Success: true, Result: fmt.Sprintf("Hello from Plugin %s!", s.pluginID)}, nil
	default:
		return &plugin_market.ExecuteResponse{Success: false, ErrorMessage: fmt.Sprintf("未知方法: %s", req.Method)}, status.Error(codes.NotFound, "Method not found")
	}
}

// StreamData 双向流式通信
func (s *ExamplePluginServer) StreamData(stream plugin_market.IPluginService_StreamDataServer) error {
	log.Printf("插件 %s 收到 StreamData 连接", s.pluginID)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Printf("插件 %s 接收流数据完毕", s.pluginID)
			return nil
		}
		if err != nil {
			log.Printf("插件 %s 接收流数据失败: %v", s.pluginID, err)
			return err
		}
		log.Printf("插件 %s 收到流数据: [Seq:%d] %s", s.pluginID, req.SequenceNum, req.Data)

		// 模拟处理后发送响应
		responseMsg := fmt.Sprintf("Plugin %s processed stream: '%s'", s.pluginID, req.Data)
		if err := stream.Send(&plugin_market.StreamDataResponse{
			ResponseData: responseMsg,
			SequenceNum:  req.SequenceNum,
		}); err != nil {
			log.Printf("插件 %s 发送流数据响应失败: %v", s.pluginID, err)
			return err
		}
	}
}
