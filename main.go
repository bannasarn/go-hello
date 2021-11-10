package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	exerciseOscar()
}

func exerciseOscar() {
	// Open File
	f, err := os.Open("/Users/bannasarn/Downloads/oscar.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Read File
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Map record into name
	name := make(map[string]int)
	for _, record := range records {
		name[record[3]] += 1
	}

	// Filter value more than 1
	for k, v := range name {
		if v > 1 {
			println(k)
		}
	}
}
