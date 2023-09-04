package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

func main() {
    rootDir := "/path/to/your/directory" // Replace with your directory path

    err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        if info.IsDir() {
            // Check if the directory name contains %20
            if strings.Contains(info.Name(), "%20") {
                newDirName := strings.Replace(info.Name(), "%20", "-", -1)
                newDirPath := filepath.Join(filepath.Dir(path), newDirName)
                if err := os.Rename(path, newDirPath); err != nil {
                    fmt.Printf("Error renaming %s to %s: %v\n", path, newDirPath, err)
                } else {
                    fmt.Printf("Renamed %s to %s\n", path, newDirPath)
                }
            }
        }
        return nil
    })

    if err != nil {
        fmt.Printf("Error walking the directory: %v\n", err)
    }
}
