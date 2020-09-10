package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func typeConversion(value string) interface{} {
	if s, err := strconv.Atoi(value); err == nil {
		return s
	}

	return value
}

func Decoder(value []string) []interface{} {
	decodedRow := make([]interface{}, len(value))
	for i := range value {
		decodedRow[i] = typeConversion(value[i])
	}

	return decodedRow
}

func main() {
	f, err := os.Open("dataset.csv")
	if err != nil {
		fmt.Println("Error open file - ", err)
	}
	defer f.Close()
	count := 0
	var data [][]interface{}

	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'
	Read:
	for {
		row, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				err = nil
				break Read
			}
			fmt.Println("Error read file - ", err)
		}
		count ++
		data = append(data, Decoder(row))
	}
	fmt.Println(count)

	file, err := os.Create("result.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		strData := make([]string, len(value))
		for i := range value {
			strData[i] = fmt.Sprint(value[i])
		}

		err := writer.Write(strData)
		if err != nil {
			log.Fatal(err)
		}
	}
}
