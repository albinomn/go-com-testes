package iteration

// Repeat the input char the number os times you pass in the second argument.
func Repeat(char string, timesToRepeat int) string {
	var repeated string
	for i := 0; i < timesToRepeat; i++ {
		repeated += char
	}

	return repeated
}
