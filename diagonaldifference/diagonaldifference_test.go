package diagonaldifference

import "testing"

func TestDiagonalDifference(t *testing.T) {

	// Input
	// 3
	// 11 2 4
	// 4 5 6
	// 10 8 -12

	//Expected Output
	// 15

	var arr [][]int32 // arr := [][]int32{}

	arr = append(arr, []int32{11, 2, 4})   // Create a new row with the values 11, 2, 4 and append it to the array
	arr = append(arr, []int32{4, 5, 6})    // Create a new row with the values 4, 5, 6 and append it to the array
	arr = append(arr, []int32{10, 8, -12}) // Create a new row with the values 10, 8, -12 and append it to the array

	result := diagonalDifference(arr) // Call the function with the input array and store the result in the variable result

	if result != 15 {
		t.Errorf("Expected 15, got %d instead.", result)
	}

}

func TestDiagonalDifference1(t *testing.T) {
	// Sample Input
	// 	1 2 3
	// 4 5 6
	// 9 8 9

	// Expected Output
	// 2

	var arr [][]int32 // arr := [][]int32{}

	arr = append(arr, []int32{1, 2, 3}) // Create a new row with the values 1, 2, 3 and append it to the array
	arr = append(arr, []int32{4, 5, 6}) // Create a new row with the values 4, 5, 6 and append it to the array
	arr = append(arr, []int32{9, 8, 9}) // Create a new row with the values 9, 8, 9 and append it to the array

	result := diagonalDifference(arr) // Call the function with the input array and store the result in the variable result

	if result != 2 {
		t.Errorf("Expected 2, got %d instead.", result)
	}
}
