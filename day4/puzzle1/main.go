package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	matrix := make([][]string, 0, 10)
	first := true
	emptySl := []string{}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if first {
			emptySl = make([]string, len(line)+6)
			for i := 0; i < 3; i++ {
				matrix = append(matrix, emptySl)
			}

			first = false
		}

		matrix = append(matrix, getSl(line))
	}

	for i := 0; i < 3; i++ {
		matrix = append(matrix, emptySl)
	}

	count := 0

	for y := 3; y < len(matrix)-3; y++ {
		for x := 3; x < len(matrix[y])-3; x++ {
			if matrix[y][x] == "X" {
				if isRightFit(matrix, x, y) {
					count++
				}
				if isLeftFit(matrix, x, y) {
					count++
				}
				if isDownFit(matrix, x, y) {
					count++
				}
				if isUpFit(matrix, x, y) {
					count++
				}
				if isRightDownDiagFit(matrix, x, y) {
					count++
				}
				if isRightUpDiagFit(matrix, x, y) {
					count++
				}
				if isLeftDownDiagFit(matrix, x, y) {
					count++
				}
				if isLeftUpDiagFit(matrix, x, y) {
					count++
				}
			}
		}
	}

	fmt.Println("count: ", count)
}

func getSl(str string) []string {
	sl := make([]string, 0, len(str)+6)

	for i := 0; i < 3; i++ {
		sl = append(sl, "")
	}

	for _, v := range str {
		sl = append(sl, string(v))
	}

	for i := 0; i < 3; i++ {
		sl = append(sl, "")
	}

	return sl
}

func isRightFit(matrix [][]string, x, y int) bool {
	return matrix[y][x+1] == "M" &&
		matrix[y][x+2] == "A" &&
		matrix[y][x+3] == "S"
}

func isLeftFit(matrix [][]string, x, y int) bool {
	a := matrix[y][x-1] == "M"
	b := matrix[y][x-2] == "A"
	c := matrix[y][3] == "S"

	_ = a
	_ = b
	_ = c
	return matrix[y][x-1] == "M" &&
		matrix[y][x-2] == "A" &&
		matrix[y][x-3] == "S"
}

func isUpFit(matrix [][]string, x, y int) bool {
	return matrix[y-1][x] == "M" &&
		matrix[y-2][x] == "A" &&
		matrix[y-3][x] == "S"
}

func isDownFit(matrix [][]string, x, y int) bool {
	return matrix[y+1][x] == "M" &&
		matrix[y+2][x] == "A" &&
		matrix[y+3][x] == "S"
}

func isRightUpDiagFit(matrix [][]string, x, y int) bool {
	return matrix[y-1][x+1] == "M" &&
		matrix[y-2][x+2] == "A" &&
		matrix[y-3][x+3] == "S"
}

func isRightDownDiagFit(matrix [][]string, x, y int) bool {
	return matrix[y+1][x+1] == "M" &&
		matrix[y+2][x+2] == "A" &&
		matrix[y+3][x+3] == "S"
}

func isLeftUpDiagFit(matrix [][]string, x, y int) bool {
	return matrix[y-1][x-1] == "M" &&
		matrix[y-2][x-2] == "A" &&
		matrix[y-3][x-3] == "S"
}

func isLeftDownDiagFit(matrix [][]string, x, y int) bool {
	return matrix[y+1][x-1] == "M" &&
		matrix[y+2][x-2] == "A" &&
		matrix[y+3][x-3] == "S"
}
