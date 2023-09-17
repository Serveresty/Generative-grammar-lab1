package main

import "GenerativeGrammar/models"

func main() {
	rules := []models.Rules{{NotTerminal: "S", Terminal: "AA"}, {NotTerminal: "A", Terminal: "aAb"}, {NotTerminal: "A", Terminal: "ab"}}
}
