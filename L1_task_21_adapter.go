// Адаптер — это структурный паттерн проектирования, 
// который позволяет объектам с несовместимыми интерфейсами работать вместе.

// Адаптер выступает прослойкой между двумя объектами, превращая вызовы одного в вызовы понятные другому.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"github.com/go-yaml/yaml"
)

type Order struct {
	ID string `yaml:"ID"`
	SupplierOrderCode string `yaml:"SupplierOrderCode"`
}

type Orders struct {
	MyOrders []Order `yaml:"Orders"`
}

func YamlToJson(data []byte) []byte{
	var yamlContent Orders
	err := yaml.Unmarshal(data, &yamlContent)
	if err != nil {
		log.Fatal("yaml.Unmarshal didn't succeed ", err)
	}
	a, err := json.Marshal(yamlContent)
	if err != nil{
		log.Fatal(err)
	}
	return a
}

func main(){
	// ADAPTER YAML->JSON
	yamlBytes, err := ioutil.ReadFile("L1_task_21.yaml")
	if err != nil {
		log.Fatal("File didn't get opened")
	}
	JsonBytes := YamlToJson(yamlBytes)
	fmt.Println(string(JsonBytes))
}

