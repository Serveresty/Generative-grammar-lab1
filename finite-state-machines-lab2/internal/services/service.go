package services

import "finite-state-machines-lab2/internal/models"

func GenerateMachine() map[string]map[string]string {
	mp := map[string]map[string]string{
		"1": {"0": "2", "1": ""},
		"2": {"0": "3", "1": ""},
		"3": {"0": "3", "1": "3"},
	}
	return mp
}

func CheckWords(matrix map[string]map[string]string, data []string) []models.Result {

	results := []models.Result{}

	for _, elem := range data {
		var currentState string = "1"
		for idx, char := range elem {
			currentState = matrix[currentState][string(char)]
			if idx < 2 && currentState == "" || len(elem) < 2 {
				results = append(results, models.Result{Word: elem, Affiliation: false, State: ""})
				break
			}
			if idx == len(elem)-1 {
				results = append(results, models.Result{Word: elem, Affiliation: true, State: currentState})
			}
		}
	}
	return results
}
