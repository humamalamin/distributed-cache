module github.com/humamalamin/distributed-cache

go 1.22

toolchain go1.22.10

require google.golang.org/protobuf v1.36.0

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/text v0.19.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241015192408-796eee8c2d53 // indirect
)

require (
	github.com/go-redis/redis/v8 v8.11.5
	github.com/google/go-cmp v0.6.0 // indirect
	google.golang.org/grpc v1.69.0
)
