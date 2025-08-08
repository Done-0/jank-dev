// Package plugin 提供gRPC插件通信实现
// 创建者：Done-0
// 创建时间：2025-08-05
package plugin

import (
	"context"

	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"

	pb "github.com/Done-0/jank/pkg/plugin/proto"
)

// GRPCPlugin gRPC 插件实现
type GRPCPlugin struct {
	plugin.Plugin
	Impl Plugin
}

// NewGRPCPlugin 创建 gRPC 插件
func NewGRPCPlugin(impl Plugin) plugin.Plugin {
	return &GRPCPlugin{Impl: impl}
}

// GRPCServer 创建 gRPC 服务端
func (p *GRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	pb.RegisterPluginServiceServer(s, &grpcServer{Impl: p.Impl})
	return nil
}

// GRPCClient 创建 gRPC 客户端
func (p *GRPCPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (any, error) {
	return &grpcClient{client: pb.NewPluginServiceClient(c)}, nil
}

// grpcServer gRPC 服务端实现
type grpcServer struct {
	pb.UnimplementedPluginServiceServer
	Impl Plugin
}

// Execute 执行插件方法
func (s *grpcServer) Execute(ctx context.Context, req *pb.ExecuteRequest) (*pb.ExecuteResponse, error) {
	data, err := s.Impl.Execute(ctx, req.Method, req.Args)
	if err != nil {
		return nil, err
	}
	return &pb.ExecuteResponse{Data: data}, nil
}

// HealthCheck 检查插件健康状态
func (s *grpcServer) HealthCheck(ctx context.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	err := s.Impl.HealthCheck(ctx)
	if err != nil {
		return &pb.HealthCheckResponse{Status: "unhealthy"}, err
	}
	return &pb.HealthCheckResponse{Status: "healthy"}, nil
}

// grpcClient gRPC 客户端实现
type grpcClient struct {
	client pb.PluginServiceClient
}

// Execute 执行插件方法
func (c *grpcClient) Execute(ctx context.Context, method string, args map[string]string) (map[string]string, error) {
	resp, err := c.client.Execute(ctx, &pb.ExecuteRequest{
		Method: method,
		Args:   args,
	})
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// HealthCheck 检查插件健康状态
func (c *grpcClient) HealthCheck(ctx context.Context) error {
	_, err := c.client.HealthCheck(ctx, &pb.HealthCheckRequest{})
	return err
}
