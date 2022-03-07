package minimaxsum

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func miniMaxSum(arr []int32) {
	//Given five positive integers, find the minimum and maximum values that can be calculated by summing exactly four of the five integers. Then print the respective minimum and maximum values as a single line of two space-separated long integers.
	// For example, arr = [1, 3, 5, 7, 9]. Our initial
	// The minimum sum is 1+3+5+7 = 16 and the maximum sum is 3+5+7+9 = 24.
	// We would print 16 24
	// Function to find the sum of all elements in the array

	// input 256741038 623958417 467905213 714532089 938071625
	// expected output 2063136757 2744467344

	var sum int64
	for i := 0; i < len(arr); i++ {
		sum = sum + int64(arr[i])
	}
	var min int64
	var max int64
	for i := 0; i < len(arr); i++ {
		if sum-int64(arr[i]) <= min || min == 0 {
			min = sum - int64(arr[i])
		}
		if sum-int64(arr[i]) > max || max == 0 {
			max = sum - int64(arr[i])
		}
	}
	fmt.Printf("%d %d", min, max)

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
