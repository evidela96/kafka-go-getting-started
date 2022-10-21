// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     DetalleDeArticulo.avsc
 */
package ArticuloEvent

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type DetalleDeArticulo struct {
	Codigo string `json:"Codigo"`

	EAN13 *UnionNullString `json:"EAN13"`

	Propietario string `json:"Propietario"`

	Lote *UnionNullLote `json:"Lote"`

	OtrosDatos *UnionNullArrayMetadato `json:"OtrosDatos"`

	Descripcion string `json:"Descripcion"`

	ClaseDeExpedicion *UnionNullString `json:"ClaseDeExpedicion"`

	ClaseDeArticulo *UnionNullString `json:"ClaseDeArticulo"`

	PaisDeOrigen *UnionNullString `json:"PaisDeOrigen"`

	EsNumeroDeSerieDeEntradaUnico *UnionNullBool `json:"EsNumeroDeSerieDeEntradaUnico"`

	RequiereCapturaDatosEntrada *UnionNullBool `json:"RequiereCapturaDatosEntrada"`

	EsNumeroDeSerieSalidaUnico *UnionNullBool `json:"EsNumeroDeSerieSalidaUnico"`

	RequiereCapturaDatosSalida *UnionNullBool `json:"RequiereCapturaDatosSalida"`

	RequierecapturaTotalNumSeries *UnionNullBool `json:"RequierecapturaTotalNumSeries"`

	Caracteristicas *UnionNullArrayMetadato `json:"Caracteristicas"`

	Notas *UnionNullString `json:"Notas"`

	InstruccionesDePreparacion *UnionNullString `json:"InstruccionesDePreparacion"`

	VidaUtilEnDias *UnionNullLong `json:"VidaUtilEnDias"`

	CodigoDeVidaUtil *UnionNullString `json:"CodigoDeVidaUtil"`

	IndicadorDeVidaUtil *UnionNullString `json:"IndicadorDeVidaUtil"`

	ConsumoEnDias *UnionNullLong `json:"ConsumoEnDias"`

	VencimientoEnDias *UnionNullLong `json:"VencimientoEnDias"`

	VidaUtilEntradaEnDias *UnionNullLong `json:"VidaUtilEntradaEnDias"`

	AcondicionamientoSecundario *UnionNullString `json:"AcondicionamientoSecundario"`

	ZonaRepo *UnionNullString `json:"ZonaRepo"`

	Grupos *UnionNullArrayMetadato `json:"Grupos"`

	Volumen *UnionNullDouble `json:"Volumen"`

	PesoBruto *UnionNullDouble `json:"PesoBruto"`

	PesoTara *UnionNullDouble `json:"PesoTara"`

	PesoNeto *UnionNullDouble `json:"PesoNeto"`

	CamposLibres *UnionNullArrayMetadato `json:"CamposLibres"`
}

const DetalleDeArticuloAvroCRC64Fingerprint = "\a\xf7sy;\x95\xf7V"

func NewDetalleDeArticulo() DetalleDeArticulo {
	r := DetalleDeArticulo{}
	r.EAN13 = nil
	r.Lote = nil
	r.OtrosDatos = nil
	r.ClaseDeExpedicion = nil
	r.ClaseDeArticulo = nil
	r.PaisDeOrigen = nil
	r.EsNumeroDeSerieDeEntradaUnico = nil
	r.RequiereCapturaDatosEntrada = nil
	r.EsNumeroDeSerieSalidaUnico = nil
	r.RequiereCapturaDatosSalida = nil
	r.RequierecapturaTotalNumSeries = nil
	r.Caracteristicas = nil
	r.Notas = nil
	r.InstruccionesDePreparacion = nil
	r.VidaUtilEnDias = nil
	r.CodigoDeVidaUtil = nil
	r.IndicadorDeVidaUtil = nil
	r.ConsumoEnDias = nil
	r.VencimientoEnDias = nil
	r.VidaUtilEntradaEnDias = nil
	r.AcondicionamientoSecundario = nil
	r.ZonaRepo = nil
	r.Grupos = nil
	r.Volumen = nil
	r.PesoBruto = nil
	r.PesoTara = nil
	r.PesoNeto = nil
	r.CamposLibres = nil
	return r
}

