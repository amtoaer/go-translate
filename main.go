package main

import (
	"fmt"
	"os"

	"github.com/amtoaer/go-translate/translator"
)

func main() {
	var (
		content, fromLang, toLang string
		method                    translator.Method
	)
	if len(os.Args) >= 6 {
		content = os.Args[1]
		fromLang = os.Args[3]
		toLang = os.Args[5]
		if len(os.Args) >= 8 {
			switch os.Args[7] {
			case "baidu":
				method = translator.BAIDU
			case "google":
				method = translator.GOOGLE
			default:
				method = translator.NIU
			}
		} else {
			method = translator.NIU
		}
		translator := translator.NewTranslator(method)
		fmt.Println(translator.Translate(content, fromLang, toLang))
	}
}