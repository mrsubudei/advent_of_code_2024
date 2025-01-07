package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Machine struct {
	buttonA [2]int
	buttonB [2]int
	prize   [2]int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	machines := []Machine{}

	currA := [2]int{}
	currB := [2]int{}

	rBut, err := regexp.Compile(`X\+\d*, Y\+\d*`)
	if err != nil {
		log.Fatal(err)
	}
	rPrize, err := regexp.Compile(`X\=\d*, Y\=\d*`)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		if strings.Contains(line, "Button A:") {
			currA = getNums(line, rBut, "+")

			continue
		}

		if strings.Contains(line, "Button B:") {
			currB = getNums(line, rBut, "+")

			continue
		}

		prize := getNums(line, rPrize, "=")

		machines = append(machines, Machine{
			buttonA: currA,
			buttonB: currB,
			prize:   prize,
		})
	}

	sum := 0

	for _, machine := range machines {
		aRepeates, bRepeats, ok := find(machine)

		if ok {
			sum += aRepeates*3 + bRepeats
		}
	}

	fmt.Println("sum: ", sum)
}

func find(machine Machine) (int, int, bool) {
	sum := math.MaxInt
	aRepeats := 0
	bRepeats := 0
	found := false

	for i := 1; i <= 100; i++ {
		for j := 1; j <= 100; j++ {
			butA := machine.buttonA
			butB := machine.buttonB
			prize := machine.prize

			if butA[0]*i+butB[0]*j == prize[0] &&
				butA[1]*i+butB[1]*j == prize[1] {

				if i*3+j < sum {
					sum = i*3 + j

					aRepeats = i
					bRepeats = j

					found = true
				}
			}
		}
	}

	return aRepeats, bRepeats, found
}

func getNums(line string, rBut *regexp.Regexp, sign string) [2]int {
	str := rBut.FindString(line)
	sl := strings.Split(str, ", ")

	slX := strings.Split(sl[0], sign)
	numX, _ := strconv.Atoi(slX[1])

	slY := strings.Split(sl[1], sign)
	numY, _ := strconv.Atoi(slY[1])

	return [2]int{numX, numY}
}
