package simplearraysum

import "testing"

func TestSimpleArraySum(t *testing.T) {
	var a []int32
	a = []int32{1, 2, 3, 4, 10, 11}
	res := simpleArraySum(a)
	if res != 31 {
		t.Error("Expected 31, got ", res)
	}

}
