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

	sl1 := []int{}
	m := make(map[int]int)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		sl := strings.Split(line, "   ")
		num1, _ := strconv.Atoi(sl[0])
		num2, _ := strconv.Atoi(sl[1])

		sl1 = append(sl1, num1)
		m[num2]++
	}

	sum := 0

	for _, v := range sl1 {
		sum += v * m[v]
	}

	fmt.Println("sum: ", sum)
}
