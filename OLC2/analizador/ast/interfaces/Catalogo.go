package interfaces

type TipoDato int

const (
	INTEGER TipoDato = iota
	FLOAT
	STRING
	BOOLEAN
	NULL
)

type RetornoType struct {
	Tipo  TipoDato
	Valor interface{}
}
