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

	matrix := [][]string{}

	scanner := bufio.NewScanner(file)
	first := true
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if first {
			sl := make([]string, len(line)+2)
			matrix = append(matrix, sl)

			first = false
		}

		sl := make([]string, 0, len(line)+2)
		sl = append(sl, "")
		for _, v := range line {
			sl = append(sl, string(v))
		}
		sl = append(sl, "")

		matrix = append(matrix, sl)
	}

	matrix = append(matrix, matrix[0])

	count := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == "0" {

				count += bfs(matrix, j, i)
			}
		}
	}

	fmt.Println("count: ", count)
}

type Position struct {
	x       int
	y       int
	currNum int
}

func getEmptyMatrix(matrix [][]string) [][]string {
	cp := make([][]string, 0, len(matrix))

	for _, v := range matrix {
		tmp := make([]string, len(v))
		copy(tmp, v)

		cp = append(cp, tmp)
	}

	return cp
}

func bfs(matrix [][]string, x, y int) int {
	sl := []Position{
		{
			x: x,
			y: y,
		},
	}

	empty := getEmptyMatrix(matrix)

	for len(sl) != 0 {
		curr := sl[0]
		sl = sl[1:]

		if curr.currNum == 9 {
			empty[curr.y][curr.x] = "X"

			continue
		}

		nextSteps := getNextSteps(matrix, curr)
		sl = append(sl, nextSteps...)
	}

	return countTrailHeads(empty)
}

func countTrailHeads(matrix [][]string) int {
	res := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == "X" {
				res++
			}
		}
	}

	return res
}

func getNextSteps(matrix [][]string, pos Position) []Position {
	result := []Position{}

	nextNum := pos.currNum + 1

	if matrix[pos.y-1][pos.x] == strconv.Itoa(nextNum) {
		result = append(result, Position{
			y:       pos.y - 1,
			x:       pos.x,
			currNum: nextNum,
		})
	}

	if matrix[pos.y+1][pos.x] == strconv.Itoa(nextNum) {
		result = append(result, Position{
			y:       pos.y + 1,
			x:       pos.x,
			currNum: nextNum,
		})
	}

	if matrix[pos.y][pos.x-1] == strconv.Itoa(nextNum) {
		result = append(result, Position{
			y:       pos.y,
			x:       pos.x - 1,
			currNum: nextNum,
		})
	}

	if matrix[pos.y][pos.x+1] == strconv.Itoa(nextNum) {
		result = append(result, Position{
			y:       pos.y,
			x:       pos.x + 1,
			currNum: nextNum,
		})
	}

	return result
}
