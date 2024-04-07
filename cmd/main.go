package main

import (
	"fmt"
	"log"

	"github.com/KozlovNikolai/gb-itog-kontr-osn-blok/internal/app"
)

var dataFile = "./initial-array"

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	fmt.Println("Start")

	s, err := app.GetData(&dataFile)
	if err != nil {
		return fmt.Errorf("run() %w", err)
	}
	fmt.Println(s)
	return nil
}
