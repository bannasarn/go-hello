package basicsyntax

import (
	"encoding/json"
	"fmt"
	"log"
)

type country struct {
	Name       string `json:"name"`
	Alpha2Code string `json:"alpha2_code"`
	Alpha3Code string `json:"alpha3_code"`
}

func Struct() {
	var c country
	jsonString := `{
		"name": "India",
		"alpha2_code": "IND2",
		"alpha3_code": "IND3"
	}`
	err := json.Unmarshal([]byte(jsonString), &c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c)
}
