package rpc

import (
	"context"
	"jank.com/jank_blog/internal/model/plugin"
	pb "jank.com/jank_blog/pkg/proto"
	"jank.com/jank_blog/pkg/serve/service/plugin"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Server RPC服务器
type Server struct {
	pb.UnimplementedPluginServiceServer
	pluginService *service.PluginService
}

// NewServer 创建RPC服务器
func NewServer() *Server {
	return &Server{
		pluginService: service.NewPluginService(),
	}
}

// Start 启动服务器
func (s *Server) Start(port string) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPluginServiceServer(grpcServer, s)

	log.Printf("插件市场RPC服务器启动在端口 %s", port)
	return grpcServer.Serve(lis)
}

// ListPlugins 列出插件
func (s *Server) ListPlugins(ctx context.Context, req *pb.ListPluginsRequest) (*pb.ListPluginsResponse, error) {
	plugins, total, err := s.pluginService.ListPlugins(
		req.Category,
		req.Tags,
		req.SearchQuery,
		int(req.Page),
		int(req.PageSize),
		req.SortBy,
		req.Ascending,
	)
	if err != nil {
		return nil, err
	}

	var pbPlugins []*pb.Plugin
	for _, plugin := range plugins {
		pbPlugins = append(pbPlugins, s.modelToProto(plugin))
	}

	return &pb.ListPluginsResponse{
		Plugins:  pbPlugins,
		Total:    int32(total),
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}

// GetPlugin 获取插件
func (s *Server) GetPlugin(ctx context.Context, req *pb.GetPluginRequest) (*pb.GetPluginResponse, error) {
	plugin, err := s.pluginService.GetPlugin(req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetPluginResponse{
		Plugin: s.modelToProto(plugin),
	}, nil
}

// RegisterPlugin 注册插件
func (s *Server) RegisterPlugin(ctx context.Context, req *pb.RegisterPluginRequest) (*pb.RegisterPluginResponse, error) {
	plugin := s.protoToModel(req.Plugin)

	id, err := s.pluginService.RegisterPlugin(plugin)
	if err != nil {
		return &pb.RegisterPluginResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &pb.RegisterPluginResponse{
		Success:  true,
		Message:  "插件注册成功",
		PluginId: id,
	}, nil
}

// UpdatePlugin 更新插件
func (s *Server) UpdatePlugin(ctx context.Context, req *pb.UpdatePluginRequest) (*pb.UpdatePluginResponse, error) {
	plugin := s.protoToModel(req.Plugin)

	err := s.pluginService.UpdatePlugin(plugin)
	if err != nil {
		return &pb.UpdatePluginResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &pb.UpdatePluginResponse{
		Success: true,
		Message: "插件更新成功",
	}, nil
}

// DeletePlugin 删除插件
func (s *Server) DeletePlugin(ctx context.Context, req *pb.DeletePluginRequest) (*pb.DeletePluginResponse, error) {
	err := s.pluginService.DeletePlugin(req.Id)
	if err != nil {
		return &pb.DeletePluginResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &pb.DeletePluginResponse{
		Success: true,
		Message: "插件删除成功",
	}, nil
}

// DownloadPlugin 下载插件
func (s *Server) DownloadPlugin(ctx context.Context, req *pb.DownloadPluginRequest) (*pb.DownloadPluginResponse, error) {
	downloadURL, err := s.pluginService.DownloadPlugin(req.Id)
	if err != nil {
		return &pb.DownloadPluginResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &pb.DownloadPluginResponse{
		Success:     true,
		Message:     "获取下载链接成功",
		DownloadUrl: downloadURL,
	}, nil
}

// SearchPlugins 搜索插件
func (s *Server) SearchPlugins(ctx context.Context, req *pb.SearchPluginsRequest) (*pb.SearchPluginsResponse, error) {
	plugins, err := s.pluginService.SearchPlugins(req.Query, int(req.Limit))
	if err != nil {
		return nil, err
	}

	var pbPlugins []*pb.Plugin
	for _, plugin := range plugins {
		pbPlugins = append(pbPlugins, s.modelToProto(plugin))
	}

	return &pb.SearchPluginsResponse{
		Plugins: pbPlugins,
	}, nil
}

// modelToProto 模型转Proto
func (s *Server) modelToProto(plugin *model.Plugin) *pb.Plugin {
	return &pb.Plugin{
		Id:            plugin.ID,
		Name:          plugin.Name,
		Version:       plugin.Version,
		Description:   plugin.Description,
		Author:        plugin.Author,
		Category:      plugin.Category,
		Tags:          plugin.Tags,
		DownloadUrl:   plugin.DownloadURL,
		DownloadCount: plugin.DownloadCount,
		Rating:        plugin.Rating,
		CreatedAt:     timestamppb.New(plugin.CreatedAt).Seconds,
		UpdatedAt:     timestamppb.New(plugin.UpdatedAt).Seconds,
		IsActive:      plugin.IsActive,
	}
}

// protoToModel Proto转模型
func (s *Server) protoToModel(pbPlugin *pb.Plugin) *model.Plugin {
	return &model.Plugin{
		ID:            pbPlugin.Id,
		Name:          pbPlugin.Name,
		Version:       pbPlugin.Version,
		Description:   pbPlugin.Description,
		Author:        pbPlugin.Author,
		Category:      pbPlugin.Category,
		Tags:          pbPlugin.Tags,
		DownloadURL:   pbPlugin.DownloadUrl,
		DownloadCount: pbPlugin.DownloadCount,
		Rating:        pbPlugin.Rating,
		CreatedAt:     time.Unix(pbPlugin.CreatedAt, 0),
		UpdatedAt:     time.Unix(pbPlugin.UpdatedAt, 0),
		IsActive:      pbPlugin.IsActive,
	}
}