func DeserializeDetalleDeArticulo(r io.Reader) (DetalleDeArticulo, error) {
	t := NewDetalleDeArticulo()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeDetalleDeArticuloFromSchema(r io.Reader, schema string) (DetalleDeArticulo, error) {
	t := NewDetalleDeArticulo()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeDetalleDeArticulo(r DetalleDeArticulo, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Codigo, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.EAN13, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Propietario, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLote(r.Lote, w)
	if err != nil {
		return err
	}
	err = writeUnionNullArrayMetadato(r.OtrosDatos, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Descripcion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ClaseDeExpedicion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ClaseDeArticulo, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.PaisDeOrigen, w)
	if err != nil {
		return err
	}
	err = writeUnionNullBool(r.EsNumeroDeSerieDeEntradaUnico, w)
	if err != nil {
		return err
	}
	err = writeUnionNullBool(r.RequiereCapturaDatosEntrada, w)
	if err != nil {
		return err
	}
	err = writeUnionNullBool(r.EsNumeroDeSerieSalidaUnico, w)
	if err != nil {
		return err
	}
	err = writeUnionNullBool(r.RequiereCapturaDatosSalida, w)
	if err != nil {
		return err
	}
	err = writeUnionNullBool(r.RequierecapturaTotalNumSeries, w)
	if err != nil {
		return err
	}
	err = writeUnionNullArrayMetadato(r.Caracteristicas, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Notas, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.InstruccionesDePreparacion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.VidaUtilEnDias, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CodigoDeVidaUtil, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.IndicadorDeVidaUtil, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.ConsumoEnDias, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.VencimientoEnDias, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.VidaUtilEntradaEnDias, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.AcondicionamientoSecundario, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ZonaRepo, w)
	if err != nil {
		return err
	}
	err = writeUnionNullArrayMetadato(r.Grupos, w)
	if err != nil {
		return err
	}
	err = writeUnionNullDouble(r.Volumen, w)
	if err != nil {
		return err
	}
	err = writeUnionNullDouble(r.PesoBruto, w)
	if err != nil {
		return err
	}
	err = writeUnionNullDouble(r.PesoTara, w)
	if err != nil {
		return err
	}
	err = writeUnionNullDouble(r.PesoNeto, w)
	if err != nil {
		return err
	}
	err = writeUnionNullArrayMetadato(r.CamposLibres, w)
	if err != nil {
		return err
	}
	return err
}

func (r DetalleDeArticulo) Serialize(w io.Writer) error {
	return writeDetalleDeArticulo(r, w)
}

func (r DetalleDeArticulo) Schema() string {
	return "{\"fields\":[{\"name\":\"Codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"EAN13\",\"type\":[\"null\",\"string\"]},{\"name\":\"Propietario\",\"type\":\"string\"},{\"default\":null,\"name\":\"Lote\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"Codigo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"LoteDeFabricante\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"LoteSecundario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FechaDeVencimiento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"OtrosDatos\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"Meta\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Contenido\",\"type\":[\"null\",\"string\"]}],\"name\":\"Metadato\",\"type\":\"record\"}]}],\"name\":\"Lote\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"OtrosDatos\",\"type\":[\"null\",{\"items\":\"Andreani.EventoWhArticulos.Events.Record.Metadato\",\"type\":\"array\"}]},{\"name\":\"Descripcion\",\"type\":\"string\"},{\"default\":null,\"name\":\"ClaseDeExpedicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ClaseDeArticulo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"PaisDeOrigen\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"EsNumeroDeSerieDeEntradaUnico\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"RequiereCapturaDatosEntrada\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"EsNumeroDeSerieSalidaUnico\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"RequiereCapturaDatosSalida\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"RequierecapturaTotalNumSeries\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"Caracteristicas\",\"type\":[\"null\",{\"items\":\"Andreani.EventoWhArticulos.Events.Record.Metadato\",\"type\":\"array\"}]},{\"default\":null,\"name\":\"Notas\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"InstruccionesDePreparacion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"VidaUtilEnDias\",\"type\":[\"null\",\"long\"]},{\"default\":null,\"name\":\"CodigoDeVidaUtil\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"IndicadorDeVidaUtil\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ConsumoEnDias\",\"type\":[\"null\",\"long\"]},{\"default\":null,\"name\":\"VencimientoEnDias\",\"type\":[\"null\",\"long\"]},{\"default\":null,\"name\":\"VidaUtilEntradaEnDias\",\"type\":[\"null\",\"long\"]},{\"default\":null,\"name\":\"AcondicionamientoSecundario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ZonaRepo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Grupos\",\"type\":[\"null\",{\"items\":\"Andreani.EventoWhArticulos.Events.Record.Metadato\",\"type\":\"array\"}]},{\"default\":null,\"name\":\"Volumen\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"PesoBruto\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"PesoTara\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"PesoNeto\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"CamposLibres\",\"type\":[\"null\",{\"items\":\"Andreani.EventoWhArticulos.Events.Record.Metadato\",\"type\":\"array\"}]}],\"name\":\"Andreani.EventoWhArticulos.Events.Record.DetalleDeArticulo\",\"type\":\"record\"}"
}

func (r DetalleDeArticulo) SchemaName() string {
	return "Andreani.EventoWhArticulos.Events.Record.DetalleDeArticulo"
}

func (_ DetalleDeArticulo) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ DetalleDeArticulo) SetInt(v int32)       { panic("Unsupported operation") }
func (_ DetalleDeArticulo) SetLong(v int64)      { panic("Unsupported operation") }
func (_ DetalleDeArticulo) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ DetalleDeArticulo) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ DetalleDeArticulo) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ DetalleDeArticulo) SetString(v string)   { panic("Unsupported operation") }
func (_ DetalleDeArticulo) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *DetalleDeArticulo) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Codigo}

		return w

	case 1:
		r.EAN13 = NewUnionNullString()

		return r.EAN13
	case 2:
		w := types.String{Target: &r.Propietario}

		return w

	case 3:
		r.Lote = NewUnionNullLote()

		return r.Lote
	case 4:
		r.OtrosDatos = NewUnionNullArrayMetadato()

		return r.OtrosDatos
	case 5:
		w := types.String{Target: &r.Descripcion}

		return w

	case 6:
		r.ClaseDeExpedicion = NewUnionNullString()

		return r.ClaseDeExpedicion
	case 7:
		r.ClaseDeArticulo = NewUnionNullString()

		return r.ClaseDeArticulo
	case 8:
		r.PaisDeOrigen = NewUnionNullString()

		return r.PaisDeOrigen
	case 9:
		r.EsNumeroDeSerieDeEntradaUnico = NewUnionNullBool()

		return r.EsNumeroDeSerieDeEntradaUnico
	case 10:
		r.RequiereCapturaDatosEntrada = NewUnionNullBool()

		return r.RequiereCapturaDatosEntrada
	case 11:
		r.EsNumeroDeSerieSalidaUnico = NewUnionNullBool()

		return r.EsNumeroDeSerieSalidaUnico
	case 12:
		r.RequiereCapturaDatosSalida = NewUnionNullBool()

		return r.RequiereCapturaDatosSalida
	case 13:
		r.RequierecapturaTotalNumSeries = NewUnionNullBool()

		return r.RequierecapturaTotalNumSeries
	case 14:
		r.Caracteristicas = NewUnionNullArrayMetadato()

		return r.Caracteristicas
	case 15:
		r.Notas = NewUnionNullString()

		return r.Notas
	case 16:
		r.InstruccionesDePreparacion = NewUnionNullString()

		return r.InstruccionesDePreparacion
	case 17:
		r.VidaUtilEnDias = NewUnionNullLong()

		return r.VidaUtilEnDias
	case 18:
		r.CodigoDeVidaUtil = NewUnionNullString()

		return r.CodigoDeVidaUtil
	case 19:
		r.IndicadorDeVidaUtil = NewUnionNullString()

		return r.IndicadorDeVidaUtil
	case 20:
		r.ConsumoEnDias = NewUnionNullLong()

		return r.ConsumoEnDias
	case 21:
		r.VencimientoEnDias = NewUnionNullLong()

		return r.VencimientoEnDias
	case 22:
		r.VidaUtilEntradaEnDias = NewUnionNullLong()

		return r.VidaUtilEntradaEnDias
	case 23:
		r.AcondicionamientoSecundario = NewUnionNullString()

		return r.AcondicionamientoSecundario
	case 24:
		r.ZonaRepo = NewUnionNullString()

		return r.ZonaRepo
	case 25:
		r.Grupos = NewUnionNullArrayMetadato()

		return r.Grupos
	case 26:
		r.Volumen = NewUnionNullDouble()

		return r.Volumen
	case 27:
		r.PesoBruto = NewUnionNullDouble()

		return r.PesoBruto
	case 28:
		r.PesoTara = NewUnionNullDouble()

		return r.PesoTara
	case 29:
		r.PesoNeto = NewUnionNullDouble()

		return r.PesoNeto
	case 30:
		r.CamposLibres = NewUnionNullArrayMetadato()

		return r.CamposLibres
	}
	panic("Unknown field index")
}

