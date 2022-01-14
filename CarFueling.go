package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInt(reader *bufio.Reader) int {
	str, _ := reader.ReadString('\n')
	result, _ := strconv.Atoi(strings.TrimSpace(str))

	return result
}

func readInts(reader *bufio.Reader) []int {
	str, _ := reader.ReadString('\n')
	args := strings.Split(strings.TrimSpace(str), " ")

	result := make([]int, len(args))
	for i, str := range args {
		result[i], _ = strconv.Atoi(str)
	}

	return result
}

func calculateCarFueling(totalDistance, maxDistance int, stations []int) int {
	lastRefill := 0
	result := 0
	distance := 0

	totalStations := len(stations)

	for i, stationDistance := range stations {

		nextStationDistance := 0
		if i == totalStations-1 {
			nextStationDistance = totalDistance
		} else {
			nextStationDistance = stations[i+1]
		}

		// can't get further without refilling
		if nextStationDistance-stations[lastRefill] > maxDistance {
			// can't get further at all
			if stationDistance-stations[lastRefill] > maxDistance {
				break
			}

			result += 1
			lastRefill = i
			distance = stationDistance
		}
	}

	// not enough gas stating along the road
	if distance+maxDistance < totalDistance {
		return -1
	}

	return result
}

// Input Format. The first line contains an integer 𝑑. The second line contains an integer 𝑚.
//  The third line specifies an integer 𝑛. Finally, the last line contains integers stop1, stop2, . . . , stop𝑛.
// Output Format. Assuming that the distance between the cities is 𝑑 miles, a car can travel at most 𝑚 miles on a full tank,
//  and there are gas stations at distances stop1 , stop2 , . . . , stop𝑛 along the way, output the minimum number of refills needed.
//  Assume that the car starts with a full tank. If it is not possible to reach the destination, output −1.
// Constraints. 1≤𝑑≤105.1≤𝑚≤400.1≤𝑛≤300.0<stop1 <stop2 <···<stop𝑛 <𝑑.
func main() {
	reader := bufio.NewReader(os.Stdin)

	totalDistance := readInt(reader)
	maxDistance := readInt(reader)
	readInt(reader)
	stations := readInts(reader)

	stations = append([]int{0}, stations...)

	// output maximum value
	fmt.Println(calculateCarFueling(totalDistance, maxDistance, stations))
}
