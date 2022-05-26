package main

import (
	"encoding/csv"
	"fmt"
	"go/token"
	"go/types"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {

	// Open file, define pointer to it
	file, err := os.Open("problems.csv")
	if err != nil {
		fmt.Println(err)
	}
	// CSV package's pointer to opened file
	csvReader := csv.NewReader(file)

	// This reads the whole file. Is it like readlines on Python? (BAD for memory usage?)
	//records, _ := csvReader.ReadAll()

	// while true, but nicer
	for {
		// line by line, see `errorResult` variable
		record, errorResult := csvReader.Read()

		if errorResult == io.EOF {
			break
		}
		// If there's an error and it's not the end of file, it dies
		if errorResult != nil {
			log.Fatal(errorResult)
		}

		// Printing the question
		fmt.Printf("%s\n", record[0])

		// A tokenizer is instantiated
		fs := token.NewFileSet()

		// Eval, maybe less evil than PHP or JS eval. See how we send errors to _
		tv, _ := types.Eval(fs, nil, token.NoPos, record[0])

		// The result of evaluating the expression!
		//fmt.Println(tv.Type)
		fmt.Println(tv.Value)

		// This must be cast/assigned here. Atoi returns the int AND the error callback, so it can't be done within the IF validation :/
		answer, _ := strconv.Atoi(record[1])
		calculated, _ := strconv.Atoi(tv.Value.String())

		// Check if answer is correct
		if answer == calculated {
			fmt.Println("Punto bien")
		} else {
			fmt.Println("Error")
		}

	}
	// You can print the whole CSV if you used readAll()
	//fmt.Println(records)
}
