package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

type Machine struct {
	DesiredState   []bool
	CurrentState   []bool
	Buttons        [][]int
	Joltage        []int
	DesiredJoltage []int
	MinPresses     int
}

func fs1(f *os.File) int {
	scanner := bufio.NewScanner(f)
	machines := []*Machine{}
	for scanner.Scan() {
		line := scanner.Text()
		machine := &Machine{
			MinPresses: 1 << 30,
		}
		machineInput := strings.Split(line, " ")
		for _, part := range machineInput {
			if strings.HasPrefix(part, "[") && strings.HasSuffix(part, "]") {
				initState := strings.Trim(part, "[]")
				for _, c := range initState {
					if c == '.' {
						machine.DesiredState = append(machine.DesiredState, false)
					} else if c == '#' {
						machine.DesiredState = append(machine.DesiredState, true)
					}
					machine.CurrentState = append(machine.CurrentState, false)
				}
			}
			if strings.HasPrefix(part, "(") && strings.HasSuffix(part, ")") {
				buttonsStr := strings.Trim(part, "()")
				buttonsParts := strings.Split(buttonsStr, ",")
				button := []int{}
				for _, b := range buttonsParts {
					val, err := strconv.Atoi(b)
					if err != nil {
						panic(err)
					}
					button = append(button, val)
				}
				machine.Buttons = append(machine.Buttons, button)
			}
		}
		machines = append(machines, machine)

	}

	ans := 0

	for _, machine := range machines {
		// Generate all possible button press combinations
		n := len(machine.Buttons)
		totalCombinations := 1 << n

		for i := 0; i < totalCombinations; i++ {
			// Reset current state
			presses := 0
			for j := range machine.CurrentState {
				machine.CurrentState[j] = false
			}

			// Apply button presses based on the current combination
			for j := 0; j < n; j++ {
				if (i & (1 << j)) != 0 {
					presses++
					for _, index := range machine.Buttons[j] {
						machine.CurrentState[index] = !machine.CurrentState[index]
					}
				}
			}

			// Check if current state matches desired state
			match := true
			for k := range machine.DesiredState {
				if machine.DesiredState[k] != machine.CurrentState[k] {
					match = false
					break
				}
			}
			if match {
				if presses < machine.MinPresses {
					machine.MinPresses = presses
				}
			}
		}
	}

	for _, machine := range machines {
		ans += machine.MinPresses
	}
	return ans
}

func fs2(f *os.File) int {
	scanner := bufio.NewScanner(f)
	ans := 0
	for scanner.Scan() {
		line := scanner.Text()
		machine := &Machine{}
		machineInput := strings.Split(line, " ")
		for _, part := range machineInput {
			if strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}") {
				initJoltage := strings.Trim(part, "{}")
				joltageParts := strings.Split(initJoltage, ",")
				for _, j := range joltageParts {
					val, err := strconv.Atoi(j)
					if err != nil {
						panic(err)
					}
					machine.DesiredJoltage = append(machine.DesiredJoltage, val)
				}
			}
			if strings.HasPrefix(part, "(") && strings.HasSuffix(part, ")") {
				buttonsStr := strings.Trim(part, "()")
				buttonsParts := strings.Split(buttonsStr, ",")
				button := []int{}
				for _, b := range buttonsParts {
					val, err := strconv.Atoi(b)
					if err != nil {
						panic(err)
					}
					button = append(button, val)
				}
				machine.Buttons = append(machine.Buttons, button)
			}
		}

		res := SolveILP(machine.Buttons, machine.DesiredJoltage)
		if res != -1 {
			ans += res
		}
	}
	return ans
}

