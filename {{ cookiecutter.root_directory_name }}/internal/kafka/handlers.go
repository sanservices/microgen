package kafka

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/sanservices/kit/kafkalistener"
)

// nolint:all  // Example on how to create a new route
func (k *Kafka) processThingUpdate(msg *message.Message) error {
	thingDecoded := Thing{}

	err := kafkalistener.DecodePayload(TopicUpdateThings, msg.Payload, &thingDecoded)
	if err != nil {
		return err
	}

	//	TODO: Do something with the thing

	return nil
}
