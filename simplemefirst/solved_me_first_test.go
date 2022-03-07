package simplemefirst

import "testing"

func TestSolvedMeFirst(t *testing.T) {
	var a, b, res uint32
	a = 2
	b = 3
	res = solveMeFirst(a, b)
	if res != 5 {
		t.Error("Expected 5, got ", res)
	}
}
