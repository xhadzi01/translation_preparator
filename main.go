package main

import (
	"fmt"
	"os"
	"path/filepath"
	"translation_preparator/translation_preparator"
)

func main() {

	// get location of current executable
	execPath, err := os.Executable()
	if err != nil {
		panic(err)
	}

	// json auth key should be in the same folder
	execDirpath := filepath.Dir(execPath)
	jsonAuthFilePath := filepath.Join(execDirpath, "translation-preparator-auth.json")

	translator, translatorErr := translation_preparator.NewTranslator(jsonAuthFilePath)
	if translatorErr != nil {
		fmt.Println(translatorErr.Error())
		return
	}
	defer translator.Close()

	supportedLang, supportedLangErr := translator.SupportedTranslationLanguages("English")
	if supportedLangErr != nil {
		fmt.Println(supportedLangErr.Error())
		return
	}

	for _, lang := range supportedLang {
		fmt.Println(lang)
	}

	translator.Translate("lietadlo", "Slovak", "AmericanEnglish")
}
