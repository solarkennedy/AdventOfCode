package main

func maxInt(numbers ...int) int {
	max := numbers[0]
	for _, n := range numbers {
		if n > max {
			max = n
		}
	}
	return max
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
