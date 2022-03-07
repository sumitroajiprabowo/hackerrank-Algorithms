package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'sherlockAndMinimax' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY arr
 *  2. INTEGER p
 *  3. INTEGER q
 */

func sherlockAndMinimax(arr []int32, p int32, q int32) int32 {
	// Input : 3
	// 5 8 14
	// 4 9

	// Output : 4
	// Write your code here
	var max int32 = 0
	var count int32 = 0

	for _, candle := range arr {
		if candle > max {
			max = candle
		}
	}

	for _, candle := range arr {
		if candle == max {
			count++
		}
	}

	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

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

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	pTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	p := int32(pTemp)

	qTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	q := int32(qTemp)

	result := sherlockAndMinimax(arr, p, q)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
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
