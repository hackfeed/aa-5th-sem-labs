package levenshtein

import "fmt"

// ReadWord used to read a word through EOL symbol.
func ReadWord() string {
	var word string
	fmt.Scanln(&word)

	return word
}
