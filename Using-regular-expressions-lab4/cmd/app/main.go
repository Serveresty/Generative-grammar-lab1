package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {
	content, err := ioutil.ReadFile("../../web_data/data.txt")
	if err != nil {
		fmt.Println("Err")
	}

	re := regexp.MustCompile(`(?:(?:8|\+7)[\- ]?)?(?:\(?\d{3}\)?[\- ]?)?(?:\d{3}[\-]\d{2}[\-]\d{2})`)
	regxp := re.FindAllString(string(content), -1)
	fmt.Println("Count of numbers: ", len(regxp))
	for _, el := range regxp {
		fmt.Println(el)
	}
}
