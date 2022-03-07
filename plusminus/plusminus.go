package plusminus

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
	// 1 2 3 -1 -2 -3 0 0

	// Output
	// 	0.375000
	// 0.375000
	// 0.250000

	var pos int32 = 0
	var neg int32 = 0
	var zero int32 = 0

	for _, v := range arr {
		if v > 0 {
			pos++
		} else if v < 0 {
			neg++
		} else {
			zero++
		}
	}

	fmt.Printf("%.6f\n", float64(pos)/float64(len(arr)))
	fmt.Printf("%.6f\n", float64(neg)/float64(len(arr)))
	fmt.Printf("%.6f\n", float64(zero)/float64(len(arr)))

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)
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
