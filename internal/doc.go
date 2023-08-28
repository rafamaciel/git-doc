package internal

import (
	"embed"
	"io"
	"os"
	"path/filepath"
)

//go:embed templates/*
var templateFiles embed.FS

func GenerateSkeleton() error {
	entries, err := templateFiles.ReadDir("templates")
	if err != nil {
		return err
	}

	err = os.MkdirAll("docs", os.ModePerm)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		templateFile, err := templateFiles.Open(filepath.Join("templates", entry.Name()))
		if err != nil {
			return err
		}
		defer templateFile.Close()

		outputFilePath := filepath.Join("docs", entry.Name())
		outputFile, err := os.Create(outputFilePath)
		if err != nil {
			return err
		}
		defer outputFile.Close()

		_, err = io.Copy(outputFile, templateFile)
		if err != nil {
			return err
		}
	}

	return nil
}
