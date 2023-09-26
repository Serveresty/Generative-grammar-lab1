package services

import (
	"fmt"
	"strings"
)

var arr = [][2]string{{"S", "var"}, {"S", "int"}, {"S", "float"}, {"S", "S+S"}, {"S", "S/S"}, {"S", "S*S"}, {"S", "S-S"}, {"S", "(S)"}}

func CheckStr(str string) {
	for str != "S" {
		for _, elem := range arr {
			str = strings.Replace(str, elem[1], elem[0], 1)
			fmt.Println(str)
		}
	}
}
