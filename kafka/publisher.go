package kafka

import (
	"log"
	"os"

	"github.com/architecture-it/go-platform/AMQStream"
	TestEvents "github.com/architecture-it/integracion-schemas-event-go/Test"
)

type PublisherTest struct{}

var event TestEvents.KafkaDemoRetro
var topics []string = []string{os.Getenv("KAFKA_TOPIC")}

func init() {
	config, err := AMQStream.AddKafka()
	if err != nil {
		log.Fatal(err)
	}
	config.ToProducer(&event, topics)
}

func (p *PublisherTest) To(event AMQStream.ISpecificRecord, key string) error {
	err := AMQStream.To(event, key)
	if err != nil {
		return err
	}
	return nil
}
