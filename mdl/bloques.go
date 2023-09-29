package mdl

type Bloques struct {
	Racionalizar bool
	Controlar    bool
}

type Funciones struct {
	Automatizar bool
	Precision   float64
}

func (B *Bloques) Iniciar() {
}

func (B *Bloques) Inferir(cantidad int, decision bool, c chan int) {
	contador := 0

	if decision {
		for i := 0; i < cantidad; i++ {
			//fmt.Println("Aumenta ", i)
			contador++
		}
	} else {
		for i := cantidad; i > 0; i-- {
			//fmt.Println("Decrementa ", i)
			contador++
		}
	}
	c <- contador

}
