package entorno

type Simbolo struct {
	Linea         int
	Columna       int
	Identificador string
	Valor         interface{}
	Tipo          TipoDato
	Constante     bool
	EsFuncion     bool
}

/**
 * @comentario  Constructor que se utiliza solo para guardar un ID en un simbolo
 * @param       identificador  -> Identificador del Simbolo.
 */
/*

	int d =  x + y;

*/

func NewSimboloIdentificador(linea int, columna int, identificador string) *Simbolo {
	return &Simbolo{
		Linea:         linea,
		Columna:       columna,
		Identificador: identificador,
		Constante:     false,
		EsFuncion:     false,
		Valor:         nil,
	}
}

func NewSimboloIdentificadorValor(linea int, columna int, identificador string, valor interface{}, dato TipoDato) Simbolo {
	e := Simbolo{
		Linea:         linea,
		Columna:       columna,
		Identificador: identificador,
		Constante:     false,
		EsFuncion:     false,
		Valor:         valor,
		Tipo:          dato,
	}
	return e
}

