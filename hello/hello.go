package hello

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
)

type Hello struct {
	Language string `json:"language"`
	Hello    string `json:"hello"`
}

var hellos []Hello

func init() {
	data, err := ioutil.ReadFile("hello/hellos.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(data), &hellos)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully opened hellos.json")
}

func RandomHellos() Hello {
	// Get a random number between 0 and the length of the array
	randomNumber := rand.Intn(len(hellos))

	// Return the random hello
	return hellos[randomNumber]
}
