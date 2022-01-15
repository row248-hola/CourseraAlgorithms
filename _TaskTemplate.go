package main

func readInt(reader *bufio.Reader) int {
	str, _ := reader.ReadString('\n')
	result, _ := strconv.Atoi(strings.TrimSpace(str))

	return result
}

func readString(reader *bufio.Reader) string {
	str, _ := reader.ReadString('\n')
	return str
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
