package grading

import (
	"fmt"
	"testing"
)

func TestGrading(t *testing.T) {

	// Write your code here
	// Test case 1
	grades := []int32{2, 37, 38}
	result := gradingStudents(grades)

	fmt.Println(result)
}
