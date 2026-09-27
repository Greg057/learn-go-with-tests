package arrays

func Sum(numbers []int) (sum int) {
	for _, v := range numbers {
		sum += v
	}
	return
}

func SumAll(slicesToSum ...[]int) []int {
	var sums []int

	for _, values := range slicesToSum {
		sums = append(sums, Sum(values))
	}
	return sums
}

func SumAllTails(slicesToSum ...[]int) []int {
	var sums []int

	for _, values := range slicesToSum {
		if len(values) < 1 {
			sums = append(sums, 0)
		} else {
			tail := values[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
