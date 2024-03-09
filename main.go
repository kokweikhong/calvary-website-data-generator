package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Product struct {
	Name        string   `json:"name"`
	Title       string   `json:"title"`
	Countries   []string `json:"countries"`
	Services    []string `json:"services"`
	Thumbnail   string   `json:"thumbnail"`
	Logo        string   `json:"logo"`
	Description string   `json:"description"`
}

const (
	PRODUCTS_JSON_OUTPUT = "output/products.json"
)

func main() {
	// read csv file
	filePath := "csv/products.csv"
	// read csv file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read csv file
	reader := csv.NewReader(file)
	// read csv file
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// read csv file
	for _, record := range records {
		fmt.Println(record)
	}

	generateProductsJSON(records)
}

func generateProductsJSON(csvData [][]string) error {
	if len(csvData) == 0 {
		return fmt.Errorf("no data in csv")
	}

	products := []*Product{}
	for recordsIdx, records := range csvData {
		if recordsIdx == 0 {
			// skip header
			continue
		}
		product := new(Product)
		for recordIdx, record := range records {
			switch strings.ToLower(csvData[0][recordIdx]) {
			case "name":
				product.Name = record
			case "title":
				product.Title = record
			case "countries":
				countries := strings.Split(record, ",")
				for i, country := range countries {
					countries[i] = strings.TrimSpace(country)
				}
				product.Countries = countries
			case "services":
				services := strings.Split(record, ",")
				for i, service := range services {
					services[i] = strings.TrimSpace(service)
				}
				product.Services = services
			case "thumbnail":
				product.Thumbnail = record
			case "logo":
				product.Logo = record
			case "description":
				product.Description = record
			}
		}

		products = append(products, product)
	}

	jsonData, err := json.MarshalIndent(products, "", "  ")
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(PRODUCTS_JSON_OUTPUT), 0755)
	if err != nil {
		return err
	}

	// create file if not exists
	file, err := os.OpenFile(PRODUCTS_JSON_OUTPUT, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}

	return nil

}
