package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	matrix = append(matrix, make([]string, len(matrix[0])+2))

	currX, currY := getCurrPos(matrix)
	direction := north
	exit := false
	for {
		if exit {
			break
		}

		switch direction {
		case north:
			if matrix[currY-1][currX] == "#" {
				direction++

				continue
			}

			currY--
			if currY == 0 {
				exit = true

				break
			}

			markVisited(matrix, currX, currY)
		case east:
			if matrix[currY][currX+1] == "#" {
				direction++

				continue
			}

			currX++
			if currX == len(matrix[0])-1 {
				exit = true

				break
			}

			markVisited(matrix, currX, currY)
		case south:
			if matrix[currY+1][currX] == "#" {
				direction++

				continue
			}

			currY++
			if currY == len(matrix)-1 {
				exit = true

				break
			}

			markVisited(matrix, currX, currY)
		case west:
			if matrix[currY][currX-1] == "#" {
				direction = 0

				continue
			}

			currX--
			if currX == 0 {
				exit = true

				break
			}

			markVisited(matrix, currX, currY)
		}
	}

	fmt.Println("count: ", countVisited(matrix))
}

func countVisited(matrix [][]string) int {
	count := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == "X" {
				count++
			}
		}
	}

	return count
}

func markVisited(matrix [][]string, x, y int) {
	matrix[y][x] = "X"
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
