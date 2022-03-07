package appleorange

import "testing"

func TestAppleOrange(t *testing.T) {
	// input
	// 7 11
	// 5 15
	// 3 2
	// -2 2 1
	// 5 -6

	// output
	// 1
	// 1

	countApplesAndOranges(7, 11, 5, 15, []int32{-2, 2, 1}, []int32{5, -6})

}
