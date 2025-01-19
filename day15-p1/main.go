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

	matrix := make([][]string, 0)

	scanner := bufio.NewScanner(file)

	isMatrix := true
	moves := ""
	currX := 0
	currY := 0
	y := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.Contains(line, "@") {
			currY = y

			for i, v := range line {
				if v == '@' {
					currX = i
				}
			}
		}

		if line == "" {
			isMatrix = false

			continue
		}

		if isMatrix {
			sl := make([]string, len(line))
			for i, v := range line {
				sl[i] = string(v)
			}
			matrix = append(matrix, sl)

			y++

			continue
		}

		moves += line
	}

	for _, v := range moves {
		switch v {
		case 'v':
			if matrix[currY+1][currX] == "#" {
				continue
			}

			if matrix[currY+1][currX] == "." {
				matrix[currY][currX] = "."
				matrix[currY+1][currX] = "@"
				currY++

				continue
			}

			for i := currY + 1; i < len(matrix); i++ {
				if matrix[i][currX] == "#" {
					break
				}

				if matrix[i][currX] == "." {
					matrix[i][currX] = "O"
					matrix[currY][currX] = "."
					matrix[currY+1][currX] = "@"
					currY++

					break
				}
			}
		case '^':
			if matrix[currY-1][currX] == "#" {
				continue
			}

			if matrix[currY-1][currX] == "." {
				matrix[currY][currX] = "."
				matrix[currY-1][currX] = "@"
				currY--

				continue
			}

			for i := currY - 1; i >= 0; i-- {
				if matrix[i][currX] == "#" {
					break
				}

				if matrix[i][currX] == "." {
					matrix[i][currX] = "O"
					matrix[currY][currX] = "."
					matrix[currY-1][currX] = "@"
					currY--

					break
				}
			}
		case '<':
			if matrix[currY][currX-1] == "#" {
				continue
			}

			if matrix[currY][currX-1] == "." {
				matrix[currY][currX] = "."
				matrix[currY][currX-1] = "@"
				currX--

				continue
			}

			for i := currX - 1; i >= 0; i-- {
				if matrix[currY][i] == "#" {
					break
				}

				if matrix[currY][i] == "." {
					matrix[currY][i] = "O"
					matrix[currY][currX] = "."
					matrix[currY][currX-1] = "@"
					currX--

					break
				}
			}
		case '>':
			if matrix[currY][currX+1] == "#" {
				continue
			}

			if matrix[currY][currX+1] == "." {
				matrix[currY][currX] = "."
				matrix[currY][currX+1] = "@"
				currX++

				continue
			}

			for i := currX + 1; i < len(matrix[currY]); i++ {
				if matrix[currY][i] == "#" {
					break
				}

				if matrix[currY][i] == "." {
					matrix[currY][i] = "O"
					matrix[currY][currX] = "."
					matrix[currY][currX+1] = "@"
					currX++

					break
				}
			}
		}
	}

	sum := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == "O" {
				sum += 100*i + j
			}
		}
	}

	fmt.Println("sum: ", sum)
}
