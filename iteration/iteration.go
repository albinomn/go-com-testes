package iteration

const timesToRepeat = 5

func Repeat(char string) string {
	var repeated string
	for i := 0; i < timesToRepeat; i++ {
		repeated += char
	}

	return repeated
}
