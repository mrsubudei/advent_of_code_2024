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

	sum := 0
	m := make(map[string]int)

	for _, v := range sl {
		currAns := 0
		iter := 0
		ans := dfs(m, v, currAns, iter)

		sum += ans
	}

	fmt.Println("sum: ", sum)
}

func dfs(m map[string]int, num, currAns, iter int) int {
	if val, ok := m[getKey(num, iter)]; ok {
		return val
	}

	if iter == 75 {
		return 1
	}

	if num == 0 {
		ans := dfs(m, 1, currAns, iter+1)
		m[getKey(num, iter)] = ans

		return ans
	}

	str := strconv.Itoa(num)
	if len(str)%2 == 0 {
		num1, _ := strconv.Atoi(str[:len(str)/2])
		num2, _ := strconv.Atoi(str[len(str)/2:])

		ans1 := dfs(m, num1, currAns, iter+1)
		ans2 := dfs(m, num2, currAns, iter+1)

		m[getKey(num, iter)] = ans1 + ans2

		return ans1 + ans2
	}

	ans := dfs(m, num*2024, currAns, iter+1)
	m[getKey(num, iter)] = ans

	return ans
}

func getKey(num, iter int) string {
	return strconv.Itoa(num) + "-" + strconv.Itoa(iter)
}
