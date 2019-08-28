package csvread

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// Ex reads a .csv file and displays the contents.
func Ex(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	r := csv.NewReader(file)
	r.Comment = '#'
	r.Comma = ';' // if it is another delimiter

	// records, err := r.ReadAll()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(records)

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			if pe, ok := err.(*csv.ParseError); ok {
				fmt.Println("bad column:", pe.Column)
				fmt.Println("bad line:", pe.Line)
				fmt.Println("Error reported", pe.Err)
				if pe.Err == csv.ErrFieldCount {
					continue
				}
			}
			log.Fatal(err)
		}
		fmt.Println("CSV Row:", record)

		i, err := strconv.Atoi(record[2])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("This is 4 times age:", i*4)

	}

}
