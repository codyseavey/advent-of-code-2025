package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	X, Y, Z int
}

type Pair struct {
	Idx1, Idx2 int
	DistSq     int
}

func distSq(p1, p2 Point) int {
	dx := p1.X - p2.X
	dy := p1.Y - p2.Y
	dz := p1.Z - p2.Z
	return dx*dx + dy*dy + dz*dz
}

type UnionFind struct {
	parent []int
	size   []int
}

func newUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return &UnionFind{parent: parent, size: size}
}

func (uf *UnionFind) Find(i int) int {
	if uf.parent[i] != i {
		uf.parent[i] = uf.Find(uf.parent[i])
	}
	return uf.parent[i]
}

func (uf *UnionFind) Union(i, j int) {
	rootI := uf.Find(i)
	rootJ := uf.Find(j)
	if rootI != rootJ {
		// Attach smaller tree to larger tree
		if uf.size[rootI] < uf.size[rootJ] {
			rootI, rootJ = rootJ, rootI
		}
		uf.parent[rootJ] = rootI
		uf.size[rootI] += uf.size[rootJ]
	}
}

func fs1(f *os.File) int {
	var points []Point
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		if len(parts) < 3 {
			continue
		}
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		points = append(points, Point{x, y, z})
	}

	var pairs []Pair
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			d := distSq(points[i], points[j])
			pairs = append(pairs, Pair{i, j, d})
		}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].DistSq < pairs[j].DistSq
	})

	var limit int
	if len(points) == 20 {
		limit = 10
	}
	if len(points) == 1000 {
		limit = 1000
	}

	uf := newUnionFind(len(points))
	for k := 0; k < limit; k++ {
		uf.Union(pairs[k].Idx1, pairs[k].Idx2)
	}

	var sizes []int
	for i := 0; i < len(points); i++ {
		if uf.parent[i] == i {
			sizes = append(sizes, uf.size[i])
		}
	}

	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})

	if len(sizes) == 0 {
		return 0
	}
	ans := 1
	count := 0
	for _, s := range sizes {
		ans *= s
		count++
		if count == 3 {
			break
		}
	}

	return ans
}

func fs2(f *os.File) int {
	var points []Point
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		if len(parts) < 3 {
			continue
		}
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		points = append(points, Point{x, y, z})
	}

	var pairs []Pair
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			d := distSq(points[i], points[j])
			pairs = append(pairs, Pair{i, j, d})
		}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].DistSq < pairs[j].DistSq
	})

	ans := 0
	uf := newUnionFind(len(points))
	// connect all points until there is a single circuit
	for uf.size[uf.Find(0)] < len(points) {
		for _, p := range pairs {
			if uf.Find(p.Idx1) != uf.Find(p.Idx2) {
				uf.Union(p.Idx1, p.Idx2)
				ans = points[p.Idx1].X * points[p.Idx2].X
				break
			}
		}
	}

	return ans
}

