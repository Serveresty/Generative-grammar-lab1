package service

import (
	"fmt"
	"regexp"
	"strings"
)

func CheckStr(str string) {
	var mp = [][]string{{"comment", `//[^\n\r]*`}, {"preprocess_directive", `#.*`}, {"keywords", `\b(?:and|auto|bool|break|case|catch|char|class|
		const|continue|default|
		delete|do|double|else|
		float|for|friend|if|int
		|long|new|private|public
		|return|static|switch|struct|
		try|this|void|while|cout|endl|string|true|false)\b`}, {"separator", `[;(){}]`}, {"stringg", `[\"'][^\"']*['\"]`},
		{"operator", `[+|-|=|<<|>>|<|>|==]`}, {"intnum", `-?\\d+`}, {"floatnum", `-?\\d+\\.\\d+`}, {"identifiers", `[a-zA-Z_][A-Za-z_0-9]*`}}

	comment := `//[^\n\r]*`
	preprocess_directive := `#.*`
	keywords := `\b(?:and|auto|bool|break|case|catch|char|class|
		const|continue|default|
		delete|do|double|else|
		float|for|friend|if|int
		|long|new|private|public
		|return|static|switch|struct|
		try|this|void|while|cout|endl|string|true|false)\b`
	separator := `[;(){}]`
	identifiers := `[a-zA-Z_][A-Za-z_0-9]*`
	operator := `[+|-|=|<<|>>|<|>|==]`
	intnum := `-?\\d+`
	floatnum := `-?\\d+\\.\\d+`
	stringg := `[\"'][^\"']*['\"]`

	rs := strings.Join([]string{comment, preprocess_directive, keywords, separator, stringg, operator, intnum, floatnum, identifiers}, "|")
	rg := regexp.MustCompile(rs)
	resultt := rg.FindAllString(str, -1)
	for _, k := range resultt {
		for _, el := range mp {
			rgg := regexp.MustCompile(el[1])
			sttt := rgg.FindString(k)
			if sttt != "" {
				fmt.Println(el[0] + " " + k)
				break
			}
		}
	}
}
