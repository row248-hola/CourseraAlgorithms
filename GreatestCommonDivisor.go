package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

func calculateGreatestCommonDivisor(n1 *big.Int, n2 *big.Int) *big.Int {
	if len(n2.Bits()) == 0 {
		return n1
	}

	mod := n1.Mod(n1, n2)
	return calculateGreatestCommonDivisor(n2, mod)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fib, _ := reader.ReadString('\n')
	fib = strings.TrimSpace(fib)

	stringSlice := strings.Split(fib, " ")

	bigIntSlice := make([]*big.Int, len(stringSlice))
	for i, str := range stringSlice {
		bigInt, ok := new(big.Int).SetString(str, 10)
		if !ok {
			fmt.Println("Can't cast", str, "to big int")
			continue
		}

		bigIntSlice[i] = bigInt
	}

	result := calculateGreatestCommonDivisor(bigIntSlice[0], bigIntSlice[1])
	fmt.Println(result)
}
