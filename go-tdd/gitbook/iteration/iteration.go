package iteration

func Repeat(str string, length int) string {
	if length <= 0 {
		return ""
	}

	var repeated string
	for i := 0; i < length; i++ {
		repeated += str
	}
	return repeated
}
