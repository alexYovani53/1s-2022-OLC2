package expresion

type Accesso struct {
	valor int
}

func NewAccesso(valor int) Accesso {
	return Accesso{valor: valor}
}
