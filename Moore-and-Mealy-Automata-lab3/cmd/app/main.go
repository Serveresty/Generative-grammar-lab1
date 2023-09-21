package main

import (
	"Moore-and-Mealy-Automata-lab3/internal/services"
	"fmt"
)

func main() {
	var coins = []int{5, 10, 8}
	var price int = 2
	results := services.Automata(coins, price)
	fmt.Println(results)
}
