package plusminus

import (
	"testing"
)

func Test(t *testing.T) {
	// 1 2 3 -1 -2 -3 0 0

	// Output
	// 	0.375000
	// 0.375000
	// 0.250000

	var arr []int32 // arr := [][]int32{}

	// // 1 2 3 -1 -2 -3 0 0
	arr = []int32{1, 2, 3, -1, -2, -3, 0, 0} // Create a new row with the values 1, 2, 3 and append it to the array

	plusMinus(arr) // Call the function with the input array and store the result in the variable result

}
