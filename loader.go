package main

import (
	"encoding/csv"
	"os"
)

type Dataset struct {
	Header []string
	Rows   [][]string
}

func load(filename string) (Dataset, error) {

	// load file
	file, err := os.Open(filename)
	if err != nil {
		return Dataset{}, err
	}
	defer file.Close()

	// read in file
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return Dataset{}, err
	}

	// create Dataset
	ds := Dataset{Header: records[0], Rows: records[1:]}

	return ds, nil
}
