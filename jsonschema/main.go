package main

import (
	"fmt"
	"path/filepath"

	"github.com/xeipuuv/gojsonschema"
)

func main() {
	filepath, err := filepath.Abs("./schema.json")
	if err != nil {
		panic(err)
	}
	schema := gojsonschema.NewReferenceLoader("file://" + filepath)
	p := Person{
		FirstName: "Naruto",
		LastName:  "Uzumaki",
		Age:       18,
	}
	actual := gojsonschema.NewGoLoader(p)
	r, err := gojsonschema.Validate(schema, actual)
	if err != nil {
		panic(err)
	}
	if r.Valid() {
		fmt.Println("No problem")
	} else {
		for _, desc := range r.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       uint   `json:"age"`
}
