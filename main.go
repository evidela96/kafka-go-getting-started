package main

import (
	"log"
	"os"

	kafkaConfig "github.com/evidela96/kafka-go-getting-started/KafkaConfig"
	v2 "github.com/evidela96/kafka-go-getting-started/models/ArticuloEvent/v2"
)

type ConsumerTest struct{}

func main() {
	/* kafkaDemoRetro := TestEvents.KafkaDemoRetro{
		Id:               12,
		Title:            "evidela",
		Attendance:       123,
		DidYouUnderstand: true,
		Me:               TestEvents.NewPerson(),
		Time:             123,
	} */
	/* articuloSol := ArticuloEvent.MantenimientoDeArticuloSolicitado{
		Contrato: "350001680",
		Almacen:  "wmwhse4",
		Planta:   "BENAVIDEZ",
		DetalleDeArticulo: ArticuloEvent.DetalleDeArticulo{
			Codigo: "7200204",
			EAN13: &ArticuloEvent.UnionNullString{
				Null:      &types.NullVal{},
				String:    "ean13",
				UnionType: ArticuloEvent.UnionNullStringTypeEnumString,
			},
			Propietario: "NOVO CUARENTENA",
			Descripcion: "jajajajaj ni idea",
			CamposLibres: &ArticuloEvent.UnionNullArrayMetadato{
				Null: &types.NullVal{},
				ArrayMetadato: []ArticuloEvent.Metadato{
					{
						Meta: &ArticuloEvent.UnionNullString{
							Null:      &types.NullVal{},
							String:    "meta1",
							UnionType: ArticuloEvent.UnionNullStringTypeEnumString,
						},
						Contenido: &ArticuloEvent.UnionNullString{
							Null:      &types.NullVal{},
							String:    "contenido1",
							UnionType: ArticuloEvent.UnionNullStringTypeEnumString,
						},
					},
					{
						Meta: &ArticuloEvent.UnionNullString{
							Null:      &types.NullVal{},
							String:    "meta2",
							UnionType: ArticuloEvent.UnionNullStringTypeEnumString,
						},
						Contenido: &ArticuloEvent.UnionNullString{
							Null:      &types.NullVal{},
							String:    "contenido2",
							UnionType: ArticuloEvent.UnionNullStringTypeEnumString,
						},
					},
				},
				UnionType: ArticuloEvent.UnionNullArrayMetadatoTypeEnumArrayMetadato,
			},
		},
	} */
	articuloV2 := v2.MantenimientoDeArticuloSolicitado{
		Contrato: "350001680",
		Almacen:  "wmwhse4",
		Planta:   "BENAVIDEZ",
		DetalleDeArticulo: v2.DetalleDeArticulo{
			Codigo:                        "7200204",
			Propietario:                   "NOVO CUARENTENA",
			Descripcion:                   "jajajajaj ni idea",
			EAN13:                         "1432143241234",
			EsNumeroDeSerieDeEntradaUnico: false,
			RequiereCapturaDatosEntrada:   true,
			EsNumeroDeSerieSalidaUnico:    true,
			VidaUtilEnDias:                1232141,
			ConsumoEnDias:                 7589403275,
			OtrosDatos: []v2.Metadato{
				{
					Meta:      "meta1",
					Contenido: "Contenido1",
				},
				{
					Meta:      "meta1",
					Contenido: "Contenido1",
				},
				{
					Meta:      "meta1",
					Contenido: "Contenido1",
				},
				{
					Meta:      "meta1",
					Contenido: "Contenido1",
				},
				{
					Meta:      "meta1",
					Contenido: "Contenido1",
				},
			},
		},
	}
	publisher := kafkaConfig.PublisherTest{}
	kafkaConfig.ConfigTopicForProducer(&articuloV2, []string{os.Getenv("KAFKA_TOPIC")})
	err := publisher.To(&articuloV2, "key123456")
	if err != nil {
		log.Fatal(err)
	}

	kafkaConfig.LisenToEvents()
}
