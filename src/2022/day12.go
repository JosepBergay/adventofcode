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

func getNextSquares(curr Point, initial *Day12Initial) []Point {
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
		if initial.heightmap[p]-initial.heightmap[curr] > 1 {
			continue
		}

		out = append(out, p)
	}

	return out
}

// GetMinimumDistance finds Point with minimum distance that is not yet visited
func getMinimumDistance(visited map[Point]bool, distances map[Point]int) Point {
	min := math.MaxInt
	point := Point{}
	for p := range distances {
		if _, exists := visited[p]; !exists && distances[p] < min {
			min = distances[p]
			point = p
		}
	}
	return point
}

func (d *day12) Part1(input *Day12Initial) (string, error) {
	visited := make(map[Point]bool)
	// prev := make(map[Point]Point) // Map from point to previous point
	// Should use a priority queue, using a map so we don't have to reorder.
	// In return we have to loop the whole map to find the minimum value.
	distances := make(map[Point]int, len(input.heightmap))

	for p := range input.heightmap {
		distances[p] = math.MaxInt
	}

	distances[input.start] = 0

	for i := 0; i < len(input.heightmap); i++ {
		// Find node with minimum distance
		curr := getMinimumDistance(visited, distances)

		visited[curr] = true

		for _, n := range getNextSquares(curr, input) {
			if _, ok := visited[n]; ok {
				continue
			}

			dist := distances[curr] + 1 // Distance between neighbour nodes is always 1.
			if dist < distances[n] {
				distances[n] = dist
				// prev[n] = curr
			}
		}
	}

	return fmt.Sprint(distances[input.end]), nil
}

func (d *day12) Part2(input *Day12Initial) (string, error) {
	return "TODO", nil
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
