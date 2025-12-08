package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Problem struct {
	numbers []*int
	op      rune
}

func fs1(f *os.File) int {
	answer := 0

	text, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	textArr := strings.Split(string(text), "\n")
	var problems []*Problem

	for _, line := range textArr {
		if line == "" {
			continue
		}
		numbers := strings.Split(line, " ")
		problemNumber := 0
		for _, numStr := range numbers {
			if numStr == "" {
				continue
			}
			if len(problems) <= problemNumber {
				problems = append(problems, &Problem{})
			}
			if numStr == "+" || numStr == "*" {
				problems[problemNumber].op = rune(numStr[0])
			} else {
				numInt, err := strconv.Atoi(numStr)
				if err != nil {
					panic(err)
				}
				problems[problemNumber].numbers = append(problems[problemNumber].numbers, &numInt)
			}
			problemNumber++
		}
	}

	for _, problem := range problems {
		switch problem.op {
		case '+':
			sum := 0
			for _, num := range problem.numbers {
				sum += *num
			}
			answer += sum
		case '*':
			prod := 1
			for _, num := range problem.numbers {
				prod *= *num
			}
			answer += prod
		default:
			panic(fmt.Sprintf("unknown operator: %c", problem.op))
		}

	}
	return answer
}

func fs2(f *os.File) int {
	answer := 0

	lines := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	// Filter out empty lines that some files may have at the end
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	// find max width
	maxWidth := 0
	for _, line := range lines {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}

	// create grid
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = make([]rune, maxWidth)
		for j := 0; j < maxWidth; j++ {
			grid[i][j] = ' '
		}
		for j, r := range line {
			grid[i][j] = r
		}
	}

	problems := [][][]rune{}
	currentProblem := [][]rune{}

	for j := maxWidth - 1; j >= 0; j-- {
		isColumnAllSpaces := true
		column := make([]rune, len(grid))
		for i := 0; i < len(grid); i++ {
			column[i] = grid[i][j]
			if grid[i][j] != ' ' {
				isColumnAllSpaces = false
			}
		}

		if !isColumnAllSpaces {
			currentProblem = append([][]rune{column}, currentProblem...)
		} else {
			if len(currentProblem) > 0 {
				problems = append(problems, currentProblem)
				currentProblem = [][]rune{}
			}
		}
	}
	if len(currentProblem) > 0 {
		problems = append(problems, currentProblem)
	}

	for _, p := range problems {
		var op rune
		numbers := []int{}

		for _, col := range p {
			numStr := ""
			for i, r := range col {
				if i == len(col)-1 { // last row
					if r == '+' || r == '*' {
						op = r
					}
				}
				if r >= '0' && r <= '9' {
					numStr += string(r)
				}
			}
			if numStr != "" {
				num, err := strconv.Atoi(numStr)
				if err != nil {
					panic(err)
				}
				numbers = append(numbers, num)
			}
		}

		switch op {
		case '+':
			sum := 0
			for _, num := range numbers {
				sum += num
			}
			answer += sum
		case '*':
			prod := 1
			for _, num := range numbers {
				prod *= num
			}
			answer += prod
		default:
			if len(numbers) > 0 {
				panic(fmt.Sprintf("unknown operator for problem with numbers %v", numbers))
			}
		}
	}

	return answer
}
