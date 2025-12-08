package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func fs1(f *os.File) int {
	answer := 0

	text, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	textArr := strings.Split(string(text), "\n\n")
	ranges := strings.Split(textArr[0], "\n")
	ids := strings.Split(textArr[1], "\n")
	for _, id := range ids {
		if id == "" {
			continue
		}
		idNum, err := strconv.Atoi(id)
		if err != nil {
			panic(err)
		}

		for _, r := range ranges {
			if r == "" {
				continue
			}
			parts := strings.Split(r, "-")
			min := parts[0]
			minNum, err := strconv.Atoi(min)
			if err != nil {
				panic(err)
			}
			max := parts[1]
			maxNum, err := strconv.Atoi(max)
			if err != nil {
				panic(err)
			}
			if idNum >= minNum && idNum <= maxNum {
				answer++
				break
			}
		}
	}

	return answer
}

type Range struct {
	min int
	max int
}

func fs2(f *os.File) int {
	answer := 0

	text, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	textArr := strings.Split(string(text), "\n\n")
	ranges := strings.Split(textArr[0], "\n")
	rangeArr := []*Range{}
	// remove overlapping ranges
	for _, r1 := range ranges {
		if r1 == "" {
			continue
		}
		parts := strings.Split(r1, "-")
		min := parts[0]
		minNum, err := strconv.Atoi(min)
		if err != nil {
			panic(err)
		}
		max := parts[1]
		maxNum, err := strconv.Atoi(max)
		if err != nil {
			panic(err)
		}
		rangeArr = append(rangeArr, &Range{min: minNum, max: maxNum})
	}

	// sort the ranges by min
	for i := 0; i < len(rangeArr)-1; i++ {
		for j := i + 1; j < len(rangeArr); j++ {
			if rangeArr[i].min > rangeArr[j].min {
				rangeArr[i], rangeArr[j] = rangeArr[j], rangeArr[i]
			}
		}
	}

	// remove overlapping ranges
	newRangeArr := []*Range{}
	for _, r := range rangeArr {
		overlap := false
		for _, existing := range newRangeArr {
			if (r.min >= existing.min && r.min <= existing.max) || (r.max >= existing.min && r.max <= existing.max) {
				// overlap
				if r.min < existing.min {
					existing.min = r.min
				}
				if r.max > existing.max {
					existing.max = r.max
				}
				overlap = true
				break
			}
		}
		if !overlap {
			newRangeArr = append(newRangeArr, r)
		}
	}
	rangeArr = newRangeArr

	for _, r := range rangeArr {
		fmt.Println(r)
		answer += (r.max - r.min + 1)
	}

	return answer
}
