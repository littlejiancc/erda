// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: cluster.proto

package client

import (
	context "context"

	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/core/pipeline/cluster/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// ClusterService cluster.proto
	ClusterService() pb.ClusterServiceClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		clusterService: pb.NewClusterServiceClient(cc),
	}
}

type serviceClients struct {
	clusterService pb.ClusterServiceClient
}

func (c *serviceClients) ClusterService() pb.ClusterServiceClient {
	return c.clusterService
}

type clusterServiceWrapper struct {
	client pb.ClusterServiceClient
	opts   []grpc1.CallOption
}

func (s *clusterServiceWrapper) ClusterHook(ctx context.Context, req *pb.ClusterHookRequest) (*pb.ClusterHookResponse, error) {
	return s.client.ClusterHook(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}