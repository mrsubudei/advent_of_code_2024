package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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
	r, _ := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)

	sum := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		sl := r.FindAllString(line, -1)

		for _, val := range sl {
			sum += getNum(val)
		}
	}

	fmt.Println("sum: ", sum)
}

func getNum(str string) int {
	str = strings.Replace(str, "mul(", "", 1)
	str = strings.Replace(str, ")", "", 1)
	sl := strings.Split(str, ",")
	num1, _ := strconv.Atoi(sl[0])
	num2, _ := strconv.Atoi(sl[1])

	return num1 * num2
}
