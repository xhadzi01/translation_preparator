// translation_api_wrapper wraps native Cloud Translation Api for simplified use

package translation_preparator

type Translator interface {
	Translate(word, from, to string) (translated string, err error)
	TranslateMultiple(words []string, from, to string) (translatedWords []string, err error)
	SupportedTranslationLanguages(lang string) (languages []string, err error)
	Close() (err error)
}
