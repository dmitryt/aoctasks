package main

import (
	"strings"
	"strconv"
	"sort"
	"math"
	"fmt"
)

type Point [3]int;

type Line struct {
	start Point
	end Point
	steps int
}

func sortNumbers(n1 int, n2 int) (int, int) {
	tmp := []int{n1, n2}
	sort.Ints(tmp)
	return tmp[0], tmp[1]
}

func findIntersections(horLines []Line, verLines []Line) []Point {
  result := []Point{}
	for _, hLine := range horLines {
		x1, x2 := sortNumbers(hLine.start[0], hLine.end[0])
		y := hLine.start[1]
		for _, vLine := range verLines {
			y1, y2 := sortNumbers(vLine.start[1], vLine.end[1])
			x := vLine.start[0]
			dx := int(math.Abs(float64(x - hLine.start[0])))
			dy := int(math.Abs(float64(y - vLine.start[1])))
			if (x > x1 && x < x2 && y > y1 && y < y2) {
				result = append(result, Point{x, y, hLine.steps + vLine.steps + dx + dy})
      }
		}
	}
  return result;
}

func collectPaths(sequence []string) ([]Line, []Line) {
	hor := []Line{}
	ver := []Line{}
	current := Point{}
	for _, el := range sequence {
		dir := el[0]
		step, _ := strconv.Atoi(el[1:])
		if dir == 'L' || dir == 'R' {
			start := current
			if dir == 'L' {
				current[0] = current[0] - step
			} else {
				current[0] = current[0] + step
			}
			end := current
			hor = append(hor, Line{ start: start, end: end, steps: current[2] })
		}
		if dir == 'U' || dir == 'D' {
			start := current
			if dir == 'D' {
				current[1] = current[1] - step
			} else {
				current[1] = current[1] + step
			}
			end := current
			ver = append(ver, Line{ start: start, end: end, steps: current[2] })
		}
		current[2] = current[2] + step
	}
	return hor, ver
}

func findMinDistance(seq1 []string, seq2 []string) float64 {
	hor1, ver1 := collectPaths(seq1)
	hor2, ver2 := collectPaths(seq2)
	intersections := findIntersections(hor1, ver2)
	intersections = append(intersections, findIntersections(hor2, ver1)...)
	distances := make([]float64, len(intersections))
	minDistance := math.Inf(1)
	for idx, point := range intersections {
		distances[idx] = math.Abs(float64(point[0])) + math.Abs(float64(point[1]))
		if distances[idx] < minDistance {
			minDistance = distances[idx]
		}
	}
	return minDistance
}

func findMinSteps(seq1 []string, seq2 []string) float64 {
	hor1, ver1 := collectPaths(seq1)
	hor2, ver2 := collectPaths(seq2)
	intersections := findIntersections(hor1, ver2)
	intersections = append(intersections, findIntersections(hor2, ver1)...)
	minStep := math.Inf(1)
	for _, point := range intersections {
		if float64(point[2]) < minStep {
			minStep = float64(point[2])
		}
	}
	return minStep
}

func run3() {
	str1 := "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51";
	str2 := "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7";
	seq1 := strings.Split(str1, ",");
	seq2 := strings.Split(str2, ",");
	fmt.Println(findMinSteps(seq1, seq2));
}