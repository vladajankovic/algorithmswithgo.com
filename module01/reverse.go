package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	out := []rune(word)
	low := 0
	high := len(out) - 1

	for low < high {
		tmp := out[low]
		out[low] = out[high]
		out[high] = tmp
		low++
		high--
	}

	return string(out)

}
