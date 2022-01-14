package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
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

func calculate(fibNum *big.Int, divider int) *big.Int {
	// find sequence for `f mod divider`
	var sequence []int
	for i := 0; ; i++ {
		fib := getFibonacci(i)
		modulo := new(big.Int).Mod(fib, new(big.Int).SetInt64(int64(divider))).Int64()

		// end of the sequence
		if i > 1 && modulo == 1 && sequence[i-1] == 0 {
			sequence = sequence[0 : i-1]
			break
		}

		sequence = append(sequence, int(modulo))
	}

	newFib := new(big.Int).Mod(fibNum, new(big.Int).SetInt64(int64(len(sequence))))
	fib := getFibonacci(int(newFib.Int64()))

	result := new(big.Int).Mod(fib, new(big.Int).SetInt64(int64(divider)))

	return result
}

// Task. Given two integers 𝑛 and 𝑚, output 𝐹𝑛 mod 𝑚 (that is, the remainder of 𝐹𝑛 when divided by 𝑚).
// Input Format. The input consists of two integers 𝑛 and 𝑚 given on the same line (separated by a space).
// Constraints. 1≤𝑛≤1014,2≤𝑚≤103.
// Output Format. Output 𝐹𝑛 mod 𝑚.
func main() {
	reader := bufio.NewReader(os.Stdin)

	fib, _ := reader.ReadString('\n')
	fib = strings.TrimSpace(fib)

	args := strings.Split(fib, " ")

	fibonacci, ok := new(big.Int).SetString(args[0], 10)
	if !ok {
		fmt.Println("Can't cast", args[0], "to big int")
	}

	divider, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Can't cast", args[1], "to int: ", err)
	}

	fmt.Println(calculate(fibonacci, divider))
}