func (r *DetalleDeArticulo) SetDefault(i int) {
	switch i {
	case 1:
		r.EAN13 = nil
		return
	case 3:
		r.Lote = nil
		return
	case 4:
		r.OtrosDatos = nil
		return
	case 6:
		r.ClaseDeExpedicion = nil
		return
	case 7:
		r.ClaseDeArticulo = nil
		return
	case 8:
		r.PaisDeOrigen = nil
		return
	case 9:
		r.EsNumeroDeSerieDeEntradaUnico = nil
		return
	case 10:
		r.RequiereCapturaDatosEntrada = nil
		return
	case 11:
		r.EsNumeroDeSerieSalidaUnico = nil
		return
	case 12:
		r.RequiereCapturaDatosSalida = nil
		return
	case 13:
		r.RequierecapturaTotalNumSeries = nil
		return
	case 14:
		r.Caracteristicas = nil
		return
	case 15:
		r.Notas = nil
		return
	case 16:
		r.InstruccionesDePreparacion = nil
		return
	case 17:
		r.VidaUtilEnDias = nil
		return
	case 18:
		r.CodigoDeVidaUtil = nil
		return
	case 19:
		r.IndicadorDeVidaUtil = nil
		return
	case 20:
		r.ConsumoEnDias = nil
		return
	case 21:
		r.VencimientoEnDias = nil
		return
	case 22:
		r.VidaUtilEntradaEnDias = nil
		return
	case 23:
		r.AcondicionamientoSecundario = nil
		return
	case 24:
		r.ZonaRepo = nil
		return
	case 25:
		r.Grupos = nil
		return
	case 26:
		r.Volumen = nil
		return
	case 27:
		r.PesoBruto = nil
		return
	case 28:
		r.PesoTara = nil
		return
	case 29:
		r.PesoNeto = nil
		return
	case 30:
		r.CamposLibres = nil
		return
	}
	panic("Unknown field index")
}

