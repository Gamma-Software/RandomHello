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

func ReadHellosEntries() []Hello {
	jsonFile, err := ioutil.ReadFile("hello/hellos.json")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened hellos.json")

	var hellos []Hello
	err = json.Unmarshal([]byte(jsonFile), &hellos)

	if err != nil {
		fmt.Println("error:", err)
	}

	return hellos
}

func RandomHellos() Hello {
	// Read json files from the directory
	hellos := ReadHellosEntries()

	// Get a random number between 0 and the length of the array
	randomNumber := rand.Intn(len(hellos))

	// Return the random hello
	return hellos[randomNumber]
}
