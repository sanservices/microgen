package kafka

import "github.com/sanservices/kit/kafkalistener"

// Create a new topic variable of type *kafkalistener.Topic
// Update the Name field to match your desired Kafka topic name

var (
	TopicUpdateThings *kafkalistener.Topic = &kafkalistener.Topic{
		Name: "my-kafka-topic",
	}
)
