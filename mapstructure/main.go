package main

import (
	"fmt"
	"log"

	"github.com/mitchellh/mapstructure"
)

type Person struct {
	Name   string
	Age    int
	Emails []string
}

func main() {
	m := map[string]interface{}{
		"name": 123,
		"age":  "28",
		"job":  "programmer",
	}

	var p Person
	var metadata mapstructure.Metadata
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Metadata:         &metadata,
		Result:           &p,
		MatchName:        nil,
	})
	if err != nil {
		log.Fatal(err)
	}

	err = decoder.Decode(m)
	if err == nil {
		fmt.Println("person:", p)
		fmt.Printf("keys:%#v, unused:%#v\n", metadata.Keys, metadata.Unused)
	} else {
		fmt.Println(err.Error())
	}
}
