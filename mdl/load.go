package mdl

import (
	"context"
	"encoding/json"
	"fmt"
)

var (
	cn    Connect
	api   API
	stack Stack
)

type Load struct {
}

func init() {

	api.GetAll()
	url := URL + "conexiones"
	cn.Conection("GET", url, nil)

}

// Analyzed sistem API Explain for Query Plan
func (L *Load) Analyzed(ctx context.Context) {
	var eject ExecApi
	eject.PreCode = "EXPLAIN"
	url := URL + "crud:541f2bd069277d053a45cac13099e185.sse"
	for _, v := range api.Data {
		eject.Function = v.Funcion
		eject.Parameters = v.Parametros
		if v.Metodo == "CONSULTAR" {
			cn.Conection("POST", url, eject.ToJson())
			L.ProcessAPI(v, &cn.DataByte)
		}
	}
	stack.Commit()
	stack.Print()
}

func (L *Load) ProcessAPI(api ApiCore, rsData *[]byte) {
	var rs RST
	err := json.Unmarshal(*rsData, &rs)
	if err != nil {
		fmt.Println(err.Error())
	}

	var node Node
	node.Name = api.Funcion
	node.QueryPlan = SetHash(cn.DataByte)
	node.Type = api.Metodo
	node.Status = true
	node.Log = rs.Msj
	if rs.Msj != "" {
		node.Status = false
	}
	_ = stack.Add(node)
	//Control de flujo en caso de fallar la API
}

func (L *Load) SendData() {

}

func (L *Load) TestinFnx() {
	var jsonData = []byte(`{
"funcion": "Fnx_ContraValor",
"tasa": "16.00",
"monto_divisa": "500.00",
"pcm": "0.10"
}`)
	var cn Connect
	url := URL + "fnx"
	cn.Conection("POST", url, jsonData)
	fmt.Println(cn.Data)
}
