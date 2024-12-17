package cluster

import (
	"context"
	"log"
	"time"

	pb "github.com/humamalamin/distributed-cache/proto"
	"google.golang.org/grpc"
)

func ConnectToNode(address string) pb.CacheServiceClient {
	conn, err := grpc.NewClient(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	return pb.NewCacheServiceClient(conn)
}

func SetKey(client pb.CacheServiceClient, key, value string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := client.Set(ctx, &pb.SetRequest{Key: key, Value: value})
	if err != nil {
		log.Printf("Could not set key: %v", err)
	}
}

func GetKey(client pb.CacheServiceClient, key string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.Get(ctx, &pb.GetRequest{Key: key})
	if err != nil {
		log.Printf("Could not get key: %v", err)
		return
	}
	if res.Found {
		log.Printf("Key: %s, Value: %s", key, res.Value)
	} else {
		log.Printf("Key not found")
	}
}
