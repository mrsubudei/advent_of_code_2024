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

	sl := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		strSl := strings.Split(line, " ")

		for _, v := range strSl {
			num, _ := strconv.Atoi(v)

			sl = append(sl, num)
		}
	}

	for i := 0; i < 25; i++ {
		tmp := []int{}
		for j := 0; j < len(sl); j++ {
			ans := process(sl[j])
			sl[j] = ans[0]

			if len(ans) == 2 {
				tmp = append(tmp, ans[1])
			}
		}

		sl = append(sl, tmp...)
	}

	fmt.Println("len: ", len(sl))
}

func process(num int) []int {
	if num == 0 {
		return []int{1}
	}

	str := strconv.Itoa(num)

	if len(str)%2 == 0 {
		num1, _ := strconv.Atoi(str[:len(str)/2])
		num2, _ := strconv.Atoi(str[len(str)/2:])

		return []int{num1, num2}
	}

	return []int{num * 2024}
}
