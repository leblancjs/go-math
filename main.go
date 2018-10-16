package main

import (
	"fmt"
	"log"

	"github.com/leblancjs/go-math/math/stats"
	"github.com/leblancjs/go-math/parsing/csv"
)

func main() {
	parser := csv.NewParser()
	file, err := parser.Parse("resources/log330_labo1_data.csv")
	if err != nil {
		log.Fatalln(err)
	}

	columnCount := file.RowSize()
	for i := 0; i < columnCount; i++ {
		fmt.Println("Column ", i)

		values := file.Column(i)

		mean := stats.Mean(values)
		variance := stats.Variance(values)
		stdDev := stats.StdDev(values)

		fmt.Println("Mean: ", mean)
		fmt.Println("Variance: ", variance)
		fmt.Println("Standard Deviation: ", stdDev)
	}
}
