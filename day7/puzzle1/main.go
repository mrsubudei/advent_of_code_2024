package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		sl := strings.Split(line, ": ")

		need := getNum(sl[0])

		slStr := strings.Split(sl[1], " ")
		nums := make([]int, len(slStr))

		for i, v := range slStr {
			nums[i] = getNum(v)
		}

		if check(nums, need) {
			count += need
		}
	}

	fmt.Println("count: ", count)
}

func countNums(str string) int {
	expr := ""
	prev := 0
	tmp := ""

	for _, v := range str {
		switch v {
		case '+', '*':
			if expr == "" {
				prev = getNum(tmp)
			} else {
				if expr == "+" {
					prev += getNum(tmp)
				} else {
					prev *= getNum(tmp)
				}
			}

			tmp = ""
			expr = string(v)
		default:
			tmp += string(v)
		}
	}

	if expr == "+" {
		return prev + getNum(tmp)
	}

	return prev * getNum(tmp)
}

func getNum(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	return num
}

func generateCombinations(nums []int, idx int, current string, result *bool, need int) {
	if idx == len(nums)-1 {
		if countNums(current) == need {
			*result = true
		}

		return
	}

	nextExpression := current + "+" + strconv.Itoa(nums[idx+1])
	generateCombinations(nums, idx+1, nextExpression, result, need)

	nextExpression = current + "*" + strconv.Itoa(nums[idx+1])
	generateCombinations(nums, idx+1, nextExpression, result, need)
}

func check(nums []int, need int) bool {
	result := false
	initial := strconv.Itoa(nums[0])
	generateCombinations(nums, 0, initial, &result, need)

	return result
}
