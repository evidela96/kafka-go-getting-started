package main

import (
	"fmt"
	"log"

	"github.com/architecture-it/go-platform/AMQStream"
	TestEvents "github.com/architecture-it/integracion-schemas-event-go/Test"
)

type PublisherTest struct{}
type SuscriberTest struct{}

func main() {
	kafkaDemoRetro := TestEvents.KafkaDemoRetro{
		Id:               12,
		Title:            "evidela",
		Attendance:       123,
		DidYouUnderstand: true,
		Me:               TestEvents.NewPerson(),
		Time:             123,
	}
	//articulo2 := EventoWhArticulosEvents.NewEventoWhArticuloAsnConfirmacion()
	/* e := EventoWhArticulosEvents.EventoWhArticuloExpedicion{
		Identificacion: EventoWhArticulosEvents.Identificacion{},
		Cabecera:       EventoWhArticulosEvents.Cabecera{},
		Detalle:        []EventoWhArticulosEvents.Detalle{},
	} */
	/* identificacion := EventoWhArticulosEvents.Identificacion{
		Id:                  "123",
		Evento:              "evidela-Evento",
		Proceso:             "Alta Articulo",
		FechaHoraGeneracion: 123,
		SistemaOrigen:       "PC",
		Almacen:             "wmwhse1",
		Propietario:         "evidela",
		Instancia:           "instancia",
	}  */
	/* articulo := EventoWhArticulosEvents.EventoWhArticuloExpedicion{
		Identificacion: EventoWhArticulosEvents.Identificacion{
			Id:                  "123",
			Evento:              "evidela-Evento",
			Proceso:             "Alta Articulo",
			FechaHoraGeneracion: 123,
			SistemaOrigen:       "PC",
			Almacen:             "wmwhse1",
			Propietario:         "evidela",
			Instancia:           "instancia",
		},
		Cabecera: EventoWhArticulosEvents.Cabecera{
			SKU:                 "sku123",
			Descripcion:         "eeeee messi",
			Propietario:         "evidela",
			TipoOrigen:          "original",
			CodigoOrigenWH:      "or1234",
			CodigoOrigenExterno: "ore1234",
		},
		Detalle: []EventoWhArticulosEvents.Detalle{
			{
				PaqueteLote:          "PaqueteLote",
				LoteCajitaFabricante: "LoteCajitaFabricante",
				LoteSecundario:       "LoteSecundario",
				FechaFabricacion: &EventoWhArticulosEvents.UnionNullLong{
					Null:      &types.NullVal{},
					Long:      12,
					UnionType: 12,
				},
				FechaVencimiento: &EventoWhArticulosEvents.UnionNullLong{
					Null:      &types.NullVal{},
					Long:      15,
					UnionType: 15,
				},
				ProductoTrazable: "ProductoTrazable",
				AlmacenConsumo:   "AlmacenConsumo",
				EstadoLote:       "EstadoLote",
				VidaUtilLote:     "VidaUtilLote",
				EntregaAntesDe: &EventoWhArticulosEvents.UnionNullLong{
					Null:      &types.NullVal{},
					Long:      15,
					UnionType: 15,
				},
				ConsumoAntesDe: &EventoWhArticulosEvents.UnionNullLong{
					Null:      &types.NullVal{},
					Long:      15,
					UnionType: 15,
				},
				StockDisponible: 100,
				StockEnTransito: 50,
			},
		},
	} */

	/* event := Events.Pedido{
		Id:                      "c2781822-5f08-4302-bb9e-8154cc12f43a",
		NumeroDePedido:          20,
		EstadoDelPedido:         "Pedido123",
		CicloDelPedido:          "c2781822-5f08-4302-bb9e-8154cc12f43a",
		CodigoDeContratoInterno: 123123123,
		CuentaCorriente:         5231235123,
		Cuando:                  "3/10/2022Z00:00:0000",
	} */
	//articuloExpedido := EventoWhArticulosEvents.NewEventoWhArticuloExpedicion()
	publish := &PublisherTest{}
	err := publish.To(&kafkaDemoRetro, "articulo2key")
	if err != nil {
		log.Fatal(err)
	}
	/* config, err := AMQStream.AddKafka()

	suscriptor := SuscriberTest{}
	//consumedEvent := Events.Pedido{}
	articulo := EventoWhArticulosEvents.EventoWhActiculoAsnConfirmacion{}
	var wg sync.WaitGroup

	config.ToConsumer(&suscriptor, &articulo, "EventoWh-ArticulosAsnConfirmado")

	wg.Add(1)
	go config.Build()
	wg.Wait()
	*/
}

func (p *PublisherTest) To(event AMQStream.ISpecificRecord, key string) error {
	config, err := AMQStream.AddKafka()
	if err != nil {
		return err
	}
	topics := []string{"KafkaDemoTest"}
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
