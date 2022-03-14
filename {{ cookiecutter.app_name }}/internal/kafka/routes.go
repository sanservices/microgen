package kafka

import (
	"context"

	"github.com/sanservices/kit/kafkalistener"
)

func (k *Kafka) handlers() []kafkalistener.RouteHandler {
	return []kafkalistener.RouteHandler{
		{
			Name:        "update_thing",
			Topic:       TopicUpdateThings,
			HandlerFunc: k.processThingUpdate,
		},
	}
}
