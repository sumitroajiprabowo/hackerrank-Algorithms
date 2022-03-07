package birthdaycakecandle

import (
	"fmt"
	"testing"
)

func TestBirthdayCakeCandle(t *testing.T) {

	var candles []int32

	candles = []int32{3, 2, 1, 3}

	fmt.Println(birthdayCakeCandles(candles))

}
