package main

import (
	"encoding/csv"
	"os"
)

func ReadCSVFile(csvPath string) ([][]string, error) {
	filePath := "csv/products.csv"
	if csvPath != "" {
		filePath = csvPath
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil

	// read csv file
	// for _, record := range records {
	// 	fmt.Println(record)
	// }

	// generateProductsJSON(records)
}
