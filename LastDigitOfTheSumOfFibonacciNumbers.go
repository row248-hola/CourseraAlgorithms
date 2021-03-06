package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

func lastDigitOfTheSumOfFibonacciNumbers(num *big.Int) int {
	sequence := []int{0, 1, 2, 4, 7, 2, 0, 3, 4, 8, 3, 2, 6, 9, 6, 6, 3, 0, 4, 5, 0, 6,
		7, 4, 2, 7, 0, 8, 9, 8, 8, 7, 6, 4, 1, 6, 8, 5, 4, 0, 5, 6, 2, 9, 2, 2, 5, 8, 4,
		3, 8, 2, 1, 4, 6, 1, 8, 0, 9, 0}

	modulo := new(big.Int).Mod(num, new(big.Int).SetInt64(int64(len(sequence))))

	return sequence[modulo.Int64()]
}

// Task. Given an integer ๐, find the last digit of the sum ๐น0 +๐น1 +ยทยทยท+๐น๐.
// Input Format. The input consists of a single integer ๐.
// Constraints. 0 โค ๐ โค 10 ** 14.
// Output Format. Output the last digit of ๐น0 + ๐น1 + ยท ยท ยท + ๐น๐.
func main() {
	reader := bufio.NewReader(os.Stdin)

	fib, _ := reader.ReadString('\n')
	fib = strings.TrimSpace(fib)

	num, ok := new(big.Int).SetString(fib, 10)
	if !ok {
		fmt.Println("Can't cast", fib, "to big.Int")
	}

	fmt.Println(lastDigitOfTheSumOfFibonacciNumbers(num))
}
