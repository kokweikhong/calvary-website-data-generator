package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

type Project struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Year        string   `json:"year"`
	Location    string   `json:"location"`
	Completion  string   `json:"completion"`
	Size        string   `json:"size"`
	ImagePath   string   `json:"imagePath"`
	Country     string   `json:"country"`
	Tags        []string `json:"tags"`
	Sectors     []string `json:"sectors"`
	Services    []string `json:"services"`
	Description string   `json:"description"`
	Products    []string `json:"products"`
	Images      []string `json:"images"`
}

func GenerateProjectsJSON(csvData [][]string, output string) error {
	var projects []*Project

	for recordsIdx, records := range csvData {
		if recordsIdx == 0 {
			// skip header
			continue
		}
		project := new(Project)
		for recordIdx, record := range records {
			switch strings.ToLower(csvData[0][recordIdx]) {
			case "id":
				project.Id = record
			case "name":
				project.Name = record
			case "year":
				project.Year = record
			case "location":
				project.Location = record
			case "completion":
				project.Completion = record
			case "size":
				project.Size = record
			case "image_path":
				project.ImagePath = record
			case "country":
				project.Country = record
			case "tags":
				tags := strings.Split(record, ",")
				for i, tag := range tags {
					tags[i] = strings.TrimSpace(tag)
				}
				project.Tags = tags
			case "sectors":
				sectors := strings.Split(record, ",")
				for i, sector := range sectors {
					sectors[i] = strings.TrimSpace(sector)
				}
				project.Sectors = sectors
			case "services":
				services := strings.Split(record, ",")
				for i, service := range services {
					services[i] = strings.TrimSpace(service)
				}
				project.Services = services
			case "description":
				project.Description = record
			case "products":
				products := strings.Split(record, ",")
				for i, product := range products {
					products[i] = strings.TrimSpace(product)
				}
				project.Products = products
			}
		}
		images, err := ListObjects(project.ImagePath)
		if err != nil {
			return err
		}
		// remove images without ext
		for i, image := range images {
			if !strings.Contains(image, ".") {
				images = append(images[:i], images[i+1:]...)
			}
		}
		project.Images = images
		projects = append(projects, project)
	}

	jsonData, err := json.MarshalIndent(projects, "", "  ")
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
