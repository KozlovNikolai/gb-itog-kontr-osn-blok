package main

import (
	"fmt"
	"log"

	"github.com/KozlovNikolai/gb-itog-kontr-osn-blok/internal/app"
)

var dataFile = "./listing"

// readStringsToArray reads string to array
func readStringsToArray() ([]string, error) {
	s, err := app.GetDataFromFile(&dataFile)
	if err != nil {
		return nil, fmt.Errorf("ошибка получения данных из файла: %w", err)
	}
	return s, nil
}

// arrayTo3CharStringsArray
func arrayTo3CharStringsArray(a []string) []string {
	var r []string
	for _, v := range a {
		if len([]rune(v)) < 4 {
			r = append(r, v)
		}
	}
	return r
}

func main() {
	fmt.Println("Start")

	a, err := readStringsToArray()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Исходный массив строк: \n%q\n", a)
	s3 := arrayTo3CharStringsArray(a)
	fmt.Printf("\nМассив строк с длинной не больше 3ех символов: \n%q\n", s3)
}
