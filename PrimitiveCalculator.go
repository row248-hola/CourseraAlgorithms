package main

import (
	"bufio"
	"fmt"
	"math"
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

func getMinimumSteps(value int) (int, string) {
	var dp = make([]int, value+1)

	for i := 1; i <= value; i++ {
		dp[i] = math.MaxInt32

		if i%3 == 0 {
			dp[i] = int(math.Min(float64(dp[i]), float64(dp[i/3]+1)))
		}
		if i%2 == 0 {
			dp[i] = int(math.Min(float64(dp[i]), float64(dp[i/2]+1)))
		}
		dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-1]+1)))
	}

	var count = dp[value] - 1

	var sequence []int
	sequence = append(sequence, 1)

	n := count
	for n != 0 {
		sequence = append(sequence, value)
		n--

		if value%3 == 0 && dp[value/3] == n+1 {
			value /= 3
		} else if value%2 == 0 && dp[value/2] == n+1 {
			value /= 2
		} else {
			value -= 1
		}
	}

	sort.Slice(sequence, func(i, j int) bool {
		return sequence[i] < sequence[j]
	})

	var stringSequence []string
	for _, value := range sequence {
		parsedInt := strconv.Itoa(value)
		stringSequence = append(stringSequence, parsedInt)
	}

	return count, strings.Join(stringSequence, " ")
}

// Task. Given an integer π, compute the minimum number of operations needed to obtain the number π starting from the number 1.
// Input Format. The input consists of a single integer 1 β€ π β€ 106.
// Output Format. In the first line, output the minimum number π of operations needed to get π from 1. In the second line output a sequence of intermediate numbers. That is, the second line should contain positiveintegersπ0,π2,...,ππβ1 suchthatπ0 =1,ππβ1 =πandforall0β€π<πβ1,ππ+1 isequalto either ππ + 1, 2ππ, or 3ππ. If there are many such sequences, output any one of them.
func main() {
	desiredValue := readInt()

	count, sequence := getMinimumSteps(desiredValue)

	fmt.Println(count)
	fmt.Println(sequence)
}
