package kafka

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/sanservices/kit/kafkalistener"
)

func (k *Kafka) processThingUpdate(msg *message.Message) error {
	ctx := msg.Context()
	thingDecoded := Thing{}

	err := kafkalistener.DecodePayload(TopicUpdateThings, msg.Payload, &thingDecoded)
	if err != nil {
		return err
	}

	//	TODO: Do something with the thing

	return nil
}
