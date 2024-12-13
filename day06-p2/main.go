package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	north int = iota
	east
	south
	west
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	matrix := [][]string{}
	first := true
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if first {
			matrix = append(matrix, make([]string, len(line)+2))

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

	matrix = append(matrix, make([]string, len(matrix[0])))

	count := 0
	for i := 1; i < len(matrix)-1; i++ {
		for j := 1; j < len(matrix[i])-1; j++ {
			copied := copyMatrix(matrix)
			copied[i][j] = "#"

			currX, currY := getCurrPos(matrix)
			direction := north
			copied[currY][currX] = "."

			if !canFindExist(copied, direction, currX, currY) {
				count++
			}
		}
	}

	fmt.Println("count: ", count)
}

func copyMatrix(matrix [][]string) [][]string {
	result := make([][]string, 0, len(matrix))

	for i := 0; i < len(matrix); i++ {
		sl := make([]string, len(matrix[i]))

		copy(sl, matrix[i])
		result = append(result, sl)
	}

	return result
}

func canFindExist(matrix [][]string, direction, currX, currY int) bool {
	for {
		switch direction {
		case north:
			if matrix[currY-1][currX] == "#" {
				direction++

				continue
			}

			currY--
			if currY == 0 {
				return true
			} else if matrix[currY][currX] == "5" {
				return false
			} else {
				increaseVisited(matrix, currX, currY)
			}
		case east:
			if matrix[currY][currX+1] == "#" {
				direction++

				continue
			}

			currX++
			if currX == len(matrix[0])-1 {
				return true
			} else if matrix[currY][currX] == "5" {
				return false
			} else {
				increaseVisited(matrix, currX, currY)
			}
		case south:
			if matrix[currY+1][currX] == "#" {
				direction++

				continue
			}

			currY++
			if currY == len(matrix)-1 {
				return true
			} else if matrix[currY][currX] == "5" {
				return false
			} else {
				increaseVisited(matrix, currX, currY)
			}
		case west:
			if matrix[currY][currX-1] == "#" {
				direction = 0

				continue
			}

			currX--
			if currX == 0 {
				return true
			} else if matrix[currY][currX] == "5" {
				return false
			} else {
				increaseVisited(matrix, currX, currY)
			}
		}
	}
}

func increaseVisited(matrix [][]string, x, y int) {
	currStr := matrix[y][x]

	if matrix[y][x] == "." {
		matrix[y][x] = "0"

		return
	}

	num, _ := strconv.Atoi(currStr)
	num++

	matrix[y][x] = strconv.Itoa(num)
}

func getCurrPos(matrix [][]string) (x int, y int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == "^" {
				return j, i
			}
		}
	}

	panic("wrong data")
}
