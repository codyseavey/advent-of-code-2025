package main

import (
	"bufio"
	"image"
	"os"
	"strconv"
	"strings"
)

func fs1(f *os.File) int {
	scanner := bufio.NewScanner(f)
	tiles := [][2]int{}
	for scanner.Scan() {
		line := scanner.Text()
		points := strings.Split(line, ",")
		x, _ := strconv.Atoi(points[0])
		y, _ := strconv.Atoi(points[1])
		tiles = append(tiles, [2]int{x, y})
	}

	rectAreas := []int{}

	for _, tile1 := range tiles {
		for _, tile2 := range tiles {
			rectLength := 1
			rectHeight := 1
			if tile1[0] > tile2[0] {
				rectLength = tile1[0] - tile2[0] + 1
			}
			if tile2[0] > tile1[0] {
				rectLength = tile2[0] - tile1[0] + 1
			}
			if tile1[1] > tile2[1] {
				rectHeight = tile1[1] - tile2[1] + 1
			}
			if tile2[1] > tile1[1] {
				rectHeight = tile2[1] - tile1[1] + 1
			}

			rectArea := rectLength * rectHeight
			rectAreas = append(rectAreas, rectArea)
		}
	}

	// Find the largest rectangle area
	largestArea := 0
	for _, area := range rectAreas {
		if area > largestArea {
			largestArea = area
		}
	}

	return largestArea
}



func fs2(f *os.File) int {
	scanner := bufio.NewScanner(f)
	var poly []image.Point
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		poly = append(poly, image.Point{X: x, Y: y})
	}

	maxArea := 0

	n := len(poly)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			p1 := poly[i]
			p2 := poly[j]

			// Define Rect R
			minX := p1.X
			maxX := p2.X
			if minX > maxX {
				minX, maxX = maxX, minX
			}
			minY := p1.Y
			maxY := p2.Y
			if minY > maxY {
				minY, maxY = maxY, minY
			}

			// Area
			width := maxX - minX + 1
			height := maxY - minY + 1
			area := width * height

			if area <= maxArea {
				continue
			}

			// Validity Check
			// image.Rect is [Min, Max)
			// So we want Min to be (minX, minY) and Max to be (maxX+1, maxY+1)
			rect := image.Rect(minX, minY, maxX+1, maxY+1)

			if isRectValid(rect, poly) {
				maxArea = area
			}
		}
	}
	return maxArea
}

func isRectValid(r image.Rectangle, poly []image.Point) bool {
	// 1. Center Check
	// Center of the rectangle in continuous space
	cx := float64(r.Min.X) + float64(r.Dx()-1)/2.0
	cy := float64(r.Min.Y) + float64(r.Dy()-1)/2.0
	if !isPointInPoly(cx, cy, poly) {
		return false
	}

	// 2. Vertex Check (Strictly Inside)
	// If a polygon vertex is strictly inside the rectangle, it's invalid
	// because that implies the boundary enters the rectangle.
	for _, v := range poly {
		if v.X > r.Min.X && v.X < r.Max.X-1 && v.Y > r.Min.Y && v.Y < r.Max.Y-1 {
			return false
		}
	}

	// 3. Edge Intersection Check
	// If any polygon edge intersects the interior of the rectangle.
	n := len(poly)
	for i := 0; i < n; i++ {
		p1 := poly[i]
		p2 := poly[(i+1)%n]

		if edgeIntersectsRectInterior(p1, p2, r) {
			return false
		}
	}

	return true
}

func isPointInPoly(x, y float64, poly []image.Point) bool {
	inside := false
	n := len(poly)
	for i := 0; i < n; i++ {
		p1 := poly[i]
		p2 := poly[(i+1)%n]

		y1 := float64(p1.Y)
		y2 := float64(p2.Y)
		x1 := float64(p1.X)
		x2 := float64(p2.X)

		if (y1 > y) != (y2 > y) {
			intersectX := x1 + (y-y1)*(x2-x1)/(y2-y1)
			if x < intersectX {
				inside = !inside
			}
		}
	}
	return inside
}

func edgeIntersectsRectInterior(p1, p2 image.Point, r image.Rectangle) bool {
	rminX := r.Min.X
	rmaxX := r.Max.X - 1
	rminY := r.Min.Y
	rmaxY := r.Max.Y - 1

	if p1.X == p2.X {
		// Vertical Edge
		x := p1.X
		if x <= rminX || x >= rmaxX {
			return false
		}
		y1 := p1.Y
		y2 := p2.Y
		if y1 > y2 {
			y1, y2 = y2, y1
		}
		// Overlap of [y1, y2] with (rminY, rmaxY)
		return max(y1, rminY) < min(y2, rmaxY)

	} else {
		// Horizontal Edge
		y := p1.Y
		if y <= rminY || y >= rmaxY {
			return false
		}
		x1 := p1.X
		x2 := p2.X
		if x1 > x2 {
			x1, x2 = x2, x1
		}
		// Overlap of [x1, x2] with (rminX, rmaxX)
		return max(x1, rminX) < min(x2, rmaxX)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
