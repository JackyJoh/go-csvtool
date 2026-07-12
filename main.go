package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strconv"
	"strings"
)

func printRows(header []string, rows [][]string) {
	fmt.Println(strings.Join(header, ", "))
	for _, row := range rows {
		fmt.Println(strings.Join(row, ", "))
	}
}

func main() {

	// load dataset
	// fallback on failed load to default csv
	ds, err := load(os.Args[2])
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			fmt.Println("file not found, using default: kalshi_crypto.csv")
			ds, err = load("kalshi_crypto.csv")
			if err != nil {
				fmt.Println("fallback also failed:", err)
				return
			}
		} else {
			fmt.Println("error:", err)
			return
		}
	}

	switch os.Args[1] {
	case "stats":
		// call commands stats

	case "head":
		n := 10 // default 10 rows
		if len(os.Args) == 5 && os.Args[3] == "--n" {
			tempN, err := strconv.Atoi(os.Args[4])
			if err != nil {
				fmt.Println("error parsing n:", err)
				return
			}
			n = tempN
		}
		headRows, err := head(ds, n)
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		printRows(ds.Header, headRows)

	case "col":
		var name string
		if len(os.Args) == 5 && os.Args[3] == "--name" {
			name = os.Args[4]
		}
		index, err := colIndex(ds.Header, name)
		if err != nil {
			fmt.Println("error getting header index:", err)
		}

	}

	fmt.Println(os.Args)
}
