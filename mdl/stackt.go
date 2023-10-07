package mdl

import (
	"encoding/json"
	"fmt"
)

type Stack struct {
	IdHash string
	Nodes  []Node
}

type Node struct {
	Name      string
	Type      string
	QueryPlan string //MD5 Plan de ejecucion
	Status    bool
	Size      int
	Rule      string
}

func (s *Stack) Get() {
}
func (s *Stack) Update() {
}

func (s *Stack) Add(node Node) {
	s.Nodes = append(s.Nodes, node)
}

func (s *Stack) Delete() {
}

func (s *Stack) Commit() {

	stack_json, _ := json.Marshal(s)
	fmt.Println(string(stack_json))
	var jsx = []byte(`{
			"coleccion": "motfa-api",
			"objeto": ` + string(stack_json) + `,
			"driver": "MGDBA",
			"upsert": true,
			"donde": "{\"oid\":\"1\"}"
		}`)

	cn.Conection("POST", URL+"ccoleccion", jsx)
	fmt.Println("Collection:  ", cn.Data)
}
