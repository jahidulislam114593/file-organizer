package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

var fileTypes = map[string]string{
    ".mp3": "Music",
    ".wav": "Music",
    ".mp4": "Videos",
    ".mkv": "Videos",
    ".jpg": "Images",
    ".png": "Images",
    ".pdf": "Documents",
    ".docx": "Documents",
    ".zip": "Archives",
    ".rar": "Archives",
    ".go": "Code",
    ".js": "Code",
    ".py": "Code",
}

func main() {
    var inputDir string
    fmt.Print("Enter the directory to organize: ")
    fmt.Scanln(&inputDir)

    err := filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        if !info.IsDir() {
            ext := strings.ToLower(filepath.Ext(info.Name()))
            category, ok := fileTypes[ext]
            if ok {
                targetDir := filepath.Join(inputDir, category)
                os.MkdirAll(targetDir, os.ModePerm)
                newPath := filepath.Join(targetDir, info.Name())
                os.Rename(path, newPath)
            }
        }
        return nil
    })

    if err != nil {
        fmt.Println("Error organizing files:", err)
    } else {
        fmt.Println("Files organized successfully.")
    }
}
