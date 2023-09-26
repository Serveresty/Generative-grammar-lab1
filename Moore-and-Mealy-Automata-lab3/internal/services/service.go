package services

import "strconv"

var nominals = []int{1, 2, 5, 10}

func Automata(coins []int, price int) [][]string {
	var results [][]string
	currentState := 0
	var payback string
	for _, elData := range coins {
		isNominal := findInputNominal(elData)
		if isNominal {
			currentState += elData
			if currentState < price {
				results = append(results, []string{strconv.Itoa(currentState), strconv.Itoa(elData) + "|-"})
			} else {
				if price >= currentState {
					payback = strconv.Itoa(price - currentState)
				} else {
					payback = strconv.Itoa(currentState - price)
				}
				currentState = 0
				results = append(results, []string{strconv.Itoa(currentState), strconv.Itoa(elData) + "|" + payback})
			}
		} else {
			payback = "Wrong nominal"
			results = append(results, []string{strconv.Itoa(currentState), strconv.Itoa(elData) + "|" + payback})
		}
	}
	return results
}

func findInputNominal(coin int) bool {
	for _, k := range nominals {
		if k == coin {
			return true
		}
	}
	return false
}
