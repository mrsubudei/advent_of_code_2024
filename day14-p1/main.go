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
	width  = 101
	height = 103

	seconds = 100
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	m := make(map[int]int)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		sl := strings.Split(line, " ")

		slP := strings.Split(sl[0][2:], ",")
		slV := strings.Split(sl[1][2:], ",")
		pX, _ := strconv.Atoi(slP[0])
		pY, _ := strconv.Atoi(slP[1])
		vX, _ := strconv.Atoi(slV[0])
		vY, _ := strconv.Atoi(slV[1])

		afterPx, afterPy := getCoordinate(pX, pY, vX, vY)
		m[getQuadrant(afterPx, afterPy)]++
	}

	sum := 1
	for k, v := range m {
		if k == 0 {
			continue
		}

		sum *= v
	}

	fmt.Println("sum: ", sum)
}

func getCoordinate(pX, pY, vX, vY int) (int, int) {
	totalX := vX * seconds
	totalY := vY * seconds

	extraX := totalX % width
	extraY := totalY % height

	for extraX != 0 {
		if vX > 0 {
			pX += 1
			if pX == width {
				pX = 0
			}
		} else {
			pX -= 1
			if pX == -1 {
				pX = width - 1
			}
		}

		if vX > 0 {
			extraX--
		} else {
			extraX++
		}
	}

	for extraY != 0 {
		if vY > 0 {
			pY += 1
			if pY == height {
				pY = 0
			}
		} else {
			pY -= 1
			if pY == -1 {
				pY = height - 1
			}
		}

		if vY > 0 {
			extraY--
		} else {
			extraY++
		}
	}

	return pX, pY
}

func getQuadrant(pX, pY int) int {
	switch {
	case pX < width/2 && pY < height/2:
		return 1
	case pX > width/2 && pY < height/2:
		return 2
	case pX < width/2 && pY > height/2:
		return 3
	case pX > width/2 && pY > height/2:
		return 4
	}

	return 0
}
