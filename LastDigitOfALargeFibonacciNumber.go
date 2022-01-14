package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func computeLastDigit(num int) *big.Int {
	results := make(map[int]*big.Int, num)
	results[0] = big.NewInt(0)
	results[1] = big.NewInt(1)

	delimiter := big.NewInt(10)

	for i := 2; i <= num; i++ {
		sumOfModulus := new(big.Int).Add(results[i-1], results[i-2])
		results[i] = new(big.Int).Mod(sumOfModulus, delimiter)
	}

	return results[num]
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fib, _ := reader.ReadString('\n')
	fib = strings.TrimSpace(fib)

	num, err := strconv.Atoi(fib)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(computeLastDigit(num))
}
