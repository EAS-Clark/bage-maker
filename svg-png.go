package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
    "image"
    "image/color"
    "image/draw"
    "github.com/srwiley/rasterx"
    "github.com/srwiley/oksvg"
    "github.com/srwiley/rasterx/png"
)

func main() {
    // Specify the input folder containing SVG files.
    inputFolder := "./input"

    // Specify the output folder where PNG files will be saved.
    outputFolder := "./output"

    // Create the output folder if it doesn't exist.
    if err := os.MkdirAll(outputFolder, os.ModePerm); err != nil {
        fmt.Printf("Error creating output folder: %v\n", err)
        return
    }

    // List all SVG files in the input folder.
    svgFiles, err := filepath.Glob(filepath.Join(inputFolder, "*.svg"))
    if err != nil {
        fmt.Printf("Error listing SVG files: %v\n", err)
        return
    }

    // Iterate through each SVG file and convert it to PNG.
    for _, svgFile := range svgFiles {
        // Open the SVG file.
        reader, err := os.Open(svgFile)
        if err != nil {
            fmt.Printf("Error opening SVG file %s: %v\n", svgFile, err)
            continue
        }
        defer reader.Close()

        // Parse the SVG file.
        svg, _ := oksvg.ReadFrom(reader, oksvg.WarnErrorMode)

        // Create a new image and drawing context.
        width, height := int(svg.ViewBox.W), int(svg.ViewBox.H)
        img := image.NewRGBA(image.Rect(0, 0, width, height))
        gc := draw.New(img)

        // Set a white background.
        draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

        // Rasterize the SVG onto the image.
        svg.Render(rasterx.NewDasher(width, height, rasterx.NewScannerGV(width, height, gc)))

        // Create the output PNG file with the same name in the output folder.
        outputFile := filepath.Join(outputFolder, strings.TrimSuffix(filepath.Base(svgFile), ".svg")+".png")

        // Save the PNG file.
        pngFile, err := os.Create(outputFile)
        if err != nil {
            fmt.Printf("Error creating PNG file %s: %v\n", outputFile, err)
            continue
        }
        defer pngFile.Close()

        // Encode and save the PNG.
        if err := png.Encode(pngFile, img); err != nil {
            fmt.Printf("Error saving PNG file %s: %v\n", outputFile, err)
        } else {
            fmt.Printf("Converted %s to %s\n", svgFile, outputFile)
        }
    }
}



