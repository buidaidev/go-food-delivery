package main

import (
	"context"
	"go-food-delivery/pubsub"
	"go-food-delivery/pubsub/localpubsub"
	"log"
)

func main() {
	var localPubSub pubsub.Pubsub = localpubsub.NewPubSub()

	var topic pubsub.Topic = "OrderCreated"

	sub1, _ := localPubSub.Subscribe(context.Background(), topic)
	sub2, _ := localPubSub.Subscribe(context.Background(), topic)

	localPubSub.Publish(context.Background(), topic, pubsub.NewMessage(1))
	localPubSub.Publish(context.Background(), topic, pubsub.NewMessage(2))

	go func() {
		for {
			log.Println("Sub1:", (<-sub1).Data())
		}
	}()

	go func() {
		for {
			log.Println("Sub2:", (<-sub2).Data())
		}
	}()
}
