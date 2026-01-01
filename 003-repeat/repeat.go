package repeat

func Repeat(input string) (output string) {
	for i := 0; i < 5; i++ {
		output += input
	}
	return output
}
