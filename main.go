package main

import (
	"log"

	"github.com/architecture-it/go-platform/AMQStream"
	"github.com/architecture-it/integracion-schemas-event-go/Onboarding/Events"
)

type publisherTest struct {
}

func (p *publisherTest) To(event AMQStream.ISpecificRecord, key string) error {
	config, err := AMQStream.AddKafka()
	if err != nil {
		return err
	}
	topics := []string{"PedidoAsignadoTest"}
	config.ToProducer(event, topics)
	return nil
}

func main() {
	event := Events.Pedido{
		Id:                      "c2781822-5f08-4302-bb9e-8154cc12f43a",
		NumeroDePedido:          20,
		EstadoDelPedido:         "Pedido123",
		CicloDelPedido:          "c2781822-5f08-4302-bb9e-8154cc12f43a",
		CodigoDeContratoInterno: 123123123,
		CuentaCorriente:         5231235123,
		Cuando:                  "Nose",
	}
	config, err := AMQStream.AddKafka()
	if err != nil {
		log.Fatal(err)
	}
	topics := []string{"PedidoAsignadoTest"}
	config.ToProducer(&event, topics)

	err = AMQStream.To(&event, "key123")
	if err != nil {
		log.Fatal(err)
	}
	//publish.publishTest(event)

	// suscriptor := SuscriberTest{}
	// config.ToConsumer(suscriptor, &event, "PedidoAsignadoTest")
	// var wg sync.WaitGroup

	// wg.Add(1)
	// config.Build()
	// wg.Wait()

}

/* type publisherTest struct {
}

type SuscriberTest struct {
}

func (p publisherTest) publish(event AMQStream.ISpecificRecord, c *AMQStream.Config) {

	c.To(event, "key1234")
}

func (p publisherTest) publishTest(event Events.Pedido) {
	config, _ := AMQStream.AddKafka()
	topics := []string{"PedidoAsignado", "PedidoAsignadoTest"}
	config.ToProducer(&event, topics)

	config.To(&event, "key123")
	fmt.Printf("Pedido publicado: %v\n", &event)
}

func (s SuscriberTest) Handler(event AMQStream.ISpecificRecord, metadata AMQStream.ConsumerMetadata) {

	fmt.Printf("EventoTest PedidoAsignado: %v\n", event)
}

func (s SuscriberTest) Setup(publisher AMQStream.IPublisher) {

} */
