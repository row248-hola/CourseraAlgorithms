package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculateMoneyChange(money int) int {
	coinsDenominations := []int{10, 5, 1}
	coins := 0

	for money > 0 {
		if coinsDenominations[0] > money {
			coinsDenominations = coinsDenominations[1:]
			continue
		}

		money -= coinsDenominations[0]
		coins++
	}

	return coins
}

// Task. The goal in this problem is to find the minimum number of coins needed to change the input value (an integer) into coins with denominations 1, 5, and 10.
// Input Format. The input consists of a single integer 𝑚.
// Constraints. 1 ≤ 𝑚 ≤ 103.
// Output Format. Output the minimum number of coins with denominations 1, 5, 10 that changes 𝑚.
func main() {
	reader := bufio.NewReader(os.Stdin)

	fib, _ := reader.ReadString('\n')
	fib = strings.TrimSpace(fib)

	num, err := strconv.Atoi(fib)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(calculateMoneyChange(num))
}