// SolveILP finds the minimum sum of non-negative integers x satisfying Ax = b.
// A is implicitly defined by buttons: buttons[j] contains the indices of counters incremented by button j.
// b is target vector.
// Returns -1 if no solution.
func SolveILP(buttons [][]int, target []int) int {
	numButtons := len(buttons)
	numCounters := len(target)

	// Build Matrix A
	A := make([][]float64, numCounters)
	for i := range A {
		A[i] = make([]float64, numButtons)
	}
	for j, button := range buttons {
		for _, counterIdx := range button {
			A[counterIdx][j] = 1.0
		}
	}

	// Build RHS b
	b := make([]float64, numCounters)
	for i, val := range target {
		b[i] = float64(val)
	}

	// Gaussian Elimination
	pivotRow := 0
	pivotCols := make([]int, 0)
	colToPivotRow := make(map[int]int)

	for col := 0; col < numButtons && pivotRow < numCounters; col++ {
		// Find pivot
		sel := -1
		for row := pivotRow; row < numCounters; row++ {
			if math.Abs(A[row][col]) > 1e-9 {
				sel = row
				break
			}
		}

		if sel == -1 {
			continue
		}

		// Swap rows
		A[pivotRow], A[sel] = A[sel], A[pivotRow]
		b[pivotRow], b[sel] = b[sel], b[pivotRow]

		// Normalize pivot row
		pivotVal := A[pivotRow][col]
		for j := col; j < numButtons; j++ {
			A[pivotRow][j] /= pivotVal
		}
		b[pivotRow] /= pivotVal

		// Eliminate other rows
		for i := 0; i < numCounters; i++ {
			if i != pivotRow {
				factor := A[i][col]
				if math.Abs(factor) > 1e-9 {
					for j := col; j < numButtons; j++ {
						A[i][j] -= factor * A[pivotRow][j]
					}
					b[i] -= factor * b[pivotRow]
				}
			}
		}

		pivotCols = append(pivotCols, col)
		colToPivotRow[col] = pivotRow
		pivotRow++
	}

	// Check for inconsistency in remaining rows
	for i := pivotRow; i < numCounters; i++ {
		if math.Abs(b[i]) > 1e-9 {
			return -1 // Inconsistent
		}
	}

	// Identify free variables
	freeCols := []int{}
	isPivot := make(map[int]bool)
	for _, p := range pivotCols {
		isPivot[p] = true
	}
	for j := 0; j < numButtons; j++ {
		if !isPivot[j] {
			freeCols = append(freeCols, j)
		}
	}

	// Calculate bounds for free variables
	freeVarBounds := make([]int, len(freeCols))
	for i, fCol := range freeCols {
		minBound := 1 << 30
		affected := false
		for _, counterIdx := range buttons[fCol] {
			affected = true
			if target[counterIdx] < minBound {
				minBound = target[counterIdx]
			}
		}
		if !affected {
			minBound = 0
		}
		freeVarBounds[i] = minBound
	}

	minTotalPresses := -1

	// Recursion
	var solveRec func(idx int, currentFreeVals []int)
	solveRec = func(idx int, currentFreeVals []int) {
		if idx == len(freeCols) {
			// Calculate pivot variables
			currentPresses := 0
			for _, v := range currentFreeVals {
				currentPresses += v
			}

			// x_pivot = b_row - sum(A[row][free] * x_free)
			valid := true

			// Map from pivot col to value.
			// pivotVals := make(map[int]int) // Not explicitly needed, just sum

			for _, pCol := range pivotCols {
				row := colToPivotRow[pCol]
				val := b[row]
				for i, fCol := range freeCols {
					coeff := A[row][fCol]
					val -= coeff * float64(currentFreeVals[i])
				}

				// Check if integer
				rounded := math.Round(val)
				if math.Abs(val-rounded) > 1e-9 {
					valid = false
					break
				}
				intVal := int(rounded)
				if intVal < 0 {
					valid = false
					break
				}
				// pivotVals[pCol] = intVal
				currentPresses += intVal
			}

			if valid {
				if minTotalPresses == -1 || currentPresses < minTotalPresses {
					minTotalPresses = currentPresses
				}
			}
			return
		}

		limit := freeVarBounds[idx]
		for val := 0; val <= limit; val++ {
			solveRec(idx+1, append(currentFreeVals, val))
		}
	}

	solveRec(0, []int{})

	return minTotalPresses
}

