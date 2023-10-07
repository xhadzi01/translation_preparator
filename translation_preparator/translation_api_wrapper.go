// translation_api_wrapper wraps native Cloud Translation Api for simplified use

package translation_preparator

import (
	"context"
	"errors"
	"fmt"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

type translator struct {
	context context.Context
	client  *translate.Client
}

// NewTranslator function creates Translator abstraction
func NewTranslator() (tr Translator, err error) {
	// create translation api client
	ctx := context.Background()
	client, err := translate.NewClient(ctx)
	if err != nil {
		return
	}

	// create translator obj
	tr = &translator{
		context: ctx,
		client:  client,
	}

	return
}

func (tr *translator) TranslateMultiple(words []string, from, to string) (translatedWords []string, err error) {
	if tr == nil {
		err = fmt.Errorf("translate object reference is invalid while tryin to translate word(s)")
		return
	}

	fromCasted, err := language.Parse(from)
	if err != nil {
		return
	}

	toCasted, err := language.Parse(to)
	if err != nil {
		return
	}

	translations, err := tr.client.Translate(tr.context,
		words, fromCasted,
		&translate.Options{
			Source: toCasted,
			Format: translate.Text,
		})
	if err != nil {
		return
	}

	if len(translations) != len(words) {
		err = fmt.Errorf("invalid number of retrieved translations, should be %d but is %d", len(words), len(translations))
	}

	translatedWords = make([]string, len(words))
	for idx, retrievedLang := range translations {
		translatedWords[idx] = retrievedLang.Text
	}

	return
}

func (tr *translator) Translate(word, from, to string) (translated string, err error) {
	wordsToTranslate := []string{
		word,
	}

	translatedWords, err := tr.TranslateMultiple(wordsToTranslate, from, to)
	if len(translatedWords) != 1 {
		err = fmt.Errorf("invalid number of retrieved translations, should be 1 but is %d", len(translatedWords))
	}
	translated = translatedWords[0]
	return
}

func (tr *translator) SupportedTranslationLanguages(lang string) (languages []string, err error) {
	if tr == nil {
		err = errors.New("translate object reference is invalid while tryin to obtain supported languages")
		return
	}

	casted, err := language.Parse(lang)
	if err != nil {
		return
	}

	retrievedLangs, err := tr.client.SupportedLanguages(tr.context, casted)
	if err != nil {
		return
	}

	languages = make([]string, len(retrievedLangs))

	for idx, retrievedLang := range retrievedLangs {
		languages[idx] = retrievedLang.Name
	}

	return
}

func (tr *translator) Close() (err error) {
	if tr == nil {
		err = errors.New("translate object reference is invalid while tryin to dispose of resources")
		return
	}

	if tr.client != nil {
		err = tr.client.Close()
	}

	return
}
