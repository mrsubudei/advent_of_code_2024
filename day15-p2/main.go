package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Data struct {
	y    int
	vals []Val
}

type Val struct {
	x   int
	val string
}

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

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			isMatrix = false

			continue
		}

		if isMatrix {
			sl := []string{}
			for _, v := range line {
				switch v {
				case '#':
					sl = append(sl, "#")
					sl = append(sl, "#")
				case 'O':
					sl = append(sl, "[")
					sl = append(sl, "]")
				case '.':
					sl = append(sl, ".")
					sl = append(sl, ".")
				case '@':
					sl = append(sl, "@")
					sl = append(sl, ".")
				}
			}
			matrix = append(matrix, sl)

			continue
		}

		moves += line
	}

	currX := 0
	currY := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == "@" {
				currY = i
				currX = j

				break
			}
		}
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

			sl := []Data{}
			first := true
			for {
				currData := Data{}

				if first {
					currData.y = currY
					currData.vals = []Val{
						{
							x:   currX,
							val: "@",
						},
					}

					first = false
				} else {
					currData = sl[len(sl)-1]
				}

				if nextRowHasHash(matrix, currData, true) {
					break
				}

				if rowAllDots(matrix, currData, true) {
					moveDown(matrix, sl, currY, currX)

					currY++
					break
				}

				nextData := getNextData(matrix, currData, true)
				sl = append(sl, nextData)
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

			sl := []Data{}
			first := true
			for {
				currData := Data{}

				if first {
					currData.y = currY
					currData.vals = []Val{
						{
							x:   currX,
							val: "@",
						},
					}

					first = false
				} else {
					currData = sl[len(sl)-1]
				}

				if nextRowHasHash(matrix, currData, false) {
					break
				}

				if rowAllDots(matrix, currData, false) {
					moveUp(matrix, sl, currY, currX)

					currY--
					break
				}

				nextData := getNextData(matrix, currData, false)
				sl = append(sl, nextData)
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
					for j := i; j < currX; j++ {
						matrix[currY][j] = matrix[currY][j+1]
						matrix[currY][j+1] = "."
					}

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
					for j := i; j > currX; j-- {
						matrix[currY][j] = matrix[currY][j-1]
						matrix[currY][j-1] = "."
					}

					currX++

					break
				}
			}
		}
	}

	sum := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == "[" {
				sum += 100*i + j
			}
		}
	}

	fmt.Println("sum: ", sum)
}

func moveUp(matrix [][]string, sl []Data, startY, startX int) {
	matrix[startY][startX] = "."

	for _, data := range sl {
		for _, v := range data.vals {
			matrix[data.y][v.x] = "."
		}
	}

	matrix[startY-1][startX] = "@"

	for _, data := range sl {
		for _, v := range data.vals {
			matrix[data.y-1][v.x] = v.val
		}
	}
}

func moveDown(matrix [][]string, sl []Data, startY, startX int) {
	matrix[startY][startX] = "."

	for _, data := range sl {
		for _, v := range data.vals {
			matrix[data.y][v.x] = "."
		}
	}

	matrix[startY+1][startX] = "@"

	for _, data := range sl {
		for _, v := range data.vals {
			matrix[data.y+1][v.x] = v.val
		}
	}
}

func getNextData(matrix [][]string, data Data, isDown bool) Data {
	nextY := data.y - 1
	if isDown {
		nextY = data.y + 1
	}

	result := Data{
		y: nextY,
	}

	for _, v := range data.vals {
		if matrix[nextY][v.x] != matrix[data.y][v.x] && matrix[nextY][v.x] == "]" {
			result.vals = append(result.vals, Val{
				x:   v.x - 1,
				val: "[",
			})
		}

		if matrix[nextY][v.x] == "[" || matrix[nextY][v.x] == "]" {
			result.vals = append(result.vals, Val{
				x:   v.x,
				val: matrix[nextY][v.x],
			})
		}

		if matrix[nextY][v.x] != matrix[data.y][v.x] && matrix[nextY][v.x] == "[" {
			result.vals = append(result.vals, Val{
				x:   v.x + 1,
				val: "]",
			})
		}
	}

	return result
}

func nextRowHasHash(matrix [][]string, data Data, isDown bool) bool {
	nextY := data.y - 1
	if isDown {
		nextY = data.y + 1
	}

	for _, val := range data.vals {
		if matrix[nextY][val.x] == "#" {
			return true
		}
	}

	return false
}

func rowAllDots(matrix [][]string, data Data, isDown bool) bool {
	nextY := data.y - 1
	if isDown {
		nextY = data.y + 1
	}

	for _, val := range data.vals {
		if matrix[nextY][val.x] != "." {
			return false
		}
	}

	return true
}
