package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type user struct {
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

func marshalJSON() {
	userExample := user{"Ed", 32}

	userInJSON, err := json.Marshal(userExample)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bytes.NewBuffer(userInJSON))
}

func unmarshalJSON() {
	userInJSON := `{"name":"Ed","age":32}`
	var userExample user

	err := json.Unmarshal([]byte(userInJSON), &userExample)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(userExample)
}

func main() {
	marshalJSON()
	unmarshalJSON()
}
