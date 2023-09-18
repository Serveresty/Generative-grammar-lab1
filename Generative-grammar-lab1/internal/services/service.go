package services

import (
	"Generative-grammar-lab1/internal/models"
	"math/rand"
	"strings"
)

func Generate(startWord string, rules map[string][]models.Rules) string {
	if startWord != strings.ToLower(startWord) {
		for startWord != strings.ToLower(startWord) {
			for key, elem := range rules {
				if len(elem) > 1 {
					rndElem := rand.Intn(len(elem) - 0)
					startWord = strings.Replace(startWord, key, elem[rndElem].Terminal, -1)
				} else {
					startWord = strings.Replace(startWord, key, elem[0].Terminal, -1)
				}
			}
		}
		return startWord
	} else {
		return startWord
	}
}

func GetCountOfTeminators(rules []models.Rules) map[string][]models.Rules {
	var ruleCount = make(map[string][]models.Rules, 0)
	for _, elem := range rules {
		ruleCount[elem.NotTerminal] = append(ruleCount[elem.NotTerminal], elem)
	}
	return ruleCount
}
