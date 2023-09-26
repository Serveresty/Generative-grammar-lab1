package main

import (
	"Create-syntax-analyzer-lab6/internal/services"
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("../../internal/data/data.cpp")
	if err != nil {
		log.Fatal()
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		services.CheckStr(scanner.Text())
	}
}
