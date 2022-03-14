package kafka

import "github.com/sanservices/kit/kafkalistener"

var (
	TopicUpdateThings *kafkalistener.Topic = &kafkalistener.Topic{
		Name: "my-kafka-topic",
	}
)
