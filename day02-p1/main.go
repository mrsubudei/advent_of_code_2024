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

		sl := strings.Split(line, " ")
		nums := make([]int, len(sl))

		for i, v := range sl {
			num, _ := strconv.Atoi(v)
			nums[i] = num
		}

		if isSafe(nums) {
			count++
		}
	}

	fmt.Println("count: ", count)
}

func isSafe(sl []int) bool {
	dir := ""
	for i := 0; i < len(sl)-1; i++ {
		if dir == "" {
			if sl[i] == sl[i+1] || abs(sl[i], sl[i+1]) > 3 {
				return false
			}

			if sl[i] < sl[i+1] {
				dir = "inc"
			} else if sl[i] > sl[i+1] {
				dir = "decr"
			}
		} else {
			switch dir {
			case "inc":
				if sl[i] >= sl[i+1] || abs(sl[i], sl[i+1]) > 3 {
					return false
				}
			case "decr":
				if sl[i] <= sl[i+1] || abs(sl[i], sl[i+1]) > 3 {
					return false
				}
			}
		}
	}

	return true
}

func abs(n1, n2 int) int {
	if n2 > n1 {
		return n2 - n1
	}

	return n1 - n2
}
