package entorno

type TipoDato int
type TipoModAccess int

const (
	INTEGER TipoDato = iota
	FLOAT
	STRING
	BOOLEAN
	NULL
	VOID
	ARREGLO
	OBJETO
)

const (
	PUBLIC TipoModAccess = iota
	PRIVATE
)

type Result3D struct {
	Tipo        TipoDato
	Codigo      string
	Temporal    string
	EtiquetaF   string
	EtiquetaV   string
	Valor       interface{}
	ValorEnHeap bool
}
