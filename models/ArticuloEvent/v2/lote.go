// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     DetalleDeArticulo.avsc
 */
package avro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Lote struct {
	Codigo string `json:"Codigo"`

	LoteDeFabricante string `json:"LoteDeFabricante"`

	LoteSecundario string `json:"LoteSecundario"`

	FechaDeVencimiento string `json:"FechaDeVencimiento"`

	OtrosDatos []Metadato `json:"OtrosDatos"`
}

const LoteAvroCRC64Fingerprint = "\xd4]Gkmđ\xd0"

func NewLote() Lote {
	r := Lote{}
	r.OtrosDatos = make([]Metadato, 0)

	return r
}

func DeserializeLote(r io.Reader) (Lote, error) {
	t := NewLote()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeLoteFromSchema(r io.Reader, schema string) (Lote, error) {
	t := NewLote()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeLote(r Lote, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Codigo, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.LoteDeFabricante, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.LoteSecundario, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.FechaDeVencimiento, w)
	if err != nil {
		return err
	}
	err = writeArrayMetadato(r.OtrosDatos, w)
	if err != nil {
		return err
	}
	return err
}

func (r Lote) Serialize(w io.Writer) error {
	return writeLote(r, w)
}

func (r Lote) Schema() string {
	return "{\"fields\":[{\"name\":\"Codigo\",\"type\":\"string\"},{\"name\":\"LoteDeFabricante\",\"type\":\"string\"},{\"name\":\"LoteSecundario\",\"type\":\"string\"},{\"name\":\"FechaDeVencimiento\",\"type\":\"string\"},{\"name\":\"OtrosDatos\",\"type\":{\"items\":{\"fields\":[{\"name\":\"Meta\",\"type\":\"string\"},{\"name\":\"Contenido\",\"type\":\"string\"}],\"name\":\"Metadato\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Andreani.EventoWhArticulos.Events.Record.Lote\",\"type\":\"record\"}"
}

func (r Lote) SchemaName() string {
	return "Andreani.EventoWhArticulos.Events.Record.Lote"
}

func (_ Lote) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Lote) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Lote) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Lote) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Lote) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Lote) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Lote) SetString(v string)   { panic("Unsupported operation") }
func (_ Lote) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Lote) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Codigo}

		return w

	case 1:
		w := types.String{Target: &r.LoteDeFabricante}

		return w

	case 2:
		w := types.String{Target: &r.LoteSecundario}

		return w

	case 3:
		w := types.String{Target: &r.FechaDeVencimiento}

		return w

	case 4:
		r.OtrosDatos = make([]Metadato, 0)

		w := ArrayMetadatoWrapper{Target: &r.OtrosDatos}

		return w

	}
	panic("Unknown field index")
}

func (r *Lote) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Lote) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Lote) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Lote) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Lote) HintSize(int)                     { panic("Unsupported operation") }
func (_ Lote) Finalize()                        {}

func (_ Lote) AvroCRC64Fingerprint() []byte {
	return []byte(LoteAvroCRC64Fingerprint)
}

func (r Lote) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Codigo"], err = json.Marshal(r.Codigo)
	if err != nil {
		return nil, err
	}
	output["LoteDeFabricante"], err = json.Marshal(r.LoteDeFabricante)
	if err != nil {
		return nil, err
	}
	output["LoteSecundario"], err = json.Marshal(r.LoteSecundario)
	if err != nil {
		return nil, err
	}
	output["FechaDeVencimiento"], err = json.Marshal(r.FechaDeVencimiento)
	if err != nil {
		return nil, err
	}
	output["OtrosDatos"], err = json.Marshal(r.OtrosDatos)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Lote) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["LoteDeFabricante"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LoteDeFabricante); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for LoteDeFabricante")
	}
	val = func() json.RawMessage {
		if v, ok := fields["LoteSecundario"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LoteSecundario); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for LoteSecundario")
	}
	val = func() json.RawMessage {
		if v, ok := fields["FechaDeVencimiento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaDeVencimiento); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for FechaDeVencimiento")
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
		return fmt.Errorf("no value specified for OtrosDatos")
	}
	return nil
}