package oscar

import (
	"encoding/csv"
	"log"
	"os"
)

// Find the name who get the oscar more than one time

func Oscar() {
	// Open File
	f, err := os.Open("./oscar/oscar.csv")
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
		name[record[3]]++
	}

	// // Filter value more than 1
	for k, v := range name {
		if v > 1 {
			println(k)
		}
	}
}
