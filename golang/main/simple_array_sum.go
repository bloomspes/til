package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func simpleArraySum(numbers []int32) int32 {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += i
	}
	return int32(sum)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	numbersCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	temp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var numbers []int32

	for i := 0; i < int(numbersCount); i++ {
		numbersTemp, err := strconv.ParseInt(temp[i], 10, 64)
		checkError(err)
		number := int32(numbersTemp)
		numbers = append(numbers, number)
	}

	result := simpleArraySum(numbers)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")

}
