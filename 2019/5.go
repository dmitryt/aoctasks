package main

import (
	"io/ioutil"
	"strings"
	"strconv"
	"bufio"
	"os"
	"fmt"
)

func toInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

func parseCode(num int) [4]int {
	result := [4]int{}
	if (num > 100) {
		str := strconv.Itoa(num)
		nums := strings.Split(str[:len(str) - 2], "")
		for i := 0; i < len(nums); i = i + 1 {
			result[3 - len(nums) + i] = toInt(nums[i])
		}
		result[3] = toInt(str[(len(str) - 2):])
	} else {
		result[3] = num
	}
	return result
}

func getInstruction(mode int, value int, codes *[]string) string {
	if mode == 0 {
		return (*codes)[value]
	}
	if mode == 1 {
		return strconv.Itoa(value)
	}
	return "0"
}

func part5(codes []string, input string) []string {
	pos := 0
	index := 0
	buffer := []int{}
	instructionsLen := 4
	jumping := false
	for {
		str := codes[pos]
		num := toInt(str)
		if (index == 0) {
			opcode := parseCode(num)
			if (opcode[3] == 3 || opcode[3] == 4) {
				instructionsLen = 2
			} else if (opcode[3] == 5 || opcode[3] == 6) {
				instructionsLen = 3
			} else {
				instructionsLen = 4
			}
		}
		if index % instructionsLen == 0 && num == 99 {
			return codes
		}
		if (index + 1) % instructionsLen == 0 {
			opcode := parseCode(buffer[0])
			result := 0
			// fmt.Println(index, opcode, buffer, num)
			if (opcode[3] == 1 || opcode[3] == 2) {
				v1 := toInt(getInstruction(opcode[2], buffer[1], &codes))
				v2 := toInt(getInstruction(opcode[1], buffer[2], &codes))
				if (opcode[3] == 1) {
					result = v1 + v2
				} else {
					result = v1 * v2
				}
				codes[num] = strconv.Itoa(result)
			} else if (opcode[3] == 5 || opcode[3] == 6) {
				v1 := toInt(getInstruction(opcode[2], buffer[1], &codes))
				v2 := toInt(getInstruction(opcode[1], num, &codes))
				if (opcode[3] == 5 && v1 != 0) || (opcode[3] == 6 && v1 == 0) {
					// fmt.Printf("JUMPING %s %s\n", v1, v2)
					pos = v2
					buffer = []int{}
					index = 0
					jumping = true
				}
			} else if (opcode[3] == 7 || opcode[3] == 8) {
				v1 := toInt(getInstruction(opcode[2], buffer[1], &codes))
				v2 := toInt(getInstruction(opcode[1], buffer[2], &codes))
				v3 := num
				if (opcode[3] == 7) {
					tmp := "0"
					if v1 < v2 {
						tmp = "1"
					}
					codes[v3] = tmp
				}
				if (opcode[3] == 8) {
					tmp := "0"
					if v1 == v2 {
						tmp = "1"
					}
					codes[v3] = tmp
				}
			} else if (opcode[3] == 3) {
				if (opcode[2] == 0) {
					codes[num] = input
				}
			} else if (opcode[3] == 4) {
				output := str
				if (opcode[2] == 0) {
					output = codes[num]
				}
				fmt.Printf("OUTPUT: %s\n", output)
			}
			buffer = []int{}
			index = 0
			// fmt.Println(codes)
		} else {
			buffer = append(buffer, num)
			index = index + 1
		}
		if !jumping {
			pos = pos + 1
		}
		jumping = false
	}
	return codes
}

func main() {
	content, _ := ioutil.ReadFile("5.txt")
	codes := strings.Split(string(content), ",")
	reader := bufio.NewReader(os.Stdin)
  fmt.Println("Simple Shell")
	fmt.Println("---------------------")
	for {
    fmt.Print("-> ")
    text, _ := reader.ReadString('\n')
    // convert CRLF to LF
    text = strings.Replace(text, "\n", "", -1)

    if text != "" {
			part5(codes, text)
			break
    }
  }
}
