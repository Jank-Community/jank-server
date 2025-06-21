package plugin_rpc

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "jank.com/jank_blog/pkg/proto/plugin_market"
	"log"
	"sync"
	"time"
)

const (
	heartbeatInterval = 30 * time.Second // 插件心跳间隔
	pluginTimeout     = 60 * time.Second // 插件超时时间 (超过此时间未收到心跳则标记为错误或停止)
)

// pluginEntry 存储插件的运行时信息
type pluginEntry struct {
	pb.PluginInfo
	LastHeartbeat time.Time
}

// PluginManagerServer 实现了 PluginManagerService
type PluginManagerServer struct {
	pb.UnimplementedPluginManagerServiceServer
	plugins sync.Map   // map[string]*pluginEntry, key is plugin ID
	mu      sync.Mutex // For protecting map operations if not using sync.Map
}

func NewPluginManagerServer() *PluginManagerServer {
	s := &PluginManagerServer{}
	go s.monitorPlugins() // 启动插件监控协程
	return s
}

// RegisterPlugin 插件注册
func (s *PluginManagerServer) RegisterPlugin(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	pluginInfo := req.GetPluginInfo()
	if pluginInfo.Id == "" || pluginInfo.Name == "" || pluginInfo.Address == "" {
		return &pb.RegisterResponse{Success: false, Message: "插件ID、名称和地址不能为空"}, status.Error(codes.InvalidArgument, "Invalid plugin info")
	}

	entry := &pluginEntry{
		PluginInfo:    *pluginInfo,
		LastHeartbeat: time.Now(),
	}
	entry.Status = pb.PluginStatus_REGISTERED // 初始状态

	if _, loaded := s.plugins.LoadOrStore(pluginInfo.Id, entry); loaded {
		// 如果插件已存在，更新信息
		log.Printf("插件已存在，更新信息: %s", pluginInfo.Id)
		s.plugins.Store(pluginInfo.Id, entry) // 强制更新
		return &pb.RegisterResponse{Success: true, Message: "插件信息已更新"}, nil
	}

	log.Printf("新插件注册: ID=%s, Name=%s, Address=%s", pluginInfo.Id, pluginInfo.Name, pluginInfo.Address)
	return &pb.RegisterResponse{Success: true, Message: "插件注册成功"}, nil
}

// UnregisterPlugin 插件注销
func (s *PluginManagerServer) UnregisterPlugin(ctx context.Context, req *pb.UnregisterRequest) (*pb.UnregisterResponse, error) {
	pluginID := req.GetPluginId()
	if pluginID == "" {
		return &pb.UnregisterResponse{Success: false, Message: "插件ID不能为空"}, status.Error(codes.InvalidArgument, "Plugin ID cannot be empty")
	}

	if _, ok := s.plugins.LoadAndDelete(pluginID); !ok {
		return &pb.UnregisterResponse{Success: false, Message: "插件未找到或已注销"}, status.Error(codes.NotFound, "Plugin not found")
	}

	log.Printf("插件已注销: ID=%s", pluginID)
	return &pb.UnregisterResponse{Success: true, Message: "插件注销成功"}, nil
}

// SendHeartbeat 接收插件心跳
func (s *PluginManagerServer) SendHeartbeat(ctx context.Context, req *pb.HeartbeatRequest) (*pb.HeartbeatResponse, error) {
	pluginID := req.GetPluginId()
	if pluginID == "" {
		return &pb.HeartbeatResponse{Success: false}, status.Error(codes.InvalidArgument, "Plugin ID cannot be empty")
	}

	if val, ok := s.plugins.Load(pluginID); ok {
		entry := val.(*pluginEntry)
		entry.LastHeartbeat = time.Now()
		// 插件状态更新，如果插件明确发送了状态
		if req.CurrentStatus != pb.PluginStatus_UNKNOWN {
			entry.Status = req.CurrentStatus
		} else if entry.Status != pb.PluginStatus_RUNNING && entry.Status != pb.PluginStatus_DISABLED {
			// 如果插件不是明确的运行中或禁用状态，且心跳正常，则设置为RUNNING
			entry.Status = pb.PluginStatus_RUNNING
		}
		s.plugins.Store(pluginID, entry) // 更新map中的值
		return &pb.HeartbeatResponse{Success: true}, nil
	}

	log.Printf("收到未知插件心跳: %s", pluginID)
	return &pb.HeartbeatResponse{Success: false}, status.Error(codes.NotFound, "Plugin not found for heartbeat")
}

