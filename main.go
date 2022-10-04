package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/architecture-it/go-platform/AMQStream"
	"github.com/architecture-it/integracion-schemas-event-go/Onboarding/Events"
)

type PublisherTest struct{}
type SuscriberTest struct{}

func main() {
	event := Events.Pedido{
		Id:                      "c2781822-5f08-4302-bb9e-8154cc12f43a",
		NumeroDePedido:          20,
		EstadoDelPedido:         "Pedido123",
		CicloDelPedido:          "c2781822-5f08-4302-bb9e-8154cc12f43a",
		CodigoDeContratoInterno: 123123123,
		CuentaCorriente:         5231235123,
		Cuando:                  "3/10/2022Z00:00:0000",
	}

	publish := PublisherTest{}
	publish.To(&event, "key123")

	config, err := AMQStream.AddKafka()
	if err != nil {
		log.Fatal(err)
	}
	suscriptor := SuscriberTest{}
	consumedEvent := Events.Pedido{}
	var wg sync.WaitGroup

	config.ToConsumer(&suscriptor, &consumedEvent, "PedidoAsignadoTest")

	wg.Add(1)
	go config.Build()
	wg.Wait()

}

func (p *PublisherTest) To(event AMQStream.ISpecificRecord, key string) error {
	config, err := AMQStream.AddKafka()
	if err != nil {
		return err
	}
	topics := []string{"PedidoAsignadoTest"}
	config.ToProducer(event, topics)
	err = AMQStream.To(event, "key123")
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (s *SuscriberTest) Handler(event AMQStream.ISpecificRecord, metadata AMQStream.ConsumerMetadata) {

	fmt.Printf("Evento: %v\n", event)
}
