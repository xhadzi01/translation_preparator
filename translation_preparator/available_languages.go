// available_languages brings stringer itf for available languages

package translation_preparator

import (
	"fmt"
	"strings"

	"golang.org/x/text/language"
)

type LanguageTag language.Tag

var languageMap map[language.Tag]string

func init() {
	languageMap = map[language.Tag]string{
		language.Afrikaans:            "Afrikaans",
		language.Amharic:              "Amharic",
		language.Arabic:               "Arabic",
		language.ModernStandardArabic: "ModernStandardArabic",
		language.Azerbaijani:          "Azerbaijani",
		language.Bulgarian:            "Bulgarian",
		language.Bengali:              "Bengali",
		language.Catalan:              "Catalan",
		language.Czech:                "Czech",
		language.Danish:               "Danish",
		language.German:               "German",
		language.Greek:                "Greek",
		language.English:              "English",
		language.AmericanEnglish:      "AmericanEnglish",
		language.BritishEnglish:       "BritishEnglish",
		language.Spanish:              "Spanish",
		language.EuropeanSpanish:      "EuropeanSpanish",
		language.LatinAmericanSpanish: "LatinAmericanSpanish",
		language.Estonian:             "Estonian",
		language.Persian:              "Persian",
		language.Finnish:              "Finnish",
		language.Filipino:             "Filipino",
		language.French:               "French",
		language.CanadianFrench:       "CanadianFrench",
		language.Gujarati:             "Gujarati",
		language.Hebrew:               "Hebrew",
		language.Hindi:                "Hindi",
		language.Croatian:             "Croatian",
		language.Hungarian:            "Hungarian",
		language.Armenian:             "Armenian",
		language.Indonesian:           "Indonesian",
		language.Icelandic:            "Icelandic",
		language.Italian:              "Italian",
		language.Japanese:             "Japanese",
		language.Georgian:             "Georgian",
		language.Kazakh:               "Kazakh",
		language.Khmer:                "Khmer",
		language.Kannada:              "Kannada",
		language.Korean:               "Korean",
		language.Kirghiz:              "Kirghiz",
		language.Lao:                  "Lao",
		language.Lithuanian:           "Lithuanian",
		language.Latvian:              "Latvian",
		language.Macedonian:           "Macedonian",
		language.Malayalam:            "Malayalam",
		language.Mongolian:            "Mongolian",
		language.Marathi:              "Marathi",
		language.Malay:                "Malay",
		language.Burmese:              "Burmese",
		language.Nepali:               "Nepali",
		language.Dutch:                "Dutch",
		language.Norwegian:            "Norwegian",
		language.Punjabi:              "Punjabi",
		language.Polish:               "Polish",
		language.Portuguese:           "Portuguese",
		language.BrazilianPortuguese:  "BrazilianPortuguese",
		language.EuropeanPortuguese:   "EuropeanPortuguese",
		language.Romanian:             "Romanian",
		language.Russian:              "Russian",
		language.Sinhala:              "Sinhala",
		language.Slovak:               "Slovak",
		language.Slovenian:            "Slovenian",
		language.Albanian:             "Albanian",
		language.Serbian:              "Serbian",
		language.SerbianLatin:         "SerbianLatin",
		language.Swedish:              "Swedish",
		language.Swahili:              "Swahili",
		language.Tamil:                "Tamil",
		language.Telugu:               "Telugu",
		language.Thai:                 "Thai",
		language.Turkish:              "Turkish",
		language.Ukrainian:            "Ukrainian",
		language.Urdu:                 "Urdu",
		language.Uzbek:                "Uzbek",
		language.Vietnamese:           "Vietnamese",
		language.Chinese:              "Chinese",
		language.SimplifiedChinese:    "SimplifiedChinese",
		language.TraditionalChinese:   "TraditionalChinese",
		language.Zulu:                 "Zulu",
	}
}

func (lg LanguageTag) GoString() string {
	return lg.String()
}

func (lg LanguageTag) String() string {
	translated := language.Tag(lg)

	if value, ok := languageMap[translated]; ok {
		return value
	}

	return "????Unknown????"
}

func ToLanguageTag(lang string) (retval LanguageTag, err error) {

	for tagValue, nameValue := range languageMap {
		if strings.EqualFold(nameValue, lang) {
			retval = LanguageTag(tagValue)
			return
		}
	}
	err = fmt.Errorf("could not translate language tag from string %s", lang)
	return
}
