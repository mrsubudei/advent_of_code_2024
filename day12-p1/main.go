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

	matrix := [][]string{}
	first := true

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if first {
			sl := make([]string, len(line)+2)
			matrix = append(matrix, sl)

			first = false
		}

		sl := make([]string, len(line)+2)
		idx := 1
		for _, v := range line {
			sl[idx] = string(v)
			idx++
		}

		matrix = append(matrix, sl)
	}
	matrix = append(matrix, matrix[0])

	m := getMap(matrix)

	sum := 0
	for key := range m {
		if key == "" {
			continue
		}

		sum += process(matrix, key)
	}

	fmt.Println("sum: ", sum)
}

func getMap(matrix [][]string) map[string]struct{} {
	m := make(map[string]struct{})

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			m[matrix[i][j]] = struct{}{}
		}
	}

	return m
}

type Data struct {
	x, y       int
	currLetter string
}

func process(matrix [][]string, currLetter string) int {
	ans := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == currLetter {
				ans += bfs(matrix, Data{
					y:          i,
					x:          j,
					currLetter: currLetter,
				})
			}
		}
	}

	return ans
}

func bfs(matrix [][]string, pos Data) int {
	edges := 0
	sl := []Data{pos}
	count := 1

	for len(sl) != 0 {
		currPos := sl[0]
		sl = sl[1:]

		matrix[currPos.y][currPos.x] = currPos.currLetter + "+"

		addIfSame(matrix, &sl, currPos.y-1, currPos.x, currPos.currLetter, &count)
		addIfSame(matrix, &sl, currPos.y+1, currPos.x, currPos.currLetter, &count)
		addIfSame(matrix, &sl, currPos.y, currPos.x-1, currPos.currLetter, &count)
		addIfSame(matrix, &sl, currPos.y, currPos.x+1, currPos.currLetter, &count)

		edges += countIfSutable(matrix, currPos.y-1, currPos.x, currPos.currLetter)
		edges += countIfSutable(matrix, currPos.y+1, currPos.x, currPos.currLetter)
		edges += countIfSutable(matrix, currPos.y, currPos.x-1, currPos.currLetter)
		edges += countIfSutable(matrix, currPos.y, currPos.x+1, currPos.currLetter)
	}

	return edges * count
}

func addIfSame(matrix [][]string, sl *[]Data, y, x int, currLetter string, count *int) {
	if matrix[y][x] == currLetter {
		matrix[y][x] = currLetter + "+"

		*sl = append(*sl, Data{
			y:          y,
			x:          x,
			currLetter: currLetter,
		})

		*count++
	}
}

func countIfSutable(matrix [][]string, y, x int, currLetter string) int {
	if matrix[y][x] == currLetter || matrix[y][x] == currLetter+"+" {

		return 0
	}

	return 1
}
