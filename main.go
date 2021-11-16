package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	templateAppFolderName  = "template-app"
	noLimit                = -1
	serviceNameTemplateKey = "ServiceName"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatalf("No service name was provided. Please pass the name of the service for which you wish to generate code")
	}
	serviceName := os.Args[1]

	vars := map[string]string{serviceNameTemplateKey: serviceName}

	files, err := getFilesIn(templateAppFolderName)
	if err != nil {
		log.Fatalf("get files under %s: %s", templateAppFolderName, err)
	}

	for _, file := range files {
		templatizedFile, err := processFile(file, vars)
		if err != nil {
			log.Fatalf("process file: %s: %s", file, err)
		}

		newPath := strings.Replace(file, templateAppFolderName, serviceName, noLimit)
		f, err := createFile(newPath)
		if err != nil {
			log.Fatalf("create file: %s", err)
		}

		_, err = f.Write([]byte(templatizedFile))
		if err != nil {
			log.Fatalf("write to file: %s: %s", file, err)
		}
	}
}

func getFilesIn(rootPath string) ([]string, error) {
	files := make([]string, 0)
	err := filepath.Walk(rootPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return fmt.Errorf("walk: %w", err)
			}
			if !info.IsDir() {
				files = append(files, path)
			}
			return nil
		})

	return files, err
}

func createFile(path string) (*os.File, error) {
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0775)
		if err != nil {
			return nil, fmt.Errorf("make dir: %s: %w", dir, err)
		}
	}

	file, err := os.Create(path)
	if err != nil {
		return nil, fmt.Errorf("make file: %s: %w", path, err)
	}

	return file, nil
}

func processFile(fileName string, vars interface{}) (string, error) {
	tmpl, err := template.ParseFiles(fileName)
	if err != nil {
		return "", fmt.Errorf("parse file: %s: %w", fileName, err)
	}

	var tmplBytes bytes.Buffer
	err = tmpl.Execute(&tmplBytes, vars)
	if err != nil {
		return "", fmt.Errorf("tokenize file: %s: %w", fileName, err)
	}

	return tmplBytes.String(), nil
}
