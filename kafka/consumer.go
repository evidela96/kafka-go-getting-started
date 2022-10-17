package kafka

import (
	"fmt"
	"log"

	"github.com/architecture-it/go-platform/AMQStream"
)

type ConsumerTest struct{}

func (c *ConsumerTest) Handler(event interface{}, metadata AMQStream.ConsumerMetadata) error {
	fmt.Printf("Evento: %v\n", event)
	return nil
}

func init() {
	config, err := AMQStream.AddKafka()
	if err != nil {
		log.Fatal(err)
	}
	Consumer := ConsumerTest{}
	config.ToConsumer(&Consumer, &event, topics)
	config.Build()
}
