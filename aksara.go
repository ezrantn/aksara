package aksara

import (
	"errors"
	"regexp"
	"strings"
)

type AksaraTranslator struct {
	latinToJavanese       map[string]string
	javaneseToLatin       map[string]string
	complexJavaneseRules  map[string]string
	complexConsonantRules map[string]string
	vowelDiacritics       map[string]string
	specialCharacters     map[string]string
}

func NewAksaraTranslator() *AksaraTranslator {
	return &AksaraTranslator{
		latinToJavanese: map[string]string{
			"a": "ꦲ", "b": "ꦧ", "c": "ꦕ", "d": "ꦢ",
			"e": "ꦺ", "f": "ꦥ꦳", "g": "ꦒ", "h": "ꦲꦃ",
			"i": "ꦶ", "j": "ꦗ", "k": "ꦏ", "l": "ꦭ",
			"m": "ꦩ", "n": "ꦤ", "o": "ꦺꦴ", "p": "ꦥ",
			"q": "ꦐ", "r": "ꦫ", "s": "ꦱ", "t": "ꦠ",
			"u": "ꦸ", "v": "ꦮ꦳", "w": "ꦮ", "x": "ꦼ",
			"y": "ꦪ", "z": "ꦗ꦳",
		},
		javaneseToLatin: map[string]string{
			"ꦲ": "a", "ꦧ": "b", "ꦕ": "c", "ꦢ": "d",
			"ꦺ": "e", "ꦥ꦳": "f", "ꦒ": "g", "ꦲꦃ": "h",
			"ꦶ": "i", "ꦗ": "j", "ꦏ": "k", "ꦭ": "l",
			"ꦩ": "m", "ꦤ": "n", "ꦺꦴ": "o", "ꦥ": "p",
			"ꦐ": "q", "ꦫ": "r", "ꦱ": "s", "ꦠ": "t",
			"ꦸ": "u", "ꦮ꦳": "v", "ꦮ": "w", "ꦼ": "x",
			"ꦪ": "y", "ꦗ꦳": "z",
		},
		complexConsonantRules: map[string]string{
			"ng": "ꦔ", "ny": "ꦚ",
			"th": "ꦛ", "dh": "ꦝ",
		},
		vowelDiacritics: map[string]string{
			"a": "ꦄ", "i": "ꦶ", "u": "ꦸ",
			"e": "ꦺ", "o": "ꦺꦴ",
		},
		specialCharacters: map[string]string{
			"(": "꧀", ")": "꧀",
			"[": "꧀", "]": "꧀",
			".": "꧁", ",": "꧂",
		},
		complexJavaneseRules: map[string]string{
			"ꦲꦶ": "i", "ꦲꦸ": "u", "ꦲꦺ": "e", "ꦲꦺꦴ": "o",

			"ꦔ": "ng", "ꦚ": "ny", "ꦛ": "th", "ꦝ": "dh",
		},
	}
}

// TranslateLatinToJavanese provides context-aware translation
func (t *AksaraTranslator) TranslateLatinToJavanese(input string) (string, error) {
	if input == "" {
		return "", errors.New("input cannot be empty")
	}

	// Normalize input
	input = strings.ToLower(input)

	// Preprocess complex consonants
	for consonant, replacement := range t.complexConsonantRules {
		input = strings.ReplaceAll(input, consonant, replacement)
	}

	var result strings.Builder
	runes := []rune(input)

	for i := 0; i < len(runes); i++ {
		char := string(runes[i])

		// Handle special characters
		if special, ok := t.specialCharacters[char]; ok {
			result.WriteString(special)
			continue
		}

		// Handle complex consonants and vowel combinations
		if i+1 < len(runes) {
			twoCharSeq := char + string(runes[i+1])
			if replacement, ok := t.complexConsonantRules[twoCharSeq]; ok {
				result.WriteString(replacement)
				i++ // Skip next character
				continue
			}
		}

		// Translate individual characters
		if javaneseChar, ok := t.latinToJavanese[char]; ok {
			result.WriteString(javaneseChar)
		} else if vowelDiacritic, ok := t.vowelDiacritics[char]; ok {
			result.WriteString(vowelDiacritic)
		} else {
			// Preserve characters not in translation map
			result.WriteString(char)
		}
	}

	return result.String(), nil
}

// TranslateJavaneseToLatin provides enhanced Latin translation
func (t *AksaraTranslator) TranslateJavaneseToLatin(input string) (string, error) {
	if input == "" {
		return "", errors.New("input cannot be empty")
	}
	
	if translation, ok := t.complexJavaneseRules[input]; ok {
		return translation, nil
	}

	var result strings.Builder

	// Iterate through characters with context awareness
	for i := 0; i < len(input); i++ {
		currentChar := string(input[i])

		// Check for complex character combinations first
		if i+1 < len(input) {
			twoCharSeq := currentChar + string(input[i+1])
			if translation, ok := t.complexJavaneseRules[twoCharSeq]; ok {
				result.WriteString(translation)
				i++ // Skip next character
				continue
			}
		}

		// Attempt direct translation
		if latinChar, ok := t.javaneseToLatin[currentChar]; ok {
			result.WriteString(latinChar)
		} else {
			// Preserve characters not in translation map
			result.WriteString(currentChar)
		}
	}

	return result.String(), nil
}

// IsValidAksaraJava checks if input is valid Javanese script
func (t *AksaraTranslator) IsValidAksaraJava(input string) bool {
	validationRegex := regexp.MustCompile(`^[ꦲ-ꧾ\s]+$`)
	return validationRegex.MatchString(input)
}

// NormalizeText provides text normalization for Javanese script
func (t *AksaraTranslator) NormalizeText(input string) string {
	input = strings.TrimSpace(input)

	if t.IsValidAksaraJava(input) {
		return input
	}

	return strings.ToLower(input)
}

// DetectScript determines whether input is Latin or Javanese script
func (t *AksaraTranslator) DetectScript(input string) string {
	for _, char := range input {
		if char >= 0xA980 && char <= 0xA9DF {
			return "Javanese"
		}
	}
	return "Latin"
}

// Translate provides context-aware bidirectional translation
func (t *AksaraTranslator) Translate(input string) (string, error) {
	script := t.DetectScript(input)

	if script == "Latin" {
		return t.TranslateLatinToJavanese(input)
	}

	return t.TranslateJavaneseToLatin(input)
}
