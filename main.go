package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func icon() {
	fmt.Println(`
 /$$$$$$$   /$$$$$$ 
| $$__  $$ /$$__  $$
| $$  \ $$| $$  \__/
| $$$$$$$/| $$      
| $$__  $$| $$      
| $$  \ $$| $$    $$
| $$  | $$|  $$$$$$/
|__/  |__/ \______/ 	React Component Generator, v1.0.0
						
	`)
}

func main() {
	err := loadConfig()
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			icon()
			log.Println("config file not found")
			generateConfig()
			main()
			return
		}

		log.Fatal(err)
	}

	if len(os.Args) < 2 {
		log.Fatal("file path is required")
	}

	path := os.Args[1]
	directory, fileName := filepath.Split(path)

	err = os.MkdirAll(directory, 0755)
	if err != nil {
		log.Fatalf("error creating directory: %v", err)
	}

	componentName := strings.Title(strings.Replace(fileName, "/", "", -1))

	filePath := filepath.Join(directory, componentName+".tsx")

	if _, err := os.Stat(filePath); err == nil {
		log.Fatalf("file %s already exists", filePath)
	}

	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("error creating file: %v", err)
	}
	defer file.Close()

	data := functionExport(componentName)
	_, err = file.WriteString(data)
	if err != nil {
		log.Fatalf("error writing to file: %v", err)
	}

	log.Printf("Component %s created successfully at %s\n", componentName, filePath)

	if config.Style.Enable {
		styleFilePath := filepath.Join(directory, componentName+"."+config.Style.Ext)
		if _, err := os.Stat(styleFilePath); err == nil {
			log.Fatalf("style file %s already exists", styleFilePath)
		}

		styleFile, err := os.Create(styleFilePath)
		if err != nil {
			log.Fatalf("error creating style file: %v", err)
		}
		defer styleFile.Close()

		_, err = styleFile.WriteString("// Styles for " + componentName + "\n")
		if err != nil {
			log.Fatalf("error writing to style file: %v", err)
		}

		log.Printf("Style file %s created successfully\n", styleFilePath)
	}
}
