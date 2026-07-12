package main

import (
	"errors"
)

func head(ds Dataset, n int) ([][]string, error) {

	if len(ds.Rows) < n {
		n = len(ds.Rows)
	}

	return ds.Rows[:n], nil
}

func colIndex(header []string, name string) (int, error) {
	for i, cName := range header {
		if name == cName {
			return i, nil
		}
	}
	return 0, errors.New("column " + name + " not found in csv")
}

func cols(ds Dataset, index int) ([]string, error) {

	if index < 0 || index >= len(ds.Header) {
		return nil, errors.New("column index out of range")
	}

	var col []string
	for _, row := range ds.Rows {
		col = append(col, row[index])
	}

	return col, nil
}
