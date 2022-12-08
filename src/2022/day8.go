package aoc2022

import (
	"bufio"
	"fmt"
	"strings"
)

type day8 struct{}

func init() {
	Days[8] = &day8{}
}

func (d *day8) Parse(input string) ([][]int, error) {
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	matrix := make([][]int, 0)
	for scanner.Scan() {
		trees := scanner.Text()

		row := make([]int, len(trees))

		for i, tree := range trees {
			row[i] = int(tree - '0')
		}

		matrix = append(matrix, row)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}

	return matrix, nil
}

type Point struct{ x, y int }

type Path struct {
	dir      Point
	distance int
}

func getPossiblePaths(point Point, matrix [][]int) [4]Path {
	return [4]Path{
		{Point{0, 1}, len(matrix) - 1 - point.y},
		{Point{1, 0}, len(matrix[0]) - 1 - point.x},
		{Point{0, -1}, point.y},
		{Point{-1, 0}, point.x},
	}
}

func createPathWalker(start Point, path Path) func() *Point {
	i := 1
	return func() *Point {
		if i > path.distance {
			return nil
		}

		p := &Point{start.x + (path.dir.x * i), start.y + (path.dir.y * i)}
		i++

		return p
	}
}

func isVisibleFromOutside(point Point, matrix [][]int) bool {
	paths := getPossiblePaths(point, matrix)

	currHeight := matrix[point.y][point.x]

	// For each direction
	for _, path := range paths {
		// Walk path till the edge
		walk := createPathWalker(point, path)

		reachedEdge := true
		for {
			p := walk()

			if p == nil {
				break
			}

			if matrix[p.y][p.x] >= currHeight {
				// Not visible, stop walking this path
				reachedEdge = false
				break
			}
		}

		if reachedEdge {
			return true
		}
	}

	return false
}

func (d *day8) Part1(matrix [][]int) (string, error) {
	total := 0

	for y, rows := range matrix {
		for x := range rows {
			if isVisibleFromOutside(Point{x, y}, matrix) {
				total += 1
			}
		}
	}

	return fmt.Sprint(total), nil
}

func computeScenicScore(point Point, matrix [][]int) int {
	paths := getPossiblePaths(point, matrix)

	currHeight := matrix[point.y][point.x]

	score := 1

	// For each direction
	for _, path := range paths {
		// Walk path till the edge
		walk := createPathWalker(point, path)

		mult := 0
		for {
			p := walk()

			if p == nil {
				break
			}

			mult += 1
			if matrix[p.y][p.x] >= currHeight {
				// Not visible, stop walking this path
				break
			}
		}

		if mult == 0 {
			return 0
		}

		score = score * mult
	}

	return score
}

func (d *day8) Part2(matrix [][]int) (string, error) {
	score := 0

	for y, rows := range matrix {
		for x := range rows {
			newScore := computeScenicScore(Point{x, y}, matrix)
			if newScore > score {
				score = newScore
			}
		}
	}

	return fmt.Sprint(score), nil
}

func (d *day8) Exec(input string) (*DayResult, error) {
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
