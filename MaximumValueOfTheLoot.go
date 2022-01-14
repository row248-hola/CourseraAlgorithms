package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Item struct {
	value        int
	weight       int
	valuePerItem float64
}

func calculateMaximumValue(sortedItems []*Item, capacity int) float64 {
	filled := 0
	value := 0.0

	for _, item := range sortedItems {
		if filled == capacity {
			break
		}

		if filled+item.weight <= capacity {
			filled += item.weight
			value += float64(item.value)
		} else {
			remainingWeight := capacity - filled
			fraction := float64(remainingWeight) / float64(item.weight)

			value += float64(item.value) * fraction
			break
		}
	}

	return value
}

// Task. The goal of this code problem is to implement an algorithm for the fractional knapsack problem.
// Input Format. The first line of the input contains the number 𝑛 of items and the capacity 𝑊 of a knapsack.
//	The next 𝑛 lines define the values and weights of the items. The 𝑖-th line contains integers 𝑣𝑖 and 𝑤𝑖—the value and the weight of 𝑖-th item, respectively.
// Constraints. 1≤𝑛≤103,0≤𝑊 ≤2·106;0≤𝑣𝑖 ≤2·106,0<𝑤𝑖 ≤2·106 forall1≤𝑖≤𝑛. All the numbers are integers.
// Output Format. Output the maximal value of fractions of items that fit into the knapsack.
// 	The absolute value of the difference between the answer of your program and the optimal value should be at most 10−3.
//	To ensure this, output your answer with at least four digits after the decimal point
//	(otherwise your answer, while being computed correctly, can turn out to be wrong because of rounding issues).
func main() {
	inputData := ""

	reader := bufio.NewReader(os.Stdin)

	countAndCapacity, _ := reader.ReadString('\n')
	inputData += countAndCapacity // debug

	args := strings.Split(strings.TrimSpace(countAndCapacity), " ")

	// parse the data
	itemsNum, _ := strconv.Atoi(args[0])
	capacity, _ := strconv.Atoi(args[1])

	var items []*Item
	for itemsNum > 0 {
		str, _ := reader.ReadString('\n')
		inputData += str // debug
		args := strings.Split(strings.TrimSpace(str), " ")

		value, _ := strconv.Atoi(args[0])
		weight, _ := strconv.Atoi(args[1])

		items = append(items, &Item{
			value:        value,
			weight:       weight,
			valuePerItem: float64(value) / float64(weight),
		})

		itemsNum -= 1
	}

	// sort the data
	sort.Slice(items, func(i, j int) bool {
		return items[i].valuePerItem > items[j].valuePerItem
	})

	fmt.Println(calculateMaximumValue(items, capacity))
}
