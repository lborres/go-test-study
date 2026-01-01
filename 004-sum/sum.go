package sum

func Sum(input [5]int) int {
	sum := 0
	for _, number := range input {
		sum += number
	}

	return sum
}
