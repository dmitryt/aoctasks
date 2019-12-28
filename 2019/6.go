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

func part1(content []byte) {
	storage := make(map[string]string)
	orbits := strings.Split(string(content), "\n")
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

func main() {
	content, _ := ioutil.ReadFile("6.txt")
	part1(content)
}
