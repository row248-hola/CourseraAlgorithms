package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func readInt() int {
	str, _ := reader.ReadString('\n')
	result, _ := strconv.Atoi(strings.TrimSpace(str))

	return result
}

func readInts() []int {
	str, _ := reader.ReadString('\n')
	args := strings.Split(strings.TrimSpace(str), " ")

	result := make([]int, len(args))
	for i, str := range args {
		result[i], _ = strconv.Atoi(str)
	}

	return result
}

func binarySearchLeftmost(list []int, search int) int {
	from := 0
	to := len(list) - 1

	for from < to {
		index := (to + from) / 2

		if list[index] < search {
			from = index + 1
		} else {
			to = index
		}
	}

	if list[from] == search {
		return from
	}
	return -1
}

func main() {
	readInt()
	sortedList := readInts()
	readInt()
	numbersToFind := readInts()

	var result []int
	for _, number := range numbersToFind {
		result = append(result, binarySearchLeftmost(sortedList, number))
	}

	var sb strings.Builder
	for _, value := range result {
		sb.WriteString(fmt.Sprint(" ", value))
	}

	fmt.Println(strings.TrimSpace(sb.String()))
}