func (r *DetalleDeArticulo) NullField(i int) {
	switch i {
	case 1:
		r.EAN13 = nil
		return
	case 3:
		r.Lote = nil
		return
	case 4:
		r.OtrosDatos = nil
		return
	case 6:
		r.ClaseDeExpedicion = nil
		return
	case 7:
		r.ClaseDeArticulo = nil
		return
	case 8:
		r.PaisDeOrigen = nil
		return
	case 9:
		r.EsNumeroDeSerieDeEntradaUnico = nil
		return
	case 10:
		r.RequiereCapturaDatosEntrada = nil
		return
	case 11:
		r.EsNumeroDeSerieSalidaUnico = nil
		return
	case 12:
		r.RequiereCapturaDatosSalida = nil
		return
	case 13:
		r.RequierecapturaTotalNumSeries = nil
		return
	case 14:
		r.Caracteristicas = nil
		return
	case 15:
		r.Notas = nil
		return
	case 16:
		r.InstruccionesDePreparacion = nil
		return
	case 17:
		r.VidaUtilEnDias = nil
		return
	case 18:
		r.CodigoDeVidaUtil = nil
		return
	case 19:
		r.IndicadorDeVidaUtil = nil
		return
	case 20:
		r.ConsumoEnDias = nil
		return
	case 21:
		r.VencimientoEnDias = nil
		return
	case 22:
		r.VidaUtilEntradaEnDias = nil
		return
	case 23:
		r.AcondicionamientoSecundario = nil
		return
	case 24:
		r.ZonaRepo = nil
		return
	case 25:
		r.Grupos = nil
		return
	case 26:
		r.Volumen = nil
		return
	case 27:
		r.PesoBruto = nil
		return
	case 28:
		r.PesoTara = nil
		return
	case 29:
		r.PesoNeto = nil
		return
	case 30:
		r.CamposLibres = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ DetalleDeArticulo) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ DetalleDeArticulo) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ DetalleDeArticulo) HintSize(int)                     { panic("Unsupported operation") }
