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

	dotIdx := 0
	for i := len(sl) - 1; i >= 0; i-- {
		if sl[i] == "." {
			continue
		}

		found := false
		for j := dotIdx; j < len(sl); j++ {
			if j == i {
				break
			}

			if sl[j] == "." {
				sl[i], sl[j] = sl[j], sl[i]
				dotIdx = j
				found = true

				break
			}
		}

		if !found {
			break
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
