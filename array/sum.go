package array

// Sum adds numbers of an input slice
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

// SumAll adds each input slice's numbers
func SumAll(slicesToSum ...[]int) []int {
	var sums []int

	for _, numbers := range slicesToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

//SumAllTails calculates the totals of the "tails" of each slice
func SumAllTails(slicesToSum ...[]int) []int {
	var sums, tail []int

	for _, numbers := range slicesToSum {
		if len(numbers) > 0 {
			tail = numbers[1:]
		} else {
			tail = numbers
		}

		sums = append(sums, Sum(tail))
	}

	return sums
}

// func main() {
// 	fmt.Println(Sum([]int{1, 2, 3, 4, 5}))
// 	fmt.Println(SumAll([]int{1, 2, 3, 4, 5}, []int{5, 9}))
// 	fmt.Println(SumAllTails([]int{1, 2, 3, 4, 5}, []int{5, 9}))
// }
