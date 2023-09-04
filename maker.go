package main

import (
    "encoding/csv"
    "fmt"
    "os"
    "strings"
)

func main() {
    // Create a "badges" folder if it doesn't exist
    err := os.MkdirAll("out", os.ModePerm)
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

    // Loop through the CSV records and generate SVGs
    for _, record := range records {
        if len(record) != 2 {
            fmt.Println("Invalid CSV record:", record)
            continue
        }

        // Extract data from CSV
        label := record[0]
        status := record[1]

        // Generate SVG based on the template
        svg := generateSVG(label, status)

        // Create the filename for the SVG
        filename := fmt.Sprintf("out/%s-%s.svg", label, status)

        // Create and write the SVG to the file
        svgFile, err := os.Create(filename)
        if err != nil {
            fmt.Println("Error creating SVG file:", err)
            continue
        }
        defer svgFile.Close()

        _, err = svgFile.WriteString(svg)
        if err != nil {
            fmt.Println("Error writing SVG:", err)
            continue
        }

        fmt.Printf("Generated %s\n", filename)
    }
}

func generateSVG(label, status string) string {
    svgTemplate := `
<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="108" height="20" role="img" aria-label="{{LABEL}}: {{STATUS}}"><title>{{LABEL}}: {{STATUS}}</title><linearGradient id="s" x2="0" y2="100%"><stop offset="0" stop-color="#bbb" stop-opacity=".1"/><stop offset="1" stop-opacity=".1"/></linearGradient><clipPath id="r"><rect width="108" height="20" rx="3" fill="#fff"/></clipPath><g clip-path="url(#r)"><rect width="59" height="20" fill="#555"/><rect x="59" width="49" height="20" fill="#4caf50"/><rect width="108" height="20" fill="url(#s)"/></g><g fill="#fff" text-anchor="middle" font-family="Verdana,Geneva,DejaVu Sans,sans-serif" text-rendering="geometricPrecision" font-size="110"><text aria-hidden="true" x="305" y="150" fill="#010101" fill-opacity=".3" transform="scale(.1)" textLength="490">{{LABEL}}</text><text x="305" y="140" transform="scale(.1)" fill="#fff" textLength="490">{{LABEL}}</text><text aria-hidden="true" x="825" y="150" fill="#010101" fill-opacity=".3" transform="scale(.1)" textLength="390">{{STATUS}}</text><text x="825" y="140" transform="scale(.1)" fill="#fff" textLength="390">{{STATUS}}</text></g></svg>
`

    // Replace placeholders with actual data
    svg := strings.ReplaceAll(svgTemplate, "{{LABEL}}", label)
    svg = strings.ReplaceAll(svg, "{{STATUS}}", status)

    return svg
}
