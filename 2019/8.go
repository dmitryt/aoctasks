package main

import (
	"io/ioutil"
	"strings"
	"sort"
	"fmt"
)

type Layer struct {
	arr [][]string
	stats map[string]int
}

func split(numbers *[]string, width int, height int) []Layer {
	result := []Layer{}
	layer := Layer{stats: map[string]int{}}
	size := width * height
	row := []string{}
	for index, num := range *numbers {
		if index > 0 && index % width == 0 {
			layer.arr = append(layer.arr, row)
			row = []string{num}
			if index % size == 0 {
				result = append(result, layer)
				layer = Layer{stats: map[string]int{}}
			}
		} else {
			row = append(row, num)
		}
		layer.stats[num] = layer.stats[num] + 1
	}
	if len(row) > 0 {
		layer.arr = append(layer.arr, row)
	}
	if len(layer.arr) > 0 {
		result = append(result, layer)
	}
	return result
}

func part8_1(numbers *[]string, width int, height int) {
	layers := split(numbers, width, height)
	sort.Slice(layers, func(i, j int) bool {
		return layers[i].stats["0"] < layers[j].stats["0"]
	})
	fmt.Println(layers[0])
	fmt.Println(layers[0].stats["1"] * layers[0].stats["2"])
}

func part8_2(numbers *[]string, width int, height int) {
	layers := split(numbers, width, height)
	result := [][]string{}
	for j := 0; j < height; j = j + 1 {
		result = append(result, []string{})
		for i := 0; i < width; i = i + 1 {
			for _, layer := range layers {
				if layer.arr[j][i] != "2" {
					result[j] = append(result[j], layer.arr[j][i]);
					break
				}
			}
			if result[j][i] == "" {
				result[j] = append(result[j], "2");
			}
		}
	}
	for _, row := range result {
		fmt.Println(row)
	}
	fmt.Println("DETAILS")
	for _, layer := range layers {
		for _, row := range layer.arr {
			fmt.Println(row)
		}
		fmt.Println("======================")
	}
}

func run8() {
	content, _ := ioutil.ReadFile("8.txt")
	numbers := strings.Split(string(content), "")
	part8_2(&numbers, 25, 6)
}
