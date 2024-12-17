
# **Distributed Cache System with Redis Pub/Sub**

This project implements a distributed cache system using **Go** (Golang), which supports **replication** across multiple nodes using **Redis Pub/Sub**. The system allows key-value pairs to be set in the cache and synchronized across different nodes in a distributed environment.

## **Features**
- **Distributed Cache**: Implements a simple cache system with replication.
- **Cache Repliation**: Automatically replicate cache updates across all nodes using **Redis Pub/Sub** or

## **Architecture**

The system consists of the following components:
1. **Cache**: Stores key-value pairs in memory using **LRU Cache** or a simple **map**.
2. **Redis Pub/Sub**: Used to send messages when the cache is updated to keep all nodes synchronized.
3. **API**: Exposes HTTP and gRPC endpoints to interact with the cache (set/get operations).
4. **Multiple Nodes**: All nodes subscribe to Redis Pub/Sub or Kafka topics to receive cache updates and replicate changes.

## **Prerequisites**
- **Go** (Golang) 1.18+
- **Redis** (optional, for Redis Pub/Sub mode)

### **Installation**

#### **1. Install Go Dependencies**
```bash
go mod tidy
```

#### **2. Set Up Redis or Kafka**

- For **Redis**: Download and install Redis from [Redis.io](https://redis.io/download) or use a Redis cloud service.

#### **3. Install Redis Go Client**
```bash
go get github.com/go-redis/redis/v8
```

### **Running the Project**

#### **1. Redis Mode**
To run the distributed cache system with **Redis Pub/Sub** for replication:

1. Start Redis server on your local machine or use a hosted Redis instance.
2. Run the Go server:
   ```bash
   go run main.go --mode=redis
   ```


### **Usage**

#### **1. Set Cache**
- **POST** `/cache`
  - Set a key-value pair in the cache.
  - Request Body (JSON):
    ```json
    {
      "key": "foo",
      "value": "bar"
    }
    ```

#### **2. Get Cache**
- **GET** `/cache/{key}`
  - Retrieve the value for a given key from the cache.

#### **3. gRPC Endpoints**
- **Set**: `POST /v1/cache/set`
- **Get**: `GET /v1/cache/get/{key}`

### **Code Overview**

The project is structured as follows:

```
├── cache/
│   ├── lru.go           # Cache management (LRU or map-based cache)
│   ├── persistence.go   # Snapshot data into file
├── cluster/
│   ├── client.go          # Communication between nodes
│   ├── node.go            # Communication between nodes
│   ├── publish.go         # Redis integration for Pub/Sub
│   ├── subcriber.go       # Redis integration for Pub/Sub
├── proto/
│   └── cache.proto        # gRPC API definitions
├── server/
│   ├── api.go          # Server for HTTP and gRPC API
├── Dockerfile             # Serverless using Dockerfile
├── main.go                # Main entry point to start the servers
├── README.md              # Project documentation
└── go.mod                 # Go modules dependencies
```

### **Cache Implementation**

The cache is implemented using a simple **LRU cache** from the `golang-lru` package or a map-based cache for simplicity.

### **License**
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
