package sum

func Sum(input []int) int {
	sum := 0

	for _, number := range input {
		sum += number
	}

	return sum
}

func SumAll(input ...[]int) []int {
	var sums []int

	for _, num := range input {
		sums = append(sums, Sum(num))
	}

	return sums
}

func SumAllTails(input ...[]int) []int {
	var sums []int

	for _, num := range input {
		if len(num) == 0 {
			sums = append(sums, 0)
		} else {
			tail := num[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
