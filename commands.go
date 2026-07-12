package main

import (
	"errors"
	"fmt"
	"strconv"
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

func stats(ds Dataset) {

	for colIdx := range ds.Header {
		// metric vars
		var min float64
		var max float64
		var sum float64
		var count int
		isNumeric := true
		seenFirst := false

		for _, row := range ds.Rows {

			// acummulate metrics
			cell, err := strconv.ParseFloat(row[colIdx], 64)
			if err != nil {
				isNumeric = false
			} else {
				if !seenFirst {
					min = cell
					max = cell
					seenFirst = true
				} else {
					if cell < min {
						min = cell
					}
					if cell > max {
						max = cell
					}
				}
				sum += cell
			}
			count++
		}

		// print stats for row
		fmt.Print(ds.Header[colIdx])
		fmt.Printf("{ count: %d", count)
		if isNumeric {
			avg := sum / float64(count)
			fmt.Printf(", min: %.0f, max: %.0f, avg: %.2f}\n", min, max, avg)
		} else {
			fmt.Println(" }")
		}

	}

}
