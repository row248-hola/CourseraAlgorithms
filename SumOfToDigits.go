package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	numbersString := strings.Replace(text, "\n", "", -1)

	numbers := strings.Split(numbersString, " ")
	result := 0
	for i := 0; i < len(numbers); i++ {

		num, err := strconv.Atoi(numbers[i])
		if err != nil {
			fmt.Println(err)
		}

		result += num
	}

	fmt.Println(result)
}
