package internal

func InText(key, text string) bool {
	matchCount := 0
	keyIndex := 0
	textIndex := 0
	for keyIndex < len(key) && keyIndex < len(text) && textIndex < len(text) {
		keyRune := key[keyIndex]
		textRune := text[textIndex]
		if keyRune == textRune {
			matchCount++
			keyIndex++
		}
		textIndex++
	}

	return len(key) == matchCount
}
