package app

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var s []string

// GetDataFromFile getting slice of string from file
func GetDataFromFile(dataFile *string) ([]string, error) {
	file, err := os.Open(*dataFile)
	if err != nil {
		return s, fmt.Errorf("ошибка открытия файла: %w", err)
	}
	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			log.Fatal(fmt.Errorf("ошибка закрытия файла: %w", err))
		}
	}(file)
	str := bufio.NewScanner(file)
	for str.Scan() {
		s = append(s, str.Text())
	}

	if err := str.Err(); err != nil {
		return s, fmt.Errorf("ошибка чтения буфера: %w", err)
	}
	return s, nil
}
