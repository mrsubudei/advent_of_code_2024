package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

	sl1 := []int{}
	sl2 := []int{}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		sl := strings.Split(line, "   ")
		num1, _ := strconv.Atoi(sl[0])
		num2, _ := strconv.Atoi(sl[1])

		sl1 = append(sl1, num1)
		sl2 = append(sl2, num2)
	}

	sort.Ints(sl1)
	sort.Ints(sl2)

	sum := 0

	for i := range sl1 {
		sum += abs(sl1[i], sl2[i])
	}

	fmt.Println("sum: ", sum)
}

func abs(n1, n2 int) int {
	if n2 > n1 {
		return n2 - n1
	}

	return n1 - n2
}
