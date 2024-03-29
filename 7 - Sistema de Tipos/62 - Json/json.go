package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	// struct para json
	p1 := produto{1, "Notebook", 3567.6, []string{"Promoção", "Eletrônico"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	// Json para struct
	var p2 produto
	p2Json := `{"id":2,"nome":"Caneta", "preco":1.25, "tags":["Papelaria", "Importado"]}`
	json.Unmarshal([]byte(p2Json), &p2)
	fmt.Println(p2)
}
