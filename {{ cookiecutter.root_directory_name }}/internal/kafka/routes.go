package kafka

import (
	"context"
	"log"

	"github.com/sanservices/kit/kafkalistener"
)

// Handlers returns a list of Kafka route handlers.
func (k *Kafka) routes() []kafkalistener.RouteHandler {
	return []kafkalistener.RouteHandler{
		{
			Name:        "update_thing",
			Topic:       TopicUpdateThings,
			HandlerFunc: k.processThingUpdate,
		},
	}
}

func (k *Kafka) StartListener(ctx context.Context) {
	err := k.mb.Listen(ctx, k.routes())
	if err != nil {
		log.Fatalln(err)
	}
}
