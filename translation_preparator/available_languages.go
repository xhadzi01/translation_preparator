// available_languages brings stringer itf for available languages

package translation_preparator

import (
	"golang.org/x/text/language"
)

type LanguageTag language.Tag

func (lg LanguageTag) GoString() string {
	return lg.String()
}

func (lg LanguageTag) String() string {
	translated := language.Tag(lg)

	switch translated {
	case language.Afrikaans:
		return "Afrikaans"
	case language.Amharic:
		return "Amharic"
	case language.Arabic:
		return "Arabic"
	case language.ModernStandardArabic:
		return "ModernStandardArabic"
	case language.Azerbaijani:
		return "Azerbaijani"
	case language.Bulgarian:
		return "Bulgarian"
	case language.Bengali:
		return "Bengali"
	case language.Catalan:
		return "Catalan"
	case language.Czech:
		return "Czech"
	case language.Danish:
		return "Danish"
	case language.German:
		return "German"
	case language.Greek:
		return "Greek"
	case language.English:
		return "English"
	case language.AmericanEnglish:
		return "AmericanEnglish"
	case language.BritishEnglish:
		return "BritishEnglish"
	case language.Spanish:
		return "Spanish"
	case language.EuropeanSpanish:
		return "EuropeanSpanish"
	case language.LatinAmericanSpanish:
		return "LatinAmericanSpanish"
	case language.Estonian:
		return "Estonian"
	case language.Persian:
		return "Persian"
	case language.Finnish:
		return "Finnish"
	case language.Filipino:
		return "Filipino"
	case language.French:
		return "French"
	case language.CanadianFrench:
		return "CanadianFrench"
	case language.Gujarati:
		return "Gujarati"
	case language.Hebrew:
		return "Hebrew"
	case language.Hindi:
		return "Hindi"
	case language.Croatian:
		return "Croatian"
	case language.Hungarian:
		return "Hungarian"
	case language.Armenian:
		return "Armenian"
	case language.Indonesian:
		return "Indonesian"
	case language.Icelandic:
		return "Icelandic"
	case language.Italian:
		return "Italian"
	case language.Japanese:
		return "Japanese"
	case language.Georgian:
		return "Georgian"
	case language.Kazakh:
		return "Kazakh"
	case language.Khmer:
		return "Khmer"
	case language.Kannada:
		return "Kannada"
	case language.Korean:
		return "Korean"
	case language.Kirghiz:
		return "Kirghiz"
	case language.Lao:
		return "Lao"
	case language.Lithuanian:
		return "Lithuanian"
	case language.Latvian:
		return "Latvian"
	case language.Macedonian:
		return "Macedonian"
	case language.Malayalam:
		return "Malayalam"
	case language.Mongolian:
		return "Mongolian"
	case language.Marathi:
		return "Marathi"
	case language.Malay:
		return "Malay"
	case language.Burmese:
		return "Burmese"
	case language.Nepali:
		return "Nepali"
	case language.Dutch:
		return "Dutch"
	case language.Norwegian:
		return "Norwegian"
	case language.Punjabi:
		return "Punjabi"
	case language.Polish:
		return "Polish"
	case language.Portuguese:
		return "Portuguese"
	case language.BrazilianPortuguese:
		return "BrazilianPortuguese"
	case language.EuropeanPortuguese:
		return "EuropeanPortuguese"
	case language.Romanian:
		return "Romanian"
	case language.Russian:
		return "Russian"
	case language.Sinhala:
		return "Sinhala"
	case language.Slovak:
		return "Slovak"
	case language.Slovenian:
		return "Slovenian"
	case language.Albanian:
		return "Albanian"
	case language.Serbian:
		return "Serbian"
	case language.SerbianLatin:
		return "SerbianLatin"
	case language.Swedish:
		return "Swedish"
	case language.Swahili:
		return "Swahili"
	case language.Tamil:
		return "Tamil"
	case language.Telugu:
		return "Telugu"
	case language.Thai:
		return "Thai"
	case language.Turkish:
		return "Turkish"
	case language.Ukrainian:
		return "Ukrainian"
	case language.Urdu:
		return "Urdu"
	case language.Uzbek:
		return "Uzbek"
	case language.Vietnamese:
		return "Vietnamese"
	case language.Chinese:
		return "Chinese"
	case language.SimplifiedChinese:
		return "SimplifiedChinese"
	case language.TraditionalChinese:
		return "TraditionalChinese"
	case language.Zulu:
		return "Zulu"
	default:
		return "????Unknown????"
	}
}
