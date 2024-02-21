package module01

func GCD(a, b int) int {
	for {
		if a%b == 0 {
			return b
		}
		a, b = b, a%b
	}
}
