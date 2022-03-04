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
)

const (
	PUBLIC TipoModAccess = iota
	PRIVATE
)

type ValorType struct {
	Tipo  TipoDato
	Valor interface{}
}
