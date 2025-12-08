package main

import (
	"bufio"
	"os"
	"strconv"
)

func fs1(f *os.File) int {

	position := 50
	password := 0

	// Read in lines from f
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		direction := line[0]
		value, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		switch direction {
		case 'L':
			// turn left
			position -= value
			position = (position%100 + 100) % 100
		case 'R':
			// turn right
			position += value
			position %= 100
		default:
			panic("unknown direction")
		}

		if position == 0 {
			password += 1
		}

	}

	return password
}

func floordiv(a, b int) int {
	res := a / b
	if (a%b != 0) && (a*b < 0) {
		res--
	}
	return res
}

func fs2(f *os.File) int {
	position := 50
	password := 0

	// Read in lines from f
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		direction := line[0]
		value, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		crosses := 0
		if direction == 'R' {
			crosses = (position + value) / 100
			position = (position + value) % 100
		} else if direction == 'L' {
			crosses = floordiv(position-1, 100) - floordiv(position-value-1, 100)
			position = (position - value%100 + 100) % 100
		}
		password += crosses
	}

	return password
}
