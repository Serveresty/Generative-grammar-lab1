package main

import (
	"Creating-a-lexical-analyzer-lab5/internal/service"
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
		service.CheckStr(scanner.Text())
	}
}
