package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func numHasRepeatingSequence(num int) bool {
	//	s := strconv.Itoa(num)
	//	n := len(s)
	//	for l := 1; l <= n/2; l++ {
	//		for i := 0; i <= n-2*l; i++ {
	//			if s[i:i+l] == s[i+l:i+2*l] {
	//				return true
	//			}
	//		}
	//	}
	//
	//	return false

	s := strconv.Itoa(num)
	n := len(s)

	// check if the first half is the same as the second half
	if n%2 != 0 {
		return false
	}
	half := n / 2
	if s[0:half] == s[half:n] {
		return true
	}
	return false
}

func numHasRepeatingSequencev2(num int) bool {
	// check if num is entirely made up of repeating digits e.g. 123123123
	s := strconv.Itoa(num)
	n := len(s)

	for l := 1; l <= n/2; l++ {
		if n%l != 0 {
			continue
		}
		pattern := s[0:l]
		matches := true
		for i := l; i < n; i += l {
			if s[i:i+l] != pattern {
				matches = false
				break
			}
		}
		if matches {
			return true
		}
	}

	return false
}

func fs1(f *os.File) int {

	sum := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		ranges := strings.Split(line, ",")
		for _, rangex := range ranges {
			if rangex == "" {
				continue
			}
			parts := strings.Split(rangex, "-")
			start, err := strconv.Atoi(parts[0])
			if err != nil {
				fmt.Println("Error converting string to int:", parts[0], err)
				panic(err)
			}
			end, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Error converting string to int:", parts[1], err)
				panic(err)
			}
			for i := start; i <= end; i++ {
				if numHasRepeatingSequence(i) {
					sum += i
				}
			}
		}

	}
	return sum
}

func fs2(f *os.File) int {

	sum := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		ranges := strings.Split(line, ",")
		for _, rangex := range ranges {
			if rangex == "" {
				continue
			}
			parts := strings.Split(rangex, "-")
			start, err := strconv.Atoi(parts[0])
			if err != nil {
				fmt.Println("Error converting string to int:", parts[0], err)
				panic(err)
			}
			end, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Error converting string to int:", parts[1], err)
				panic(err)
			}
			for i := start; i <= end; i++ {
				if numHasRepeatingSequencev2(i) {
					sum += i
				}
			}
		}

	}
	return sum
}
