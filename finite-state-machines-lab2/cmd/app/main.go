package main

import (
	"finite-state-machines-lab2/internal/services"
	"fmt"
)

func main() {
	var data = []string{"", "0", "1", "00", "01", "10", "11", "0001", "1000", "0110", "1100", "0011", "001100", "110011"}
	matrix := services.GenerateMachine()
	result := services.CheckWords(matrix, data)

	for _, el := range result {
		fmt.Println("--------------")
		fmt.Println("Word: " + el.Word)
		fmt.Println("Affiliation: ", el.Affiliation)
		fmt.Println("State: " + el.State)
	}
}
