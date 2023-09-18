package main

import (
	"GenerativeGrammar/internal/models"
	"GenerativeGrammar/internal/services"
	"fmt"
)

func main() {
	rules := []models.Rules{{NotTerminal: "S", Terminal: "AA"}, {NotTerminal: "A", Terminal: "aAb"}, {NotTerminal: "A", Terminal: "ab"}}
	fmt.Println("********* 1 вариант *********\n---------  Правила  ---------")
	for _, el := range rules {
		fmt.Print(el.NotTerminal + "->" + el.Terminal + "\n")
	}
	startWord := "S"
	mapOfRules := services.GetCountOfTeminators(rules)
	word := services.Generate(startWord, mapOfRules)
	fmt.Println("Сгенерированное слово: " + word)
}
