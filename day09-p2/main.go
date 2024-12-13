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
	sl := []string{}
	isBlock := true
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		id := 0
		for _, v := range line {
			times, _ := strconv.Atoi(string(v))

			symb := "."
			if isBlock {
				symb = strconv.Itoa(id)
			}
			for i := 0; i < times; i++ {
				sl = append(sl, symb)
			}

			if isBlock {
				isBlock = false
				id++
			} else {
				isBlock = true
			}
		}
	}

	for i := len(sl) - 1; i >= 0; i-- {
		if sl[i] == "." {
			continue
		}

		curr := sl[i]
		count := 1
		for j := i - 1; j >= 0; j-- {
			if sl[j] != curr {
				break
			}
			count++
		}

		replaced := false
		for fromDot := 0; fromDot <= i-count; fromDot++ {
			if sl[fromDot] != "." {
				continue
			}

			have := 1
			for k := fromDot + 1; k <= i-count; k++ {
				if sl[k] == "." {
					have++
				} else {
					break
				}
			}

			if have >= count {
				replaced = true

				t := i
				for k := 0; k < count; k++ {
					sl[fromDot], sl[t] = sl[t], sl[fromDot]
					fromDot++
					t--
				}

				i -= count - 1

				break
			}
		}

		if !replaced {
			i -= count - 1
		}
	}

	sum := 0
	mul := 0
	for _, v := range sl {
		num, _ := strconv.Atoi(v)
		sum += mul * num

		mul++
	}

	fmt.Println("sum: ", sum)
}
