package cluster

import (
	"context"
	"log"
	"net"

	"github.com/humamalamin/distributed-cache/cache"
	pb "github.com/humamalamin/distributed-cache/proto"
	"google.golang.org/grpc"
)

type Node struct {
	pb.UnimplementedCacheServiceServer
	Cache *cache.Cache
}

func (n *Node) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	value, found := n.Cache.Get(req.Key)
	return &pb.GetResponse{Value: value, Found: found}, nil
}

func (n *Node) Set(ctx context.Context, req *pb.SetRequest) (*pb.SetResponse, error) {
	n.Cache.Set(req.Key, req.Value)
	return &pb.SetResponse{Success: true}, nil
}

func StartGRPCServer(port string, c *cache.Cache) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCacheServiceServer(grpcServer, &Node{Cache: c})
	log.Printf("gRPC server listening on %s", port)
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
