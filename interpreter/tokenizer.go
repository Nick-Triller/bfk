package interpreter

var validTokens = map[rune]bool{
	'<': true,
	'>': true,
	'+': true,
	'-': true,
	'.': true,
	',': true,
	'[': true,
	']': true,
}

// tokenize converts a string to a rune slice containing only valid tokens.
func Tokenize(in string) []rune {
	tokens := make([]rune, 0, len(in))
	for _, ch := range in {
		if _, ok := validTokens[ch]; ok {
			tokens = append(tokens, ch)
		}
	}
	return tokens
}
