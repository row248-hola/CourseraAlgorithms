package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

func getFibonacci(num int) *big.Int {
	results := make(map[int]*big.Int, num)
	results[0] = big.NewInt(0)
	results[1] = big.NewInt(1)

	for i := 2; i <= num; i++ {
		results[i] = new(big.Int).Add(results[i-1], results[i-2])
	}

	return results[num]
}

func findPisanoPeriod(divider *big.Int) {
	var period []int
	for i := 0; ; i++ {
		fib := getFibonacci(i)
		modulo := new(big.Int).Mod(fib, divider).Int64()

		// end of the period
		if i > 1 && modulo == 1 && period[i-1] == 0 {
			period = period[0 : i-1]
			break
		}

		period = append(period, int(modulo))
	}

	fmt.Println("Pisano period for modulo", divider, "is:", period)
}

func findPisanoPeriodV2(from *big.Int) {
	var period []int
	for i := from; ; i.Add(i, new(big.Int).SetInt64(1)) {
		fib := getFibonacci(i)
		modulo := new(big.Int).Mod(fib, divider).Int64()

		// end of the period
		if i > 1 && modulo == 1 && period[i-1] == 0 {
			period = period[0 : i-1]
			break
		}

		period = append(period, int(modulo))
	}

	fmt.Println("Pisano period for modulo", divider, "is:", period)
}

// Task. Given two non-negative integers ๐ and ๐, where ๐ โค ๐, find the last digit of the sum ๐น๐ + ๐น๐+1 + ยทยทยท+๐น๐.
// Input Format. The input consists of two non-negative integers ๐ and ๐ separated by a space.
// Constraints. 0 โค ๐ โค ๐ โค 1014.
// Output Format. Output the last digit of ๐น๐ + ๐น๐+1 + ยท ยท ยท + ๐น๐.
func main() {
	reader := bufio.NewReader(os.Stdin)

	fib, _ := reader.ReadString('\n')
	nums := strings.Split(strings.TrimSpace(fib), " ")

	if len(nums) != 2 {
		fmt.Println("Provide two numbers")
		return
	}

	from, _ := new(big.Int).SetString(nums[0], 10)
	to, _ := new(big.Int).SetString(nums[1], 10)

	findPisanoPeriod(from)
	findPisanoPeriod(to)
}
