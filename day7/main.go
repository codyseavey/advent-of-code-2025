package main

import (
	"bufio"
	"os"
)

var splitCoords = []struct{ x, y int }{
	{1, 1},
	{1, -1},
}

func fs1(f *os.File) int {
	var grid [][]string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		var row []string
		for _, ch := range line {
			row = append(row, string(ch))
		}
		grid = append(grid, row)
	}

	splitCount := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			// if we hit a spliter (^) and there is a S above it, then split it and mark the two below as S
			if grid[i][j] == "^" && i > 0 && grid[i-1][j] == "S" {
				splitCount++
				// mark the two split coordinates below
				for _, coord := range splitCoords {
					newX := i + coord.x
					newY := j + coord.y
					if newX >= 0 && newX < len(grid) && newY >= 0 && newY < len(grid[i]) {
						grid[newX][newY] = "S"
					}
				}
				// if we hit a S propgate it downward unless the space is a ^
			} else if grid[i][j] == "S" && (i+1 < len(grid)) && grid[i+1][j] != "^" {
				// propagate S downwards
				if i+1 < len(grid) {
					grid[i+1][j] = "S"
				}
			}
		}
	}

	return splitCount
}

func fs2(f *os.File) int {
	scanner := bufio.NewScanner(f)
	grid := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		var row []string
		for _, ch := range line {
			row = append(row, string(ch))
		}
		grid = append(grid, row)
	}

	if len(grid) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	dp := make([][]int, rows)
	for i := range dp {
		dp[i] = make([]int, cols)
	}

	for r := rows - 1; r >= 0; r-- {
		for c := 0; c < cols; c++ {
			if r+1 >= rows {
				dp[r][c] = 1
				continue
			}

			next_r, next_c := r+1, c
			if grid[next_r][next_c] == "^" {
				ways1, ways2 := 0, 0

				coord1 := splitCoords[0]
				branch1_r, branch1_c := next_r+coord1.x, next_c+coord1.y
				if branch1_r >= rows || branch1_c < 0 || branch1_c >= cols {
					ways1 = 1
				} else {
					ways1 = dp[branch1_r][branch1_c]
				}

				coord2 := splitCoords[1]
				branch2_r, branch2_c := next_r+coord2.x, next_c+coord2.y
				if branch2_r >= rows || branch2_c < 0 || branch2_c >= cols {
					ways2 = 1
				} else {
					ways2 = dp[branch2_r][branch2_c]
				}
				dp[r][c] = ways1 + ways2
			} else {
				dp[r][c] = dp[next_r][next_c]
			}
		}
	}

	startX, startY := -1, -1
	for j := 0; j < cols; j++ {
		if grid[0][j] == "S" {
			startX = 0
			startY = j
			break
		}
	}

	if startX == -1 {
		return 0
	}

	return dp[startX][startY]
}

