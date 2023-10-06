package services

import (
	"fmt"
	"strings"
)

var arr = [][2]string{{"S", "var"}, {"S", "int"}, {"S", "float"}, {"S", "S+S"}, {"S", "S/S"}, {"S", "S*S"}, {"S", "S-S"}, {"S", "(S)"}}

func CheckStr(str string) {
	checker := ""
	fmt.Println(str)
	for str != "S" {
		for _, elem := range arr {
			str = strings.Replace(str, elem[1], elem[0], 1)
			if checker != str {
				fmt.Println(str)
			}
			checker = str
		}
	}
}
