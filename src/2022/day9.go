package aoc2022

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type day9 struct{}

func init() {
	Days[9] = &day9{}
}

func parseLine(line string) ([]Point, error) {
	spl := strings.Split(line, " ")

	d := spl[0]
	n, err := strconv.Atoi(spl[1])

	if err != nil {
		return nil, err
	}

	moves := make([]Point, 0)

	m := map[string]Point{
		"R": {1, 0},
		"L": {-1, 0},
		"U": {0, 1},
		"D": {0, -1},
	}

	for i := 0; i < n; i++ {
		moves = append(moves, m[d])
	}

	return moves, nil
}

func (d *day9) Parse(input string) ([]Point, error) {
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)

	scanner.Split(bufio.ScanLines)

	moves := make([]Point, 0)

	for scanner.Scan() {
		line := scanner.Text()

		m, err := parseLine(line)
		if err != nil {
			return nil, err
		}

		moves = append(moves, m...)
	}

	return moves, nil
}

func addToVisited(visited map[int]map[int]bool, p Point) {
	if _, exists := visited[p.x]; !exists {
		visited[p.x] = make(map[int]bool)
	}
	visited[p.x][p.y] = true
}

func mustMoveKnot(h, t Point) bool {
	return math.Abs(float64(h.x-t.x)) > 1 || math.Abs(float64(h.y-t.y)) > 1
}

func countVisited(visited map[int]map[int]bool) int {
	count := 0

	for _, row := range visited {
		count += len(row)
	}

	return count
}

func moveKnot(h, t *Point) {
	if h.x != t.x && h.y != t.y {
		// Not in the same row or column, so move diagonally
		if h.x-t.x > 0 {
			t.x += 1
		} else {
			t.x -= 1
		}

		if h.y-t.y > 0 {
			t.y += 1
		} else {
			t.y -= 1
		}

		return
	}

	if h.x-t.x > 1 {
		t.x += 1
	} else if h.x-t.x < -1 {
		t.x -= 1
	} else if h.y-t.y > 1 {
		t.y += 1
	} else if h.y-t.y < -1 {
		t.y -= 1
	}
}

func (d *day9) Part1(moves []Point) (string, error) {
	visited := make(map[int]map[int]bool)

	h := Point{}
	t := Point{}

	addToVisited(visited, t)

	for _, move := range moves {
		h.x += move.x
		h.y += move.y

		if mustMoveKnot(h, t) {
			moveKnot(&h, &t)

			addToVisited(visited, t)
		}
	}

	count := countVisited(visited)

	return fmt.Sprint(count), nil
}

func (d *day9) Part2(moves []Point) (string, error) {
	visited := make(map[int]map[int]bool)

	knots := [10]Point{}

	addToVisited(visited, knots[9])

	for _, move := range moves {
		// Head always moves
		knots[0].x += move.x
		knots[0].y += move.y

		for i := 1; i < 10; i++ {
			if mustMoveKnot(knots[i-1], knots[i]) {
				moveKnot(&knots[i-1], &knots[i])
			}
		}

		addToVisited(visited, knots[9])
	}

	count := countVisited(visited)

	return fmt.Sprint(count), nil
}

func (d *day9) Exec(input string) (*DayResult, error) {
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
