package mdl

import "runtime"

// Sensor de decisiones
type Sensor struct {
	Codigo     int
	Tareas     int
	Memoria    string
	Cantidad   int
	Limite     int
	Espacio    int
	MaxNucleos int
	Min        int
}

func init() {
	var ld Load
	ld.Conections()
}

func (S *Sensor) Evaluar() {
	S.MaxNucleos = runtime.GOMAXPROCS(1)
	S.Cantidad = 1000
}
