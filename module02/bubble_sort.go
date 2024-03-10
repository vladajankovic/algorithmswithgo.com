package module02

// BubbleSortInt will sort a list of integers using the bubble sort algorithm.
func BubbleSortInt(list []int) {
	for i := 0; i < len(list); i++ {
		sorted := true
		for j := 1; j < len(list)-i; j++ {
			if list[j-1] > list[j] {
				list[j-1], list[j] = list[j], list[j-1]
				sorted = false
			}
		}
		if sorted {
			break
		}
	}
}
