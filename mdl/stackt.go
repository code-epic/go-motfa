package mdl

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Stack struct {
	IdHash string
	Nodes  []Node
}

type Node struct {
	Name      string
	Type      string
	Map       string
	QueryPlan string //MD5 Plan de ejecucion
	Status    bool
	Size      int
	Rule      string
	Log       string
}

// Get Obtener la pila de las API
func (s *Stack) Get(idNode string) (err error, node Node) {
	for _, v := range s.Nodes {
		if v.Name == idNode {
			return nil, v
		}
	}
	return errors.New("No se econtro el elemento"), node
}

// Update Actualizar los datos de la pila
func (s *Stack) Update() {

}

// Add agregar elementos al pila
func (s *Stack) Add(node Node) error {

	for _, v := range s.Nodes {
		if v.Name == node.Name {
			return errors.New("El registro ya esta en la pila")
		}
	}
	s.Nodes = append(s.Nodes, node)
	return nil
}

func (s *Stack) Delete() {
}

func (s *Stack) Commit() {
	if MT_ACTIVE_COMMIT {
		stack_json, _ := json.Marshal(s)
		var jsx = []byte(`{
			"coleccion": "motfaapi",
			"objeto": ` + string(stack_json) + `,
			"driver": "` + MT_COMMIT + `",
			"upsert": true,
			"donde": "{\"oid\":\"1\"}"
		}`)

		cn.Conection("POST", URL+"ccoleccion", jsx)
		fmt.Println("Collection:  ", cn.Data)
	}
}

func (s *Stack) Print() {
	for _, v := range s.Nodes {
		fmt.Println(v.Name, "\t", v.Status)
	}
}
