package analizador

type ErrorSemantico struct {
	linea   int
	columna int
	msg     string
}

func NewErrorSemantico(linea int, columna int, msg string) ErrorSemantico {
	return ErrorSemantico{linea: linea, columna: columna, msg: msg}
}

var ListaErrores []ErrorSemantico
