:ackage main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func maxDigit(s string) int {
	max := -1
	for _, r := range s {
		d, _ := strconv.Atoi(string(r))
		if d > max {
			max = d
		}
	}
	return max
}

func fs1(f *os.File) int {
	answer := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		maxJoltageForBank := 0

		if len(line) < 2 {
			continue
		}

		for i := 0; i < len(line)-1; i++ {
			d1, _ := strconv.Atoi(string(line[i]))
			d2 := maxDigit(line[i+1:])

			currentNum := d1*10 + d2
			if currentNum > maxJoltageForBank {
				maxJoltageForBank = currentNum
			}
		}
		answer += maxJoltageForBank
	}

	return answer
}



// findLargestNumber finds the largest number of length d that can be formed from s.
func findLargestNumber(s string, d int) string {
	k := len(s) - d
	if k < 0 {
		k = 0
	}

	stack := make([]rune, 0)

	for _, r := range s {
		for len(stack) > 0 && stack[len(stack)-1] < r && k > 0 {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, r)
	}

	if len(stack) > d {
		stack = stack[:d]
	}

	return string(stack)
}

func fs2(f *os.File) int {
	answer := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) < 12 {
			continue
		}

		// get highest 12 digits from line
		highest12Digits := findLargestNumber(line, 12)
		maxJoltageForBank, err := strconv.Atoi(highest12Digits)
		if err != nil {
			fmt.Println("Error converting highest 12 digits to int: ", err)
			continue
		}

		answer += maxJoltageForBank
	}

	return answer
}
