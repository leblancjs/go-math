package main

import (
	"fmt"
	"github.com/leblancjs/log330-math-go/math/stats"
	"github.com/leblancjs/log330-math-go/parsing/csv"
	"log"
)

func main() {
	p := csv.NewParser()
	f, err := p.Parse("resources/log330_labo1_data.csv")
	if err != nil {
		log.Fatalln(err)
	}

	columnCount := f.RowSize()
	for i := 0; i < columnCount; i++ {
		fmt.Println("Column ", i)

		values := f.Column(i)

		mean := stats.Mean(values)
		variance := stats.Variance(values)
		stdDev := stats.StdDev(values)

		fmt.Println("Mean: ", mean)
		fmt.Println("Variance: ", variance)
		fmt.Println("Standard Deviation: ", stdDev)
	}
}
