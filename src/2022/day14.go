package aoc2022

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type day14 struct {
	rocks        map[Point]bool
	furthestRock Point
	start        Point
}

func init() {
	Days[14] = &day14{}
}

func (d *day14) Parse(input string) (map[Point]bool, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)

	points := make(map[Point]bool)

	for scanner.Scan() {
		line := scanner.Text()

		pointsStr := strings.Split(line, " -> ")

		from := Point{}
		for _, pointStr := range pointsStr {
			coords := strings.Split(pointStr, ",")
			x, err := strconv.Atoi(coords[0])
			if err != nil {
				return nil, err
			}
			y, err := strconv.Atoi(coords[1])
			if err != nil {
				return nil, err
			}

			to := Point{x, y}
			points[to] = true

			if to.y > d.furthestRock.y {
				d.furthestRock = to
			}

			empty := Point{}
			if from == empty {
				from = to
				continue
			}

			min := math.Min(float64(from.x), float64(to.x))
			max := math.Max(float64(from.x), float64(to.x))
			for i := min; i < max; i++ {
				points[Point{int(i), to.y}] = true
			}

			min = math.Min(float64(from.y), float64(to.y))
			max = math.Max(float64(from.y), float64(to.y))
			for i := min; i < max; i++ {
				points[Point{to.x, int(i)}] = true
			}

			from = to
		}
	}

	d.rocks = points
	d.start = Point{500, 0}

	return points, nil
}

func copyMap[K comparable, V any](m map[K]V) map[K]V {
	m2 := make(map[K]V, len(m))
	for k := range m {
		m2[k] = m[k]
	}
	return m2
}

func (d *day14) simulateFallingSand(
	isEmptySpace func(next Point, cave map[Point]bool) bool,
	stop func(sandUnit Point) bool,
) int {
	dirs := [3]Point{
		{0, 1},
		{-1, 1},
		{1, 1},
	}

	sandUnits := 0

	cave := copyMap(d.rocks)

	for {
		sandUnit := Point{d.start.x, d.start.y}

		for {
			moved := false
			for _, dir := range dirs {
				next := Point{sandUnit.x + dir.x, sandUnit.y + dir.y}

				if isEmptySpace(next, cave) {
					// Found empty space so move sandUnit there.
					sandUnit = next
					moved = true
					break
				}
			}

			if stop(sandUnit) {
				return sandUnits
			}

			if !moved {
				// Sand unit came to rest, let's break and produce another one.
				cave[sandUnit] = true
				break
			}
		}

		sandUnits++
	}
}

func (d *day14) Part1(input map[Point]bool) (string, error) {
	sandUnits := d.simulateFallingSand(
		func(next Point, cave map[Point]bool) bool {
			// Empty space.
			return !cave[next]
		},
		func(sandUnit Point) bool {
			// Falling to the endless void.
			return sandUnit.y > d.furthestRock.y
		})

	return fmt.Sprint(sandUnits), nil
}

func (d *day14) Part2(input map[Point]bool) (string, error) {
	floorHeight := d.furthestRock.y + 2

	sandUnits := d.simulateFallingSand(
		func(next Point, cave map[Point]bool) bool {
			// Empty space and not reached the floor.
			return !cave[next] && next.y < floorHeight
		},
		func(sandUnit Point) bool {
			// Sand unit didn't move at all.
			return sandUnit == d.start
		})

	// We stop just before adding sandUnit to start space so we need to add one.
	return fmt.Sprint(sandUnits + 1), nil
}

func (d *day14) Exec(input string) (*DayResult, error) {
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
