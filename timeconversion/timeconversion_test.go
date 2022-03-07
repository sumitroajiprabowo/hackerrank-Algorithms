package main

import "testing"

func TestTimeConversion(t *testing.T) {

	// Sample input 07:05:45PM
	// Output 19:05:45

	// Write your code here

	var s string = "07:05:45PM"
	var expected string = "19:05:45"

	actual := timeConversion(s)

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
