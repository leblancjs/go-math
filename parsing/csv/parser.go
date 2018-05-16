package csv

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strconv"
)

type Parser struct {
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(name string) (File, error) {
	f, err := os.Open("resources/log330_labo1_data.csv")
	if err != nil {
		return nil, errors.New("parser: failed to open CSV file")
	}
	defer f.Close()

	r := csv.NewReader(f)

	record, err := r.Read()
	if err != nil {
		return nil, errors.New("parser: failed to read size record")
	}

	size, err := strconv.ParseInt(record[0], 10, 64)
	if err != nil {
		return nil, errors.New("parser: failed to convert size to int")
	}

	records := make([][]float64, 0, size)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, errors.New("parser: failed to read record")
		}

		values := make([]float64, len(record))
		for idx, v := range record {
			val, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return nil, errors.New("parser: failed to convert value to float")
			}

			values[idx] = val
		}

		records = append(records, values)
	}

	return records, nil
}
