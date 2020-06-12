package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

type Character struct {
	Name string `json:"name"`
}

const url = "https://api.got.show/api/book/characters"

func main() {
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var characters []Character

	err = json.Unmarshal(body, &characters)
	if err != nil {
		panic(err.Error())
	}

	rand.Seed(time.Now().Unix())
	c := characters[rand.Intn(len(characters))]

	fmt.Println(c.Name)
}
