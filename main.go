package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("resources/log330_labo1_data.csv")
	if err != nil {
		log.Fatal("Failed to open CSV")
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal("Failed to read CSV")
	}

	fmt.Println(records)
}
