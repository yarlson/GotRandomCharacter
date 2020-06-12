package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

type Character []struct {
	V         int64    `json:"__v"`
	ID        string   `json:"_id"`
	Locations []string `json:"locations"`
	Name      string   `json:"name"`
	Slug      string   `json:"slug"`
}

const url = "https://api.got.show/api/characters/locations"

func main() {
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	characters := Character{}
	err = json.Unmarshal(body, &characters)
	if err != nil {
		panic(err.Error())
	}

	rand.Seed(time.Now().Unix())
	c := characters[rand.Intn(len(characters))]

	fmt.Println(c.Name)
}
