package mdl

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

var (
	cn    Connect
	api   API
	stack Stack
)

type EjectAPI struct {
	Function   string `json:"funcion"`
	Parameters string `json:"parametros"`
	PreCode    string `json:"precodigo"`
}

type Load struct {
}

func init() {

	api.GetAll()
	url := URL + "conexiones"
	cn.Conection("GET", url, nil)

	// api.List()
}
func (E *EjectAPI) ToJson() []byte {
	jSon, _ := json.Marshal(E)
	return jSon
}

// Analyzed sistem API Explain for Query Plan
func (L *Load) Analyzed(ctx context.Context) {
	var eject EjectAPI
	eject.PreCode = "EXPLAIN"
	eject.Parameters = ""
	url := URL + "crud:541f2bd069277d053a45cac13099e185.sse"
	for _, v := range api.Data {
		eject.Function = v.Funcion
		cn.Conection("POST", url, eject.ToJson())
		L.ProcessAPI(v, &cn.DataByte)
	}
	stack.Commit()
}

func (L *Load) ProcessAPI(api ApiCore, rsData *[]byte) {
	var rs RS
	_ = json.Unmarshal(*rsData, &rs)
	var node Node
	node.Name = api.Funcion
	h := sha256.New()
	h.Write(cn.DataByte)
	node.QueryPlan = hex.EncodeToString(h.Sum(nil))
	node.Type = api.Metodo
	node.Status = false
	if len(rs.Cabecera) > 0 {
		node.Status = true
	}
	stack.Add(node)
	//Control de flujo en caso de fallar la API
}

func (L *Load) ListApi() {
	for _, v := range stack.Nodes {
		fmt.Println(v.Name)
	}
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
