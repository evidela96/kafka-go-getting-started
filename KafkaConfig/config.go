package KafkaConfig

import (
	"fmt"
	"log"
	"os"

	"github.com/architecture-it/go-platform/AMQStream"
	TestEvents "github.com/architecture-it/integracion-schemas-event-go/Test"
)

type ConsumerTest struct{}
type PublisherTest struct{}

func init() {
	config, err := AMQStream.AddKafka()
	if err != nil {
		log.Fatal(err)
	}
	subscriber := ConsumerTest{}
	config.ToConsumer(&subscriber, &TestEvents.KafkaDemoRetro{}, []string{os.Getenv("KAFKA_TOPIC")})
}

func (c *ConsumerTest) Handler(event interface{}, metadata AMQStream.ConsumerMetadata) error {
	fmt.Printf("Evento: %v\n", event)
	return nil
}

func (p *PublisherTest) To(event AMQStream.ISpecificRecord, key string) error {
	err := AMQStream.To(event, key)
	if err != nil {
		return err
	}
	return nil
}

func ConfigTopicForProducer(event AMQStream.ISpecificRecord, topics []string) {
	config, err := AMQStream.AddKafka()
	if err != nil {
		log.Fatal(err)
	}
	config.ToProducer(event, topics)
}

func LisenToEvents() {
	config, err := AMQStream.AddKafka()
	if err != nil {
		log.Fatal(err)
	}
	config.Build()
}
