package main

import (
	"io/ioutil"
	"strconv"
	"strings"
	"regexp"
	"math"
	"time"
	"fmt"
)

type Moon struct {
	pos [3]int
	vel [3]int
}


func Equal(a, b [3]int) bool {
	for i, v := range a {
			if v != b[i] {
					return false
			}
	}
	return true
}

func EqualMoons(a, b []Moon) bool {
	for i, v := range a {
			if !(Equal(v.pos, b[i].pos) && Equal(v.vel, b[i].vel)) {
					return false
			}
	}
	return true
}

func (m *Moon) energy() int {
	pot := 0
	for _, v := range m.pos {
		pot = pot + int(math.Abs(float64(v)))
	}
	kin := 0
	for _, v := range m.vel {
		kin = kin + int(math.Abs(float64(v)))
	}
	return pot * kin
}

func iterate(moons []Moon) ([]Moon, int) {
	result := []Moon{}
	changePos := func(moon *Moon) {
		for i, _ := range moon.pos {
			moon.pos[i] = moon.pos[i] + moon.vel[i]
		}
	}
	changeVel := func(baseMoon *Moon, moon Moon) {
		for i, coord := range baseMoon.pos {
			if coord > moon.pos[i] {
				baseMoon.vel[i] = baseMoon.vel[i] - 1
			} else if coord < moon.pos[i] {
				baseMoon.vel[i] = baseMoon.vel[i] + 1
			}
		}
	}
	energy := 0
	for i, moon := range moons {
		for j, m := range moons {
			if i == j {
				continue
			}
			changeVel(&moon, m)
		}
		changePos(&moon)
		energy = energy + moon.energy()
		result = append(result, moon)
	}
	return result, energy
}

func part1(moons *[]Moon) {
	tmp := *moons
	for i := 0; i < 1000; i = i + 1 {
		tmp, _ = iterate(tmp)
	}
	result := 0
	for _, moon := range tmp {
		result = result + moon.energy()
	}
	fmt.Println(result)
}

func part2(moons *[]Moon) {
	tmp := *moons
	_, firstEnergy := iterate(*moons)
	result := -1
	start := time.Now()
	for i := 0;; i = i + 1 {
		tmpEnergy := 0
		tmp, tmpEnergy = iterate(tmp)
		if i % 10000000 == 0 {
			fmt.Println(i)
			fmt.Println("Took %s", time.Since(start))
			start = time.Now()
		}
		if i > 0 && firstEnergy == tmpEnergy {
			result = i
			break
		}
	}
	fmt.Println(result)
}

func preparePos(arr []string) [3]int {
	var pos [3]int
	slice := []int{}
	for _, i := range arr {
		j, err := strconv.Atoi(i)
		if err != nil {
				panic(err)
		}
		slice = append(slice, j)
	}
	copy(pos[:], slice[:3])
	return pos
}

func main() {
	content, _ := ioutil.ReadFile("12.txt")
	moonsData := strings.Split(string(content), "\n")
	moons := []Moon{}
	pattern := regexp.MustCompile(`^<x=(-?\d+),\s+y=(-?\d+),\s+z=(-?\d+)>$`)
	for _, data := range moonsData {
		moons = append(moons, Moon{pos: preparePos(pattern.FindAllStringSubmatch(data, -1)[0][1:])})
	}
	part2(&moons)
}
