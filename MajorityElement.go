package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

// fill values to hashmap
// sort hashmap values
// compare first value of sorted hashmap
func hasMajorElement(nums []int) int {
	size := len(nums)
	occurrence := make(map[int]int, size)

	for _, n := range nums {
		if _, ok := occurrence[n]; ok {
			occurrence[n]++
		} else {
			occurrence[n] = 1
		}
	}

	var values []int
	for k, _ := range occurrence {
		values = append(values, k)
	}

	sort.Slice(values, func(i, j int) bool {
		return occurrence[values[i]] > occurrence[values[j]]
	})

	if occurrence[values[0]] > size/2 {
		return 1
	}
	return 0
}

// Task. The goal in this code problem is to check whether an input sequence contains a majority element.
// Input Format. The first line contains an integer 𝑛, the next one contains a sequence of 𝑛 non-negative integers 𝑎0, 𝑎1, . . . , 𝑎𝑛−1.
// Constraints. 1≤𝑛≤105;0≤𝑎𝑖 ≤109 forall0≤𝑖<𝑛.
// Output Format. Output 1 if the sequence contains an element that appears strictly more than 𝑛/2 times, and 0 otherwise.
func main() {
	readInt()
	nums := readInts()

	fmt.Println(hasMajorElement(nums))
}
