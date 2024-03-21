package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	PRODUCTS_JSON_OUTPUT = "output/products.json"
)

type Product struct {
	Name        string   `json:"name"`
	Title       string   `json:"title"`
	Href        string   `json:"href"`
	Countries   []string `json:"countries"`
	Services    []string `json:"services"`
	Thumbnail   string   `json:"thumbnail"`
	Logo        string   `json:"logo"`
	Description string   `json:"description"`
}

func GenerateProductsJSON(csvData [][]string, output string) error {
	if len(csvData) == 0 {
		return fmt.Errorf("no data in csv")
	}

	if output == "" {
		output = PRODUCTS_JSON_OUTPUT
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
			case "href":
				product.Href = record
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

	err = os.MkdirAll(filepath.Dir(output), 0755)
	if err != nil {
		return err
	}

	// create file if not exists
	file, err := os.OpenFile(output, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}

	return nil

}
