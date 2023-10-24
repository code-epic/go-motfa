package mdl

import (
	"encoding/json"
	"fmt"
	"time"
)

// Definicion Describe una oficinas
type Definicion struct {
	Nombre string `json:"nombre"`
	Tipo   string `json:"tipo"`
}

// RS Establece la respuesta de la API
type RS struct {
	Cabecera []Definicion `json:"Cabecera"`
	Cuerpo   interface{}  `json:"Cuerpo"`
	Pie      interface{}  `json:"Pie"`
}

// RS Establece la respuesta de la API
type RST struct {
	Msj         string `json:"msj"`
	Tipo        int    `json:"tipo"`
	Fecha       string `json:"fecha"`
	Responsable int    `json:"responsable"`
	Contenido   string `json:"contenido"`
}

// ApiCore Estructura de conexion
type ApiCore struct {
	ID            string      `json:"id"`
	Modulo        string      `json:"modulo"`
	Driver        string      `json:"driver"`
	Tipo          string      `json:"tipo"`
	PreCodigo     string      `json:"precodigo"`
	Coleccion     string      `json:"coleccion"`
	Accion        bool        `json:"accion"` //Definir Insercion C.Traza.F Aplicar Reglas de injeccion
	Query         string      `json:"query"`
	Parametros    string      `json:"parametros"`
	Totalizar     string      `json:"totalizar"`
	Columnas      string      `json:"columnas"`
	Ruta          string      `json:"ruta"`
	Funcion       string      `json:"funcion"`
	Retorna       bool        `json:"retorna"`
	Concurrencia  bool        `json:"concurrencia"`
	Migrar        bool        `json:"migrar"`
	Metodo        string      `json:"metodo"`
	Destino       string      `json:"destino"`
	PuertoHttp    int         `json:"puertohttp"`
	PuertoHttps   int         `json:"puertohttps"`
	Protocolo     string      `json:"protocolo"`
	Consumidores  int         `json:"consumidores"`
	Prioridad     string      `json:"prioridad"`
	Entorno       string      `json:"entorno"`
	Estatus       bool        `json:"estatus" bson:"estatus,omitempty"`
	Relacional    bool        `json:"relacional"`
	Valores       interface{} `json:"valores" bson:"valores,omitempty"`
	Logs          bool        `json:"logs"`
	Cache         int         `json:"cache"`
	Descripcion   string      `json:"descripcion"`
	Version       string      `json:"version"`
	Categoria     string      `json:"categoria"`
	Entradas      interface{} `json:"entradas"`
	Salidas       interface{} `json:"salidas"`
	Funcionalidad string      `json:"funcionalidad"`
	Resultado     interface{} `json:"resultado"`
	Fecha         time.Time   `json:"fecha"`
	Autor         string      `json:"autor"`
	Aplicacion    string      `json:"aplicacion"`
	Distribucion  int         `json:"distribucion"`
	Duracion      int         `json:"duracion"`
	TipoDuracion  int         `json:"tipoduracion"` //0 Segundos, 1 Minutas, 2 Horas
}

type ExecApi struct {
	Function   string      `json:"funcion"`
	Parameters string      `json:"parametros"`
	Values     interface{} `json:"valores"`
	PreCode    string      `json:"precodigo"`
}

// "parametros": "PGMOTFA"
type API struct {
	Data   []ApiCore
	Result []RS
	Stack  Stack
	Size   int
}

func (E *ExecApi) ToJson() []byte {
	jSon, _ := json.Marshal(E)

	//fmt.Println(string(jSon))
	return jSon
}

// getDefApi Metodo privado de generar la definicion de APICore
func getDefApi() []byte {
	var fnx ExecApi
	fnx.Function = FNX_GetAPI
	fnx.Parameters = "MYCFEPD"
	jsonP, err := json.Marshal(fnx)

	if err != nil {
		fmt.Println("Error en el formato ", err.Error())
	}
	return jsonP
}

// GetAll Obtener Todas las api de un driver
func (a *API) GetAll() {
	var cn Connect
	url := URL + "crud:541f2bd069277d053a45cac13099e185.sse"
	cn.Conection("POST", url, getDefApi())

	err := json.Unmarshal(cn.DataByte, &a.Data)
	if err != nil {
		fmt.Println("Error en el formato ", err.Error())
	}
	a.Size = len(a.Data)
}

// GetSql Obtener listado de SQL
func (a *API) GetSql() {
	var cn Connect

	url := URL + "crud:541f2bd069277d053a45cac13099e185.sse"
	cn.Conection("POST", url, getDefApi())
	//fmt.Println(cn.Data)
	err := json.Unmarshal(cn.DataByte, &a.Result)
	if err != nil {
		fmt.Println("Error en el formato ", err.Error())
	}
	a.Size = len(a.Result)
}

func (a *API) Print() {
	for _, v := range a.Data {
		fmt.Println(v.Funcion, v.Metodo)
	}
}
