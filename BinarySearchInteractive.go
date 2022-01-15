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

func readString() string {
	str, _ := reader.ReadString('\n')
	return str
}

func askUser(number int) string {
	fmt.Println("Is", number, "= X?")
	return readString()
}

func search(numbers int) {
	from := 0.0
	to := float64(numbers)

	for to > from {
		guessedNumber := math.Round((to-from)/2) + from

		answer := askUser(int(guessedNumber))
		if strings.HasPrefix(answer, ">") {
			from = guessedNumber
		} else if strings.HasPrefix(answer, "<") {
			to = guessedNumber
		} else if strings.HasPrefix(answer, "=") {
			fmt.Println("The number was found. End of the program")
			return
		} else {
			fmt.Println("You should type of these symbols: <, =, >")
		}
	}
}

// Performs binary search in N sorted numbers asking each time whether X is found.
// The users must answer by one of these symbols: < (less), = (equal), > (greater)
func main() {
	numbers := readInt()
	search(numbers)
}
