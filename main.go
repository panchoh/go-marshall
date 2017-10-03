package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name string `json:"name"`
}

func main() {
	var sentPeople = []Person{Person{"Rai"}, Person{"Alex"}}
	var (
		payload []byte
		err     error
	)
	payload, err = json.Marshal(sentPeople)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(string(payload))

	var receivedPeople []Person
	if err = json.Unmarshal(payload, &receivedPeople); err != nil {
		log.Panic(err)
	}

	for _, p := range receivedPeople {
		fmt.Println(p.Name)
	}
}
