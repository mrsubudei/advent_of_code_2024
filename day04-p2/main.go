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
			if matrix[y][x] == "A" {
				if ver1(matrix, x, y) {
					count++
				}
				if ver2(matrix, x, y) {
					count++
				}
				if ver3(matrix, x, y) {
					count++
				}
				if ver4(matrix, x, y) {
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

func ver1(matrix [][]string, x, y int) bool {
	return matrix[y-1][x-1] == "M" &&
		matrix[y+1][x-1] == "M" &&
		matrix[y-1][x+1] == "S" &&
		matrix[y+1][x+1] == "S"
}

func ver2(matrix [][]string, x, y int) bool {
	return matrix[y-1][x-1] == "M" &&
		matrix[y+1][x-1] == "S" &&
		matrix[y-1][x+1] == "M" &&
		matrix[y+1][x+1] == "S"
}

func ver3(matrix [][]string, x, y int) bool {
	return matrix[y-1][x-1] == "S" &&
		matrix[y+1][x-1] == "S" &&
		matrix[y-1][x+1] == "M" &&
		matrix[y+1][x+1] == "M"
}

func ver4(matrix [][]string, x, y int) bool {
	return matrix[y-1][x-1] == "S" &&
		matrix[y+1][x-1] == "M" &&
		matrix[y-1][x+1] == "S" &&
		matrix[y+1][x+1] == "M"
}
