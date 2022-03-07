package averybigsum

import "testing"

func TestAVeryBigSum(t *testing.T) {
	var a []int64
	a = []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}
	res := aVeryBigSum(a)
	if res != 5000000015 {
		t.Error("Expected 5000000015, got ", res)
	}
}
