package aoc2022

import (
	"bufio"
	"fmt"
	"strings"
)

var dirs = [4]Point{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} // N, S, W, E

type elf struct{}

type day23 struct {
	elvesMap map[Point]*elf
}

func init() {
	Days[23] = &day23{
		elvesMap: make(map[Point]*elf),
	}
}

func (d *day23) Parse(input string) (string, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))

	scanner.Split(bufio.ScanLines)

	y := 0
	for scanner.Scan() {
		line := scanner.Text()

		var x int
		for _, v := range line {
			if v == '#' {
				p := Point{x, y}
				d.elvesMap[p] = nil
			}
			x++
		}
		y++
	}

	return "", nil
}

func getAdjacents(p Point) [8]Point {
	n := Point{p.x, p.y - 1}
	s := Point{p.x, p.y + 1}
	e := Point{p.x + 1, p.y}
	w := Point{p.x - 1, p.y}

	nw := Point{p.x - 1, p.y - 1}
	ne := Point{p.x + 1, p.y - 1}

	sw := Point{p.x - 1, p.y + 1}
	se := Point{p.x + 1, p.y + 1}

	return [8]Point{n, s, e, w, nw, ne, sw, se}
}

func areAllPositionsEmpty(points []Point, elvesMap map[Point]*elf) bool {
	for _, p := range points {
		if _, ok := elvesMap[p]; ok {
			return false
		}
	}

	return true
}

func hasSomeAdjacents(p Point, elvesMap map[Point]*elf) bool {
	adj := getAdjacents(p)

	return !areAllPositionsEmpty(adj[:], elvesMap)
}

func getNextValidPosition(p Point, elvesMap map[Point]*elf, lastDir int) Point {
	var next *Point

	for i := 0; i < 4; i++ {
		var mustCheck [3]Point

		switch (lastDir + i) % 4 {
		case 0:
			// Check North
			mustCheck = [3]Point{
				{p.x, p.y - 1}, // N
				{p.x + 1, p.y - 1},
				{p.x - 1, p.y - 1}}
		case 1:
			// Check South
			mustCheck = [3]Point{
				{p.x, p.y + 1}, // S
				{p.x + 1, p.y + 1},
				{p.x - 1, p.y + 1}}
		case 2:
			// Check West
			mustCheck = [3]Point{
				{p.x - 1, p.y}, // W
				{p.x - 1, p.y - 1},
				{p.x - 1, p.y + 1}}
		case 3:
			// Check West
			mustCheck = [3]Point{
				{p.x + 1, p.y}, // E
				{p.x + 1, p.y - 1},
				{p.x + 1, p.y + 1}}
		}

		if ok := areAllPositionsEmpty(mustCheck[:], elvesMap); ok {
			next = &mustCheck[0]
			break
		}
	}

	if next == nil {
		return p
	}

	return *next
}

func (d *day23) Part1(input string) (string, error) {
	rounds := 10

	var dirIdx int // index of proposed direction

	for i := 0; i < rounds; i++ {
		// First half
		tmpMap := make(map[Point]Point) // map from next point to previous
		repeated := make(map[Point]bool)

		for p := range d.elvesMap {
			yes := hasSomeAdjacents(p, d.elvesMap)

			var next Point
			if !yes {
				// The elf does not do anything
				next = p
			} else {
				next = getNextValidPosition(p, d.elvesMap, dirIdx)
			}

			if _, ok := tmpMap[next]; ok {
				repeated[next] = true
			}
			tmpMap[next] = p
		}

		// Second half
		for next, prev := range tmpMap {
			if ok := repeated[next]; ok {
				continue
			}
			delete(d.elvesMap, prev)
			d.elvesMap[next] = nil
		}

		dirIdx++
		dirIdx %= 4
	}

	var minX, maxX, minY, maxY int
	for p := range d.elvesMap {
		if p.x < minX {
			minX = p.x
		}
		if p.x > maxX {
			maxX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	surface := (maxX - minX + 1) * (maxY - minY + 1)
	emptyTiles := surface - len(d.elvesMap)

	return fmt.Sprint(emptyTiles), nil
}

func (d *day23) Part2(input string) (string, error) {
	return "TODO", nil
}

func (d *day23) Exec(input string) (*DayResult, error) {
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
