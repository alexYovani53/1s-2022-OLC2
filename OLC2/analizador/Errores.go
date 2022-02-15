package analizador

type ErrorSemantico struct {
	linea   int
	columna int
	msg     string
}

var Error []ErrorSemantico
