package rpc

import (
	"context"
	pb "jank.com/jank_blog/pkg/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Client RPC客户端
type Client struct {
	conn   *grpc.ClientConn
	client pb.PluginServiceClient
}

// NewClient 创建RPC客户端
func NewClient(address string) (*Client, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pb.NewPluginServiceClient(conn)
	return &Client{
		conn:   conn,
		client: client,
	}, nil
}

// Close 关闭连接
func (c *Client) Close() error {
	return c.conn.Close()
}

// ListPlugins 列出插件
func (c *Client) ListPlugins(category string, tags []string, searchQuery string, page, pageSize int32, sortBy string, ascending bool) (*pb.ListPluginsResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return c.client.ListPlugins(ctx, &pb.ListPluginsRequest{
		Category:    category,
		Tags:        tags,
		SearchQuery: searchQuery,
		Page:        page,
		PageSize:    pageSize,
		SortBy:      sortBy,
		Ascending:   ascending,
	})
}

// GetPlugin 获取插件
func (c *Client) GetPlugin(id string) (*pb.GetPluginResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return c.client.GetPlugin(ctx, &pb.GetPluginRequest{Id: id})
}

// RegisterPlugin 注册插件
func (c *Client) RegisterPlugin(plugin *pb.Plugin) (*pb.RegisterPluginResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return c.client.RegisterPlugin(ctx, &pb.RegisterPluginRequest{Plugin: plugin})
}

// UpdatePlugin 更新插件
func (c *Client) UpdatePlugin(plugin *pb.Plugin) (*pb.UpdatePluginResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return c.client.UpdatePlugin(ctx, &pb.UpdatePluginRequest{Plugin: plugin})
}

// DeletePlugin 删除插件
func (c *Client) DeletePlugin(id string) (*pb.DeletePluginResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return c.client.DeletePlugin(ctx, &pb.DeletePluginRequest{Id: id})
}

// DownloadPlugin 下载插件
func (c *Client) DownloadPlugin(id string) (*pb.DownloadPluginResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return c.client.DownloadPlugin(ctx, &pb.DownloadPluginRequest{Id: id})
}

// SearchPlugins 搜索插件
func (c *Client) SearchPlugins(query string, limit int32) (*pb.SearchPluginsResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return c.client.SearchPlugins(ctx, &pb.SearchPluginsRequest{
		Query: query,
		Limit: limit,
	})
}
