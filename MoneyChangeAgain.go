package main

import (
	"bufio"
	"fmt"
	"math"
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

var coins = []int{1, 3, 4}

func getMinimumCoinsDpVersion(money int) int {
	var minCoins = make([]int, money+1)

	for m := 1; m <= money; m++ {
		minCoins[m] = math.MaxInt32
		minCoinsLocal := 0

		for _, coin := range coins {
			if m >= coin {
				minCoinsLocal = minCoins[m-coin] + 1

				if minCoinsLocal < minCoins[m] {
					minCoins[m] = minCoinsLocal
				}
			}
		}
	}

	return minCoins[money]
}

// Input Format. Integer money.
// Output Format. The minimum number of coins with denominations 1, 3, 4 that changes money.
// Constraints. 1 ≤ money ≤ 10 ** 3.
func main() {
	sum := readInt()

	fmt.Println(getMinimumCoinsDpVersion(sum))
}
