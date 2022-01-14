package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	_, _ = reader.ReadString('\n')
	text, _ := reader.ReadString('\n')

	numbersString := strings.Replace(text, "\n", "", -1)

	stringSlice := strings.Split(numbersString, " ")

	bigIntSlice := make([]*big.Int, len(stringSlice))
	for i, number := range stringSlice {
		bigInt, ok := new(big.Int).SetString(number, 10)
		if !ok {
			fmt.Println("can't covert", bigInt)
		}

		bigIntSlice[i] = bigInt
	}

	sort.Sort(sort.IntSlice{})

	sort.Slice(bigIntSlice, func(i, j int) bool {
		//fmt.Println("Comparing",  bigIntSlice[i], "with",  bigIntSlice[j])
		compareResult := bigIntSlice[i].Cmp(bigIntSlice[j])

		if compareResult == 1 {
			return true
		}
		return false
	})

	//fmt.Println("Sorted bigIntSlice:", bigIntSlice)

	if len(bigIntSlice) == 1 {
		fmt.Println(bigIntSlice[0])
		return
	}

	val1 := bigIntSlice[0]
	val2 := bigIntSlice[1]

	result := val1.Mul(val1, val2)
	fmt.Println(result)
}
