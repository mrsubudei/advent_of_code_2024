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
	empty := [][]string{}
	m := make(map[string][][2]int)

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
				coor1, coor2 := getAntinodeCoords(val[i], val[j])
				putAntinode(empty, coor1, coor2)
			}
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

func putAntinode(matrix [][]string, coord1, coord2 [2]int) {
	if coord1[0] >= 0 && coord1[0] < len(matrix) &&
		coord1[1] >= 0 && coord1[1] < len(matrix[0]) {
		matrix[coord1[0]][coord1[1]] = "#"
	}

	if coord2[0] >= 0 && coord2[0] < len(matrix) &&
		coord2[1] >= 0 && coord2[1] < len(matrix[0]) {
		matrix[coord2[0]][coord2[1]] = "#"
	}
}

func getAntinodeCoords(arr1, arr2 [2]int) ([2]int, [2]int) {
	absY := abs(arr1[0], arr2[0])
	absX := abs(arr1[1], arr2[1])

	if arr1[0] < arr2[0] && arr1[1] < arr2[1] {
		return [2]int{arr1[0] - absY, arr1[1] - absX}, [2]int{arr2[0] + absY, arr2[1] + absX}
	}

	if arr1[0] < arr2[0] && arr1[1] >= arr2[1] {
		return [2]int{arr1[0] - absY, arr1[1] + absX}, [2]int{arr2[0] + absY, arr2[1] - absX}
	}

	if arr1[0] >= arr2[0] && arr1[1] < arr2[1] {
		return [2]int{arr1[0] + absY, arr1[1] - absX}, [2]int{arr2[0] - absY, arr2[1] + absX}
	}

	if arr1[0] >= arr2[0] && arr1[1] >= arr2[1] {
		return [2]int{arr1[0] + absY, arr1[1] + absX}, [2]int{arr2[0] - absY, arr2[1] - absX}
	}

	panic(fmt.Sprintf("arr1: %v, arr2: %v\n", arr1, arr2))
}

func abs(n1, n2 int) int {
	if n2 > n1 {
		return n2 - n1
	}

	return n1 - n2
}
