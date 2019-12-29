package main

import (
	"io/ioutil"
	"strings"
	"fmt"
)

type Vertex struct {
	value string;
	children []*Vertex;
}

func part1(orbits []string) {
	storage := make(map[string]string)
	for _, orbit := range orbits {
		parsed := strings.Split(string(orbit), ")")
		storage[parsed[1]] = parsed[0]
	}
	calculate := func(node string, cnt *int) {
		current := node
		for {
			_, ok := storage[current]
			if !ok {
				return
			}
			*cnt = *cnt + 1
			current = storage[current]
		}
	}
	result := 0
	for child := range storage {
		calculate(child, &result)
	}
	fmt.Println(result)
}

func part2(orbits []string) {
	storage := make(map[string]string)
	for _, orbit := range orbits {
		parsed := strings.Split(string(orbit), ")")
		storage[parsed[1]] = parsed[0]
	}
	indexOf := func(arr []string, el string) int {
		result := -1
		for index, item := range arr {
			if item == el {
				return index
			}
		}
		return result
	}

	collectParents := func(node string) []string {
		result := []string{}
		current := node
		for {
			parent, ok := storage[current]
			if !ok {
				return result
			}
			result = append(result, parent)
			current = parent
		}
		return result
	}
	left := collectParents("YOU")
	right := collectParents("SAN")
	result := []string{}
	for lIndex, el := range left {
		rIndex := indexOf(right, el)
		if rIndex != -1 {
			result = append(left[:lIndex + 1], right[:rIndex]...)
			break
		}
	}
	fmt.Println(len(result) - 1)
}

func main() {
	content, _ := ioutil.ReadFile("6.txt")
	part2(strings.Split(string(content), "\n"))
}
