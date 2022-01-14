package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func compute(num int) *big.Int {
	results := make(map[int]*big.Int, num)
	results[0] = big.NewInt(0)
	results[1] = big.NewInt(1)

	for i := 2; i <= num; i++ {
		results[i] = new(big.Int).Add(results[i-1], results[i-2])
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

	fmt.Println(compute(num))
}
