package module02

// InsertionStortInt will sort a list of integers using the insertion sort
// algorithm.
func InsertionStortInt(list []int) {
	for i := 1; i < len(list); i++ {
		for j := i; j > 0 && list[j] < list[j-1]; j-- {
			list[j-1], list[j] = list[j], list[j-1]
		}
	}
}
