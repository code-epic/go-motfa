package main

import (
	"fmt"
	"go-motfa/mdl"
)

func main() {
	var ss mdl.Sensor
	var bloque mdl.Bloques

	canal := make(chan int)

	bloque.Iniciar()
	ss.Evaluar()

	fmt.Println("Pudes ejecutar un maximo de ", ss.MaxCore)

	go bloque.Inferir(ss.Count, true, canal)
	go bloque.Inferir(ss.Count, false, canal)

	x, y := <-canal, <-canal
	fmt.Println("Resultado del canal", x+y)

	fmt.Println("Finalizando")

}
