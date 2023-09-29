package mdl

import "fmt"

type Load struct {
}

func (L *Load) Conections() {
	var cn Connect
	url := "http://localhost:8080/devel/api/conexiones"
	cn.Conection("GET", url, nil)
	fmt.Println(cn.Data)

	L.TestinFnx()
}

func (L *Load) TestinFnx() {
	var jsonData = []byte(`{
	"funcion": "Fnx_ContraValor",
	"tasa": "16.00",
	"monto_divisa": "500.00",
	"pcm": "0.10"
}`)
	var cn Connect
	url := "http://localhost:8080/devel/api/fnx"
	cn.Conection("POST", url, jsonData)
	fmt.Println(cn.Data)
}
