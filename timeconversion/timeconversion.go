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
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {

	// Sample input 07:05:45PM
	// Output 19:05:45

	// Write your code here

	time := strings.Split(s, "")
	firstSlice := strings.Split(s, ":")[0]

	meridiem := time[len(time)-2]
	var left string
	var i int

	hour, err := strconv.Atoi(firstSlice)

	if meridiem == "P" {
		if err == nil {
			if hour < 12 {
				i = hour + 12
				left = strconv.Itoa(i)
			} else if hour >= 24 {
				i = hour - 24
				if i < 10 {
					left += "0"
				}
				left += strconv.Itoa(i)
			} else {
				left = firstSlice
			}
		}
	}

	if meridiem == "A" {
		if err == nil {
			if hour >= 12 {
				i = hour - 12
				left += "0"
				left += strconv.Itoa(i)
			} else {
				left = firstSlice
			}
		}
	}

	trim := strings.Join(time[:len(time)-2], "")
	right := strings.Split(trim, firstSlice)[1]
	final := left + right

	return final

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

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
