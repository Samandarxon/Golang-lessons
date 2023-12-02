package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func main() {

	f, err := os.Open("product.json")
	if err != nil {
		log.Println("File open error:", err.Error())
		return
	}

	body, err := io.ReadAll(f)
	if err != nil {
		log.Println("Ioutil read all error:", err.Error())
		return
	}

	var productData interface{}
	err = json.Unmarshal(body, &productData)
	if err != nil {
		log.Println("body unmarshal error:", err.Error())
		return
	}

	var product = make(map[string]interface{})

	if _, ok := productData.(map[string]interface{}); ok {
		product = productData.(map[string]interface{})
	}

	product["metadata"] = map[string]interface{}{
		"Name": "Asadbek",
	}

	body, err = json.MarshalIndent(product, "", "    ")
	if err != nil {
		log.Println("body marshal error:", err.Error())
		return
	}

	err = os.WriteFile("product.json", body, os.ModePerm)
	if err != nil {
		log.Println("WriteFile error:", err.Error())
		return
	}
}
