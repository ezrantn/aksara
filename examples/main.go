package main

import (
	"fmt"

	"github.com/ezrantn/aksara"
)

func main() {
	translator := aksara.NewAksaraTranslator()

	javaneseText, err := translator.TranslateLatinToJavanese("kanca")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(javaneseText)

	// Translate from Javanese to Latin
	kanca := translator.NormalizeText("ꦏꦚ꧀ꦕ")
	latinText, err := translator.TranslateJavaneseToLatin(kanca)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(latinText)

	// Automatic bidirectional translation
	translatedText, err := translator.Translate("kanca")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(translatedText)
}
