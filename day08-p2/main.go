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

	count := 0
	matrix := [][]string{}
	m := make(map[string][][2]int)
	empty := [][]string{}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		sl := make([]string, len(line))
		emptySl := make([]string, len(line))
		empty = append(empty, emptySl)
		for i, v := range line {
			sl[i] = string(v)
		}

		matrix = append(matrix, sl)
	}

	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[y]); x++ {
			if matrix[y][x] != "." {
				m[matrix[y][x]] = append(m[matrix[y][x]], [2]int{y, x})
			}
		}
	}

	for _, val := range m {
		for i := 0; i < len(val); i++ {
			for j := i + 1; j < len(val); j++ {
				handleCoords(matrix, empty, val[i], val[j])
			}
		}

		for _, v := range val {
			putAntinode(empty, v)
		}
	}

	for i := 0; i < len(empty); i++ {
		for j := 0; j < len(empty[i]); j++ {
			if empty[i][j] == "#" {
				count++
			}
		}
	}

	fmt.Println("count: ", count)
}

func putAntinode(matrix [][]string, coord [2]int) {
	matrix[coord[0]][coord[1]] = "#"
}

func handleCoords(matrix, empty [][]string, arr1, arr2 [2]int) {
	absY := abs(arr1[0], arr2[0])
	absX := abs(arr1[1], arr2[1])

	currY1 := arr1[0]
	currX1 := arr1[1]

	currY2 := arr2[0]
	currX2 := arr2[1]

	if arr1[0] < arr2[0] && arr1[1] < arr2[1] {
		currY1 -= absY
		currX1 -= absX
		for currY1 >= 0 && currX1 >= 0 {
			putAntinode(empty, [2]int{currY1, currX1})

			currY1 -= absY
			currX1 -= absX
		}

		currY2 += absY
		currX2 += absX
		for currY2 < len(matrix) && currX2 < len(matrix[0]) {
			putAntinode(empty, [2]int{currY2, currX2})

			currY2 += absY
			currX2 += absX
		}
	}

	if arr1[0] < arr2[0] && arr1[1] >= arr2[1] {
		currY1 -= absY
		currX1 += absX
		for currY1 >= 0 && currX1 < len(matrix[0]) {
			putAntinode(empty, [2]int{currY1, currX1})

			currY1 -= absY
			currX1 += absX
		}

		currY2 += absY
		currX2 -= absX
		for currY2 < len(matrix) && currX2 >= 0 {
			putAntinode(empty, [2]int{currY2, currX2})

			currY2 += absY
			currX2 -= absX
		}
	}

	if arr1[0] >= arr2[0] && arr1[1] < arr2[1] {
		currY1 += absY
		currX1 -= absX
		for currY1 < len(matrix) && currX1 >= 0 {
			putAntinode(empty, [2]int{currY1, currX1})

			currY1 += absY
			currX1 -= absX
		}

		currY2 -= absY
		currX2 += absX
		for currY2 >= 0 && currX2 < len(matrix[0]) {
			putAntinode(empty, [2]int{currY2, currX2})

			currY2 -= absY
			currX2 += absX
		}
	}

	if arr1[0] >= arr2[0] && arr1[1] >= arr2[1] {
		currY1 += absY
		currX1 += absX
		for currY1 < len(matrix) && currX1 < len(matrix[0]) {
			putAntinode(empty, [2]int{currY1, currX1})

			currY1 += absY
			currX1 += absX
		}

		currY2 -= absY
		currX2 -= absX
		for currY2 >= 0 && currX2 >= 0 {
			putAntinode(empty, [2]int{currY2, currX2})

			currY2 -= absY
			currX2 -= absX
		}
	}
}

func abs(n1, n2 int) int {
	if n2 > n1 {
		return n2 - n1
	}

	return n1 - n2
}
