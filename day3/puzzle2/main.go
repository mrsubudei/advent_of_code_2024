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
	rDo, _ := regexp.Compile(`do\(\)`)
	rDont, _ := regexp.Compile(`don't\(\)`)

	sum := 0
	for scanner.Scan() {
		fmt.Println("rtas")
		line := strings.TrimSpace(scanner.Text())
		sl := r.FindAllString(line, -1)
		indexes := r.FindAllStringIndex(line, -1)

		indexesDo := rDo.FindAllStringIndex(line, -1)
		indexesDont := rDont.FindAllStringIndex(line, -1)

		enabled := true
		for i, val := range sl {
			mulIdx := indexes[i][0]

			switch enabled {
			case true:
				lastDoIdx, ok2 := getLastIdx(indexesDo)
				if ok2 && mulIdx > lastDoIdx {
					indexesDo = indexesDo[1:]
				}

				lastDontIdx, ok := getLastIdx(indexesDont)

				if ok && mulIdx > lastDontIdx {
					enabled = false
					indexesDont = indexesDont[1:]
				} else {
					sum += getNum(val)
				}
			case false:
				lastDontIdx, ok := getLastIdx(indexesDont)
				if ok && mulIdx > lastDontIdx {
					indexesDont = indexesDont[1:]
				}

				lastDoIdx, ok2 := getLastIdx(indexesDo)
				if ok2 && mulIdx > lastDoIdx {
					enabled = true
					sum += getNum(val)
				}
			}
		}
	}

	fmt.Println("sum: ", sum)
}

func getLastIdx(sl [][]int) (int, bool) {
	if len(sl) == 0 {
		return 0, false
	}

	cp := sl[0][0]

	return cp, true
}

func getNum(str string) int {
	str = strings.Replace(str, "mul(", "", 1)
	str = strings.Replace(str, ")", "", 1)
	sl := strings.Split(str, ",")
	num1, _ := strconv.Atoi(sl[0])
	num2, _ := strconv.Atoi(sl[1])

	return num1 * num2
}
