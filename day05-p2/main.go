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

	mRules := make(map[[2]int]struct{})
	sum := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.Contains(line, "|") {
			sl := strings.Split(line, "|")

			num1, _ := strconv.Atoi(sl[0])
			num2, _ := strconv.Atoi(sl[1])
			mRules[[2]int{num1, num2}] = struct{}{}
		} else if line == "" {

		} else {
			sl := strings.Split(line, ",")
			nums := make([]int, len(sl))

			for i, v := range sl {
				num, _ := strconv.Atoi(v)
				nums[i] = num
			}

			if isCorrect(mRules, nums) {
				continue
			}

			bubbleFix(mRules, nums)
			sum += getMiddle(nums)
		}
	}

	fmt.Println("sum: ", sum)
}

func bubbleFix(mRules map[[2]int]struct{}, nums []int) {
	for i := 0; i < len(nums); i++ {
		if isInPlace(mRules, nums, nums[i], i) {
			continue
		}

		curr := i
		i--
		for j := curr + 1; j < len(nums); j++ {
			nums[curr], nums[j] = nums[j], nums[curr]
			if isInPlace(mRules, nums, nums[j], j) {
				break
			}
			curr++
		}
	}
}

func isInPlace(mRules map[[2]int]struct{}, nums []int, curr, idx int) bool {
	for j, num := range nums {
		if idx == j {
			continue
		}

		if idx < j {
			if _, ok := mRules[[2]int{curr, num}]; !ok {
				return false
			}
		} else {
			if _, ok := mRules[[2]int{num, curr}]; !ok {
				return false
			}
		}
	}

	return true
}

func isCorrect(mRules map[[2]int]struct{}, nums []int) bool {
	for i, curr := range nums {
		for j, num := range nums {
			if i == j {
				continue
			}

			if i < j {
				if _, ok := mRules[[2]int{curr, num}]; !ok {
					return false
				}
			} else {
				if _, ok := mRules[[2]int{num, curr}]; !ok {
					return false
				}
			}
		}
	}

	return true
}

func getMiddle(sl []int) int {
	return sl[len(sl)/2]
}
