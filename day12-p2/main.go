package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
	sl := []Data{pos}
	count := 1

	mSides := make(map[string][]int)

	for len(sl) != 0 {
		currPos := sl[0]
		sl = sl[1:]

		matrix[currPos.y][currPos.x] = currPos.currLetter + "+"

		addIfSame(matrix, &sl, currPos.y-1, currPos.x, currPos.currLetter, &count)
		addIfSame(matrix, &sl, currPos.y+1, currPos.x, currPos.currLetter, &count)
		addIfSame(matrix, &sl, currPos.y, currPos.x-1, currPos.currLetter, &count)
		addIfSame(matrix, &sl, currPos.y, currPos.x+1, currPos.currLetter, &count)

		if countIfSutable(matrix, currPos.y-1, currPos.x, currPos.currLetter) {
			key := strconv.Itoa(currPos.y) + "up"
			mSides[key] = append(mSides[key], currPos.x)
		}
		if countIfSutable(matrix, currPos.y+1, currPos.x, currPos.currLetter) {
			key := strconv.Itoa(currPos.y) + "down"
			mSides[key] = append(mSides[key], currPos.x)
		}
		if countIfSutable(matrix, currPos.y, currPos.x-1, currPos.currLetter) {
			key := strconv.Itoa(currPos.x) + "left"
			mSides[key] = append(mSides[key], currPos.y)
		}
		if countIfSutable(matrix, currPos.y, currPos.x+1, currPos.currLetter) {
			key := strconv.Itoa(currPos.x) + "right"
			mSides[key] = append(mSides[key], currPos.y)
		}
	}

	sides := 0
	for _, v := range mSides {
		sides += countSides(v)
	}

	return count * sides
}

func countSides(sl []int) int {
	sort.Ints(sl)

	count := 1
	prev := sl[0]

	for _, v := range sl[1:] {
		if abs(prev, v) > 1 {
			count++
		}

		prev = v
	}

	return count
}

func abs(n1, n2 int) int {
	if n1 > n2 {
		return n1 - n2
	}

	return n2 - n1
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

func countIfSutable(matrix [][]string, y, x int, currLetter string) bool {
	if matrix[y][x] == currLetter || matrix[y][x] == currLetter+"+" {

		return false
	}

	return true
}
