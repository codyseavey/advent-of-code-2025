package main

import (
	"bufio"
	"os"
)

var directions = [8][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func fs1(f *os.File) int {
	scanner := bufio.NewScanner(f)

	answer := 0
	var grid [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	// check grid for adjacent positions
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			// i is row, j is column

			filledPositions := 0
			if grid[i][j] != '@' {
				continue
			}
			for _, dir := range directions {
				ni := i + dir[0]
				nj := j + dir[1]

				if ni >= 0 && ni < len(grid) && nj >= 0 && nj < len(grid[i]) {
					if grid[ni][nj] == '@' {
						filledPositions++
					}

				}
			}
			if filledPositions < 4 {
				answer++
			}
		}
	}

	return answer
}

func fs2(f *os.File) int {
	scanner := bufio.NewScanner(f)

	answer := 0
	var grid [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	for {
		// check grid for adjacent positions
		positionsFilledIteration := 0
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				// i is row, j is column

				filledPositions := 0
				if grid[i][j] != '@' {
					continue
				}
				for _, dir := range directions {
					ni := i + dir[0]
					nj := j + dir[1]

					if ni >= 0 && ni < len(grid) && nj >= 0 && nj < len(grid[i]) {
						if grid[ni][nj] == '@' {
							filledPositions++
						}

					}
				}
				if filledPositions < 4 {
					positionsFilledIteration++
					grid[i][j] = '.'
				}
			}
		}
		if positionsFilledIteration == 0 {
			break
		}
		answer += positionsFilledIteration
	}

	return answer
}
