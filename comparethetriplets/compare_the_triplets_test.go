package comparethetriplets

import "testing"

func TestCompareTheTriplets(t *testing.T) {

	var a, b []int32
	a = []int32{1, 2, 3}
	b = []int32{3, 2, 1}
	res := compareTriplets(a, b)
	if res[0] != 1 || res[1] != 1 {
		t.Error("Expected 1, 1, got ", res)
	}
}

func TestCompareTheTriplets2(t *testing.T) {

	var a, b []int32
	a = []int32{5, 6, 7}
	b = []int32{3, 6, 10}
	res := compareTriplets(a, b)
	if res[0] != 1 || res[1] != 1 {
		t.Error("Expected 1, 1, got ", res)
	}
}

func TestCompareTheTriplets3(t *testing.T) {

	var a, b []int32
	a = []int32{17, 28, 30}
	b = []int32{99, 16, 8}
	res := compareTriplets(a, b)
	if res[0] != 2 || res[1] != 1 {
		t.Error("Expected 2, 1, got ", res)
	}
}
