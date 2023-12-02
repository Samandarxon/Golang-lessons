package main

import (
	"encoding/json"
	"log"
)

type Doctor struct {
	Id      string  `json:"-"`
	Name    string  `json:"name"`
	Surname string  `json:"surname"`
	Phone   string  `json:"phone"`
	Weight  float64 `json:"wieght,omitempty"`
	Status  bool
}

func main() {

	var d1 = Doctor{
		Id:      "1",
		Name:    "Alex",
		Surname: "Bob",
		Phone:   "9993-2332-23-23",
	}

	body, err := json.Marshal(d1)
	if err != nil {
		log.Println(body)
		return
	}

	log.Println(string(body))
}
