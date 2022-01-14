package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

func greatestCommonDivisor(n1 *big.Int, n2 *big.Int) *big.Int {
	if len(n2.Bits()) == 0 {
		return n1
	}

	mod := new(big.Int).Mod(n1, n2)
	return greatestCommonDivisor(n2, mod)
}

func calculateLeastCommonMultiple(n1 *big.Int, n2 *big.Int) *big.Int {
	divisor := greatestCommonDivisor(n1, n2)
	return new(big.Int).Div(new(big.Int).Mul(n1, n2), divisor)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	nums, _ := reader.ReadString('\n')
	nums = strings.TrimSpace(nums)

	stringSlice := strings.Split(nums, " ")

	bigIntSlice := make([]*big.Int, len(stringSlice))
	for i, str := range stringSlice {
		num, ok := new(big.Int).SetString(str, 10)
		if !ok {
			fmt.Println("cant parse", str, "as big int")
			return
		}

		bigIntSlice[i] = num
	}

	result := calculateLeastCommonMultiple(bigIntSlice[0], bigIntSlice[1])
	fmt.Println(result)
}
