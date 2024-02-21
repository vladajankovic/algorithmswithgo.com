package module01

import (
	"math"
)

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//	BaseToDec("E", 16) => 14
//	BaseToDec("1110", 2) => 14
func BaseToDec(value string, base int) int {
	charset := make(map[byte]int)
	const table = "0123456789ABCDEF"
	for i := 0; i < len(table); i++ {
		charset[table[i]] = i
	}
	res := 0
	for i := 0; i < len(value); i++ {
		res += charset[value[i]] * int(math.Pow(float64(base), float64(len(value)-1-i)))
	}
	return res
}
