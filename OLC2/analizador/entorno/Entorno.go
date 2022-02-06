package entorno

type Entorno struct {
	Nombre      string
	EntAnterior *Entorno
	Tabla       map[string]interface{}
}

func NewEntorno(nombre string, entAnterior *Entorno) Entorno {

	en := Entorno{Nombre: nombre, EntAnterior: entAnterior}
	en.Tabla = make(map[string]interface{})
	return en
}