// GetPlugin 获取单个插件信息
func (s *PluginManagerServer) GetPlugin(ctx context.Context, req *pb.GetPluginRequest) (*pb.GetPluginResponse, error) {
	pluginID := req.GetPluginId()
	if pluginID == "" {
		return nil, status.Error(codes.InvalidArgument, "Plugin ID cannot be empty")
	}

	if val, ok := s.plugins.Load(pluginID); ok {
		entry := val.(*pluginEntry)
		return &pb.GetPluginResponse{PluginInfo: &entry.PluginInfo}, nil
	}
	return nil, status.Error(codes.NotFound, "Plugin not found")
}

// ListPlugins 列出所有插件信息
func (s *PluginManagerServer) ListPlugins(ctx context.Context, req *pb.ListPluginsRequest) (*pb.ListPluginsResponse, error) {
	var pluginInfos []*pb.PluginInfo
	filterStatus := req.GetStatusFilter()

	s.plugins.Range(func(key, value interface{}) bool {
		entry := value.(*pluginEntry)
		if filterStatus == pb.PluginStatus_UNKNOWN || entry.Status == filterStatus {
			pluginInfos = append(pluginInfos, &entry.PluginInfo)
		}
		return true
	})
	return &pb.ListPluginsResponse{Plugins: pluginInfos}, nil
}

// UpdatePluginStatus 更新插件状态（例如宿主禁用/启用）
func (s *PluginManagerServer) UpdatePluginStatus(ctx context.Context, req *pb.UpdatePluginStatusRequest) (*pb.UpdatePluginStatusResponse, error) {
	pluginID := req.GetPluginId()
	newStatus := req.GetNewStatus()

	if pluginID == "" {
		return &pb.UpdatePluginStatusResponse{Success: false, Message: "插件ID不能为空"}, status.Error(codes.InvalidArgument, "Plugin ID cannot be empty")
	}

	if val, ok := s.plugins.Load(pluginID); ok {
		entry := val.(*pluginEntry)
		if entry.Status != newStatus { // 只有当状态发生变化时才更新
			entry.Status = newStatus
			s.plugins.Store(pluginID, entry)
			log.Printf("插件状态已更新: ID=%s, NewStatus=%s", pluginID, newStatus.String())
		}
		return &pb.UpdatePluginStatusResponse{Success: true, Message: "插件状态更新成功"}, nil
	}
	return &pb.UpdatePluginStatusResponse{Success: false, Message: "插件未找到"}, status.Error(codes.NotFound, "Plugin not found")
}

// monitorPlugins 定期检查插件心跳，更新状态
func (s *PluginManagerServer) monitorPlugins() {
	ticker := time.NewTicker(heartbeatInterval / 2) // 频率稍高于心跳间隔
	defer ticker.Stop()

	for range ticker.C {
		s.plugins.Range(func(key, value interface{}) bool {
			entry := value.(*pluginEntry)
			// 如果插件状态不是DISABLED，且长时间未收到心跳，则标记为错误
			if entry.Status != pb.PluginStatus_DISABLED && time.Since(entry.LastHeartbeat) > pluginTimeout {
				if entry.Status != pb.PluginStatus_ERROR {
					log.Printf("插件超时无心跳，标记为错误: ID=%s, Name=%s", entry.Id, entry.Name)
					entry.Status = pb.PluginStatus_ERROR
					s.plugins.Store(entry.Id, entry) // 更新map
				}
			}
			return true
		})
	}
}
