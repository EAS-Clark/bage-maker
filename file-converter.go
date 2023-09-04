package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

func main() {
    rootDir := "input" // Replace with your directory path

    err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        // Check if the directory or filename contains %20
        if strings.Contains(info.Name(), "%20") {
            newName := strings.Replace(info.Name(), "%20", "-", -1)
            newPath := filepath.Join(filepath.Dir(path), newName)

            if err := os.Rename(path, newPath); err != nil {
                fmt.Printf("Error renaming %s to %s: %v\n", path, newPath, err)
            } else {
                fmt.Printf("Renamed %s to %s\n", path, newPath)
            }
        }
        return nil
    })

    if err != nil {
        fmt.Printf("Error walking the directory: %v\n", err)
    }
}
