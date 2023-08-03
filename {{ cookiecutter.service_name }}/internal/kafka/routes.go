package kafka

import (
	"context"
	"log"

	"github.com/sanservices/kit/kafkalistener"
)

func (k *Kafka) routes() []kafkalistener.RouteHandler {
	return []kafkalistener.RouteHandler{
		{},
	}
}

func (k *Kafka) StartListener(ctx context.Context) {
	err := k.mb.Listen(ctx, k.routes())
	if err != nil {
		log.Fatalln(err)
	}
}
