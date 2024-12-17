package cluster

import (
	"context"
	"fmt"
	"log"
)

func (p *Publish) StartRedisSubscriber(ctx context.Context) {
	sub := p.redis.Subscribe(ctx, "cache-channel")
	defer sub.Close()

	for {
		msg, err := sub.ReceiveMessage(ctx)
		if err != nil {
			log.Fatalf("Error while receiving message: %v", err)
		}
		fmt.Printf("Received message: %s\n", msg.Payload)

		var action, key, value string
		_, err = fmt.Sscanf(msg.Payload, "%s %s %s", &action, &key, &value)
		if err != nil {
			log.Printf("Failed to parse message: %v", err)
			continue
		}

		if action == "SET" {
			p.Cache.Set(key, value)
			log.Printf("Replicated SET for key %s with value %s", key, value)
		}
	}
}
