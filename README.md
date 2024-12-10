# Aksara

Aksara is a comprehensive Go package designed to facilitate bidirectional translation between Latin and Javanese (Aksara Jawa) scripts. This library provides robust, context-aware translation capabilities for Javanese language characters.

## Features

- Bidirectional translation between Latin and Javanese scripts
- Context-aware translation
- Support for complex character combinations
- Script detection
- Text normalization
- Comprehensive character mapping

## Limitations

- Supports standard Latin and Javanese script characters
- May not cover all dialectical variations (in progress to improve this!)
- Requires well-formed input

## Installation

To install the library, use Go modules:

```bash
go get github.com/ezrantn/aksara
```

## Usage

### Basic Translation

```go
package main

import (
    "fmt"
    "github.com/ezrantn/aksara"
)

func main() {
    translator := aksara.NewAksaraTranslator()

    // Latin to Javanese Translation
    javaneseText, _ := translator.TranslateLatinToJavanese("hello")
    fmt.Println(javaneseText)

    // Javanese to Latin Translation
    latinText, _ := translator.TranslateJavaneseToLatin("ꦲꦺꦭꦺꦴ")
    fmt.Println(latinText)
}
```

### Other Methods

`Translate`

Automatically detects and translates script:

```go
func ExampleTranslate() {
    translator := aksara.NewAksaraTranslator()
    
    // Automatic Latin to Javanese
    result, _ := translator.Translate("hello")
    fmt.Println(result) // Output: ꦲꦺꦭꦺꦴ
    
    // Automatic Javanese to Latin
    result, _ = translator.Translate("ꦲꦺꦭꦺꦴ")
    fmt.Println(result) // Output: hello
}
```

`IsValidAksaraJava`

Validates Javanese script:

```go
func ExampleIsValidAksaraJava() {
    translator := aksara.NewAksaraTranslator()
    
    // Valid Javanese script
    valid := translator.IsValidAksaraJava("ꦲꦺꦭꦺꦴ")
    fmt.Println(valid) // Output: true
    
    // Invalid script
    valid = translator.IsValidAksaraJava("hello")
    fmt.Println(valid) // Output: false
}
```

`DetectScript`

Identifies the script type:

```go
func ExampleDetectScript() {
    translator := aksara.NewAksaraTranslator()
    
    // Detect Latin script
    script := translator.DetectScript("hello")
    fmt.Println(script) // Output: Latin
    
    // Detect Javanese script
    script = translator.DetectScript("ꦲꦺꦭꦺꦴ")
    fmt.Println(script) // Output: Javanese
}
```

`NormalizeText`

Normalizes input text with different behaviors for Latin and Javanese scripts:

```go
func ExampleNormalizeText() {
    translator := aksara.NewAksaraTranslator()
    
    // Normalize Latin text (converts to lowercase)
    normalized := translator.NormalizeText("HeLLo WoRLd")
    fmt.Println(normalized) // Output: hello world
    
    // Preserve Javanese script case
    javaneseText := "ꦲꦺꦭꦺꦴ ꦮꦺꦴꦫꦭꦢ"
    normalized = translator.NormalizeText(javaneseText)
    fmt.Println(normalized) // Output: ꦲꦺꦭꦺꦴ ꦮꦺꦴꦫꦭꦢ
    
    // Remove extra whitespaces
    spacedText := "  hello   world  "
    normalized = translator.NormalizeText(spacedText)
    fmt.Println(normalized) // Output: hello world
}
```

## License

This tool is open-source and available under the [MIT](https://github.com/ezrantn/aksara/blob/main/LICENSE) License.

## Contributions

Contributions are welcome! Please feel free to submit a pull request.
