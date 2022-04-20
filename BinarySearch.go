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

func binarySearch(list []int, search int) int {
	from := 0
	to := len(list) - 1

	for from <= to {
		index := (to + from) / 2
		guess := list[index]

		if guess == search {
			return index
		}

		if guess > search {
			to = index - 1
		} else {
			from = index + 1
		}
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
		result = append(result, binarySearch(sortedList, number))
	}

	var sb strings.Builder
	for _, value := range result {
		sb.WriteString(fmt.Sprint(" ", value))
	}

	fmt.Println(strings.TrimSpace(sb.String()))
}
