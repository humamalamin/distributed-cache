package main

import (
	"context"
	"log"
	"sync"

	"github.com/humamalamin/distributed-cache/cache"
	"github.com/humamalamin/distributed-cache/cluster"
	"github.com/humamalamin/distributed-cache/server"
)

func main() {
	c := cache.NewCache(100)
	rds := cluster.InitRedis(c)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		restServer := server.NewServer(c)
		log.Println("Starting REST API server on port 8080")
		restServer.Start("8080")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("Starting gRPC server on port 50051")
		cluster.StartGRPCServer("50051", c)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("Starting Redis Subcriber")
		rds.StartRedisSubscriber(context.Background())
	}()

	wg.Wait()
}