func (_ DetalleDeArticulo) Finalize()                        {}

func (_ DetalleDeArticulo) AvroCRC64Fingerprint() []byte {
	return []byte(DetalleDeArticuloAvroCRC64Fingerprint)
}

func (r DetalleDeArticulo) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Codigo"], err = json.Marshal(r.Codigo)
	if err != nil {
		return nil, err
	}
	output["EAN13"], err = json.Marshal(r.EAN13)
	if err != nil {
		return nil, err
	}
	output["Propietario"], err = json.Marshal(r.Propietario)
	if err != nil {
		return nil, err
	}
	output["Lote"], err = json.Marshal(r.Lote)
	if err != nil {
		return nil, err
	}
	output["OtrosDatos"], err = json.Marshal(r.OtrosDatos)
	if err != nil {
		return nil, err
	}
	output["Descripcion"], err = json.Marshal(r.Descripcion)
	if err != nil {
		return nil, err
	}
	output["ClaseDeExpedicion"], err = json.Marshal(r.ClaseDeExpedicion)
	if err != nil {
		return nil, err
	}
	output["ClaseDeArticulo"], err = json.Marshal(r.ClaseDeArticulo)
	if err != nil {
		return nil, err
	}
	output["PaisDeOrigen"], err = json.Marshal(r.PaisDeOrigen)
	if err != nil {
		return nil, err
	}
	output["EsNumeroDeSerieDeEntradaUnico"], err = json.Marshal(r.EsNumeroDeSerieDeEntradaUnico)
	if err != nil {
		return nil, err
	}
	output["RequiereCapturaDatosEntrada"], err = json.Marshal(r.RequiereCapturaDatosEntrada)
	if err != nil {
		return nil, err
	}
	output["EsNumeroDeSerieSalidaUnico"], err = json.Marshal(r.EsNumeroDeSerieSalidaUnico)
	if err != nil {
		return nil, err
	}
	output["RequiereCapturaDatosSalida"], err = json.Marshal(r.RequiereCapturaDatosSalida)
	if err != nil {
		return nil, err
	}
	output["RequierecapturaTotalNumSeries"], err = json.Marshal(r.RequierecapturaTotalNumSeries)
	if err != nil {
		return nil, err
	}
	output["Caracteristicas"], err = json.Marshal(r.Caracteristicas)
	if err != nil {
		return nil, err
	}
	output["Notas"], err = json.Marshal(r.Notas)
	if err != nil {
		return nil, err
	}
	output["InstruccionesDePreparacion"], err = json.Marshal(r.InstruccionesDePreparacion)
	if err != nil {
		return nil, err
	}
	output["VidaUtilEnDias"], err = json.Marshal(r.VidaUtilEnDias)
	if err != nil {
		return nil, err
	}
	output["CodigoDeVidaUtil"], err = json.Marshal(r.CodigoDeVidaUtil)
	if err != nil {
		return nil, err
	}
	output["IndicadorDeVidaUtil"], err = json.Marshal(r.IndicadorDeVidaUtil)
	if err != nil {
		return nil, err
	}
	output["ConsumoEnDias"], err = json.Marshal(r.ConsumoEnDias)
	if err != nil {
		return nil, err
	}
	output["VencimientoEnDias"], err = json.Marshal(r.VencimientoEnDias)
	if err != nil {
		return nil, err
	}
	output["VidaUtilEntradaEnDias"], err = json.Marshal(r.VidaUtilEntradaEnDias)
	if err != nil {
		return nil, err
	}
	output["AcondicionamientoSecundario"], err = json.Marshal(r.AcondicionamientoSecundario)
	if err != nil {
		return nil, err
	}
	output["ZonaRepo"], err = json.Marshal(r.ZonaRepo)
	if err != nil {
		return nil, err
	}
	output["Grupos"], err = json.Marshal(r.Grupos)
	if err != nil {
		return nil, err
	}
	output["Volumen"], err = json.Marshal(r.Volumen)
	if err != nil {
		return nil, err
	}
	output["PesoBruto"], err = json.Marshal(r.PesoBruto)
	if err != nil {
		return nil, err
	}
	output["PesoTara"], err = json.Marshal(r.PesoTara)
	if err != nil {
		return nil, err
	}
	output["PesoNeto"], err = json.Marshal(r.PesoNeto)
	if err != nil {
		return nil, err
	}
	output["CamposLibres"], err = json.Marshal(r.CamposLibres)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *DetalleDeArticulo) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Codigo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Codigo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Codigo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["EAN13"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EAN13); err != nil {
			return err
		}
	} else {
		r.EAN13 = NewUnionNullString()

		r.EAN13 = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Propietario"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Propietario); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Propietario")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Lote"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Lote); err != nil {
			return err
		}
	} else {
		r.Lote = NewUnionNullLote()

		r.Lote = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["OtrosDatos"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.OtrosDatos); err != nil {
			return err
		}
	} else {
		r.OtrosDatos = NewUnionNullArrayMetadato()

		r.OtrosDatos = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Descripcion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Descripcion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Descripcion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ClaseDeExpedicion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ClaseDeExpedicion); err != nil {
			return err
		}
	} else {
		r.ClaseDeExpedicion = NewUnionNullString()

		r.ClaseDeExpedicion = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["ClaseDeArticulo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ClaseDeArticulo); err != nil {
			return err
		}
	} else {
		r.ClaseDeArticulo = NewUnionNullString()

		r.ClaseDeArticulo = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["PaisDeOrigen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PaisDeOrigen); err != nil {
			return err
		}
	} else {
		r.PaisDeOrigen = NewUnionNullString()

		r.PaisDeOrigen = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["EsNumeroDeSerieDeEntradaUnico"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EsNumeroDeSerieDeEntradaUnico); err != nil {
			return err
		}
	} else {
		r.EsNumeroDeSerieDeEntradaUnico = NewUnionNullBool()

		r.EsNumeroDeSerieDeEntradaUnico = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["RequiereCapturaDatosEntrada"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.RequiereCapturaDatosEntrada); err != nil {
			return err
		}
	} else {
		r.RequiereCapturaDatosEntrada = NewUnionNullBool()

		r.RequiereCapturaDatosEntrada = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["EsNumeroDeSerieSalidaUnico"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EsNumeroDeSerieSalidaUnico); err != nil {
			return err
		}
	} else {
		r.EsNumeroDeSerieSalidaUnico = NewUnionNullBool()

		r.EsNumeroDeSerieSalidaUnico = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["RequiereCapturaDatosSalida"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.RequiereCapturaDatosSalida); err != nil {
			return err
		}
	} else {
		r.RequiereCapturaDatosSalida = NewUnionNullBool()

		r.RequiereCapturaDatosSalida = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["RequierecapturaTotalNumSeries"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.RequierecapturaTotalNumSeries); err != nil {
			return err
		}
	} else {
		r.RequierecapturaTotalNumSeries = NewUnionNullBool()

		r.RequierecapturaTotalNumSeries = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Caracteristicas"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Caracteristicas); err != nil {
			return err
		}
	} else {
		r.Caracteristicas = NewUnionNullArrayMetadato()

		r.Caracteristicas = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Notas"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Notas); err != nil {
			return err
		}
	} else {
		r.Notas = NewUnionNullString()

		r.Notas = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["InstruccionesDePreparacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.InstruccionesDePreparacion); err != nil {
			return err
		}
	} else {
		r.InstruccionesDePreparacion = NewUnionNullString()

		r.InstruccionesDePreparacion = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["VidaUtilEnDias"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.VidaUtilEnDias); err != nil {
			return err
		}
	} else {
		r.VidaUtilEnDias = NewUnionNullLong()

		r.VidaUtilEnDias = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["CodigoDeVidaUtil"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoDeVidaUtil); err != nil {
			return err
		}
	} else {
		r.CodigoDeVidaUtil = NewUnionNullString()

		r.CodigoDeVidaUtil = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["IndicadorDeVidaUtil"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IndicadorDeVidaUtil); err != nil {
			return err
		}
	} else {
		r.IndicadorDeVidaUtil = NewUnionNullString()

		r.IndicadorDeVidaUtil = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["ConsumoEnDias"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ConsumoEnDias); err != nil {
			return err
		}
	} else {
		r.ConsumoEnDias = NewUnionNullLong()

		r.ConsumoEnDias = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["VencimientoEnDias"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.VencimientoEnDias); err != nil {
			return err
		}
	} else {
		r.VencimientoEnDias = NewUnionNullLong()

		r.VencimientoEnDias = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["VidaUtilEntradaEnDias"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.VidaUtilEntradaEnDias); err != nil {
			return err
		}
	} else {
		r.VidaUtilEntradaEnDias = NewUnionNullLong()

		r.VidaUtilEntradaEnDias = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["AcondicionamientoSecundario"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.AcondicionamientoSecundario); err != nil {
			return err
		}
	} else {
		r.AcondicionamientoSecundario = NewUnionNullString()

		r.AcondicionamientoSecundario = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["ZonaRepo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ZonaRepo); err != nil {
			return err
		}
	} else {
		r.ZonaRepo = NewUnionNullString()

		r.ZonaRepo = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Grupos"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Grupos); err != nil {
			return err
		}
	} else {
		r.Grupos = NewUnionNullArrayMetadato()

		r.Grupos = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Volumen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Volumen); err != nil {
			return err
		}
	} else {
		r.Volumen = NewUnionNullDouble()

		r.Volumen = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["PesoBruto"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PesoBruto); err != nil {
			return err
		}
	} else {
		r.PesoBruto = NewUnionNullDouble()

		r.PesoBruto = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["PesoTara"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PesoTara); err != nil {
			return err
		}
	} else {
		r.PesoTara = NewUnionNullDouble()

		r.PesoTara = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["PesoNeto"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PesoNeto); err != nil {
			return err
		}
	} else {
		r.PesoNeto = NewUnionNullDouble()

		r.PesoNeto = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["CamposLibres"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CamposLibres); err != nil {
			return err
		}
	} else {
		r.CamposLibres = NewUnionNullArrayMetadato()

		r.CamposLibres = nil
	}
	return nil
}