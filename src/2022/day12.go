package aoc2022

import (
	"bufio"
	"fmt"
	"math"
	"strings"
)

type day12 struct{}

func init() {
	Days[12] = &day12{}
}

type Day12Initial struct {
	start, end Point
	heightmap  map[Point]rune
	maxX, maxY int
}

func (d *day12) Parse(input string) (*Day12Initial, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)

	heightmap := make(map[Point]rune)

	start, end := Point{}, Point{}
	maxX, maxY := 0, 0

	for scanner.Scan() {
		line := scanner.Text()

		for x, r := range line {
			p := Point{x, maxY}

			heightmap[p] = r

			switch r {
			case 'S':
				start = Point{x, maxY}
				heightmap[p] = 'a'
			case 'E':
				end = Point{x, maxY}
				heightmap[p] = 'z'
			}
			maxX = x
		}

		maxY++
	}

	maxY--

	return &Day12Initial{start, end, heightmap, maxX, maxY}, nil
}

func getAdjacentSquares(curr Point, initial *Day12Initial, skipFn func(Point) bool) []Point {
	dirs := [4]Point{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	out := make([]Point, 0)

	for _, d := range dirs {
		p := Point{curr.x + d.x, curr.y + d.y}

		// Out of bounds
		if p.x < 0 || p.x > initial.maxX || p.y < 0 || p.y > initial.maxY {
			continue
		}

		// Damn, can't get out my climbing gear!
		if skipFn(p) {
			continue
		}

		out = append(out, p)
	}

	return out
}

var intfinity = int(math.Float64bits(math.Inf(0)))

func hasUnvisited(visited map[Point]bool, distances map[Point]int) bool {
	for p, v := range visited {
		if !v && distances[p] == intfinity {
			return true
		}
	}

	return false
}

// GetMinimumDistance finds Point with minimum distance that is not yet visited
func getMinimumUnvisitedDistance(visited map[Point]bool, distances map[Point]int) Point {
	min := intfinity
	point := Point{}
	for p := range distances {
		if yes := visited[p]; !yes && distances[p] < min {
			min = distances[p]
			point = p
		}
	}
	return point
}

func getDistancesFromStart(
	start Point,
	reachedEnd func(Point) bool,
	input *Day12Initial,
	getNeighbours func(Point) []Point,
) map[Point]int {
	visited := make(map[Point]bool)
	// prev := make(map[Point]Point) // Map from point to previous point
	// Should use a priority queue, using a map so we don't have to reorder.
	// In return we have to loop the whole map to find the minimum value.
	distances := make(map[Point]int, len(input.heightmap))

	for p := range input.heightmap {
		distances[p] = intfinity
		visited[p] = false
	}

	distances[start] = 0

	for hasUnvisited(visited, distances) {
		// Find node with minimum distance
		curr := getMinimumUnvisitedDistance(visited, distances)

		if reachedEnd(curr) {
			break
		}

		visited[curr] = true

		for _, n := range getNeighbours(curr) {
			if yes := visited[n]; yes {
				continue
			}

			dist := distances[curr] + 1 // Distance between neighbour nodes is always 1.
			if dist < distances[n] {
				distances[n] = dist
				// prev[n] = curr
			}
		}
	}

	return distances
}

func (d *day12) Part1(input *Day12Initial) (string, error) {
	reachedEnd := func(curr Point) bool {
		return curr == input.end
	}

	getNeighbours := func(curr Point) []Point {
		return getAdjacentSquares(curr, input, func(p Point) bool {
			// Damn, can't get out my climbing gear!
			return input.heightmap[p]-input.heightmap[curr] > 1
		})
	}

	distances := getDistancesFromStart(input.start, reachedEnd, input, getNeighbours)

	return fmt.Sprint(distances[input.end]), nil
}

func (d *day12) Part2(input *Day12Initial) (string, error) {
	endRune := 'a'

	reachedEnd := func(curr Point) bool {
		return input.heightmap[curr] == endRune
	}

	getNeighbours := func(curr Point) []Point {
		return getAdjacentSquares(curr, input, func(p Point) bool {
			// What about my downhill gear?!
			step := input.heightmap[p] - input.heightmap[curr]
			return step < -1
		})
	}

	distances := getDistancesFromStart(input.end, reachedEnd, input, getNeighbours)

	min := intfinity
	for p, dist := range distances {
		if input.heightmap[p] == endRune && dist < min {
			min = dist
		}
	}

	return fmt.Sprint(min), nil
}

func (d *day12) Exec(input string) (*DayResult, error) {
	parsed, err := d.Parse(input)

	if err != nil {
		return nil, err
	}

	part1, err := d.Part1(parsed)

	if err != nil {
		return nil, err
	}

	part2, err := d.Part2(parsed)

	if err != nil {
		return nil, err
	}

	result := &DayResult{part1, part2}

	return result, nil
}
