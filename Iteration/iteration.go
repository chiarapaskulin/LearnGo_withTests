package iteration

//Lesson 1
func Repeat1(character string) string {
	var repeated string

	for i := 0; i < 5; i++ {
		repeated += character
	}

	return repeated
}

func Repeat2(character string, maxIt int) string {
	var repeated string

	for i := 0; i < maxIt; i++ {
		repeated += character
	}

	return repeated
}
