package main

import (
	"io/ioutil"
	"strings"
	"strconv"
	"fmt"
)

func part1(codes []string) []string {
	index := 0
	buffer := []int{}
	toInt := func(str string) int {
		i, _ := strconv.Atoi(str)
		return i
	}
	for _, str := range codes {
		num := toInt(str)
		if index % 4 == 0 && num == 99 {
			return codes
		}
		if (index + 1) % 4 == 0 {
			result := 0
			if (buffer[0] == 1) {
				result = toInt(codes[buffer[1]]) + toInt(codes[buffer[2]])
			} else if (buffer[0] == 2) {
				result = toInt(codes[buffer[1]]) * toInt(codes[buffer[2]])
			}
			codes[num] = strconv.Itoa(result)
			buffer = []int{}
			index = 0
		} else {
			buffer = append(buffer, num)
			index = index + 1
		}
	}
	return codes
}

func part2(pCodes *[]string) {
	codes := *pCodes
	const min = 0
	const max = 100
	// i := 12
	// j := 2
	expected := "19690720"
	for i := min; i < max; i = i + 1 {
		for j := min; j < max; j = j + 1 {
			normalzedCodes := append([]string{codes[0], strconv.Itoa(i), strconv.Itoa(j)}, codes[3:]...)
			if part1(normalzedCodes)[0] == expected {
				fmt.Println(100 * i + j)
			}
		}
	}
}

func main() {
	content, _ := ioutil.ReadFile("4.txt")
	codes := strings.Split(string(content), ",")
	part2(&codes)
}
