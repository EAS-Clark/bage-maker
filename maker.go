package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	// Create a "badges" folder if it doesn't exist
	err := os.MkdirAll("badges", os.ModePerm)
	if err != nil {
		fmt.Println("Error creating 'badges' folder:", err)
		return
	}

	// Open the CSV file
	csvfile, err := os.Open("input.csv")
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer csvfile.Close()

	// Parse the CSV
	csvreader := csv.NewReader(csvfile)
	records, err := csvreader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	// Loop through the CSV records and generate SVG badges
	for _, record := range records {
		if len(record) != 2 {
			fmt.Println("Invalid CSV record:", record)
			continue
		}

		// Extract data from CSV
		label := record[0]
		status := record[1]

		// Make a web request to Shields.io to generate the SVG badge
		url := fmt.Sprintf("https://img.shields.io/badge/%s-%s-brightgreen.svg", label, status)

		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error making the request:", err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Println("Received a non-OK response:", resp.Status)
			continue
		}

		badgeData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			continue
		}

		// Create the filename for the SVG badge
		filename := fmt.Sprintf("output/%s-%s.svg", label, status)

		// Write the SVG badge to the file
		err = ioutil.WriteFile(filename, badgeData, 0644)
		if err != nil {
			fmt.Println("Error writing SVG badge to file:", err)
			continue
		}

		fmt.Printf("Generated %s\n", filename)
	}
}
