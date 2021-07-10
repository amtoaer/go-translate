package main

import (
	"fmt"
	"os"

	"github.com/amtoaer/go-translate/config"
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
			case "niu":
				method = translator.NIU
			default:
				method = translator.YOUDAO
			}
		} else {
			method = translator.YOUDAO
		}
		if err := config.InitConfig(); err != nil {
			panic(err)
		}
		translator := translator.NewTranslator(method)
		fmt.Println(translator.Translate(content, fromLang, toLang))
	}
}
