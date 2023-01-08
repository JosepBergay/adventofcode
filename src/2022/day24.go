package aoc2022

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

type valleyMap struct {
	start, end             Point
	blizzards              map[Point][]rune // Blizzard motions at this point. > | v | < | ^
	minX, maxX, minY, maxY int
}

type day24 struct{}

func init() {
	Days[24] = &day24{}
}

func (d *day24) Parse(input string) (valleyMap, error) {
	lines := strings.Split(strings.TrimRight(input, "\n"), "\n")

	vm := valleyMap{}
	vm.blizzards = make(map[Point][]rune)

	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			switch lines[y][x] {
			case '#':
			case '.':
				if y == 0 {
					vm.start = Point{x, y}
				} else if y == len(lines)-1 {
					vm.end = Point{x, y}
				}
			default:
				vm.blizzards[Point{x, y}] = append(vm.blizzards[Point{x, y}], rune(lines[y][x]))
			}
		}
	}

	vm.minX, vm.minY = 1, 1
	vm.maxX, vm.maxY = len(lines[0])-2, len(lines)-2

	return vm, nil
}

func (vm *valleyMap) moveBlizzard(r rune, p Point) Point {
	var next Point
	switch r {
	case 'v':
		next = Point{p.x, p.y + 1}
	case '^':
		next = Point{p.x, p.y - 1}
	case '>':
		next = Point{p.x + 1, p.y}
	case '<':
		next = Point{p.x - 1, p.y}
	}

	if next.x > vm.maxX {
		next.x = vm.minX
	}
	if next.x < vm.minX {
		next.x = vm.maxX
	}
	if next.y > vm.maxY {
		next.y = vm.minY
	}
	if next.y < vm.minY {
		next.y = vm.maxY
	}

	return next
}

func (vm *valleyMap) moveBlizzards(from map[Point][]rune) map[Point][]rune {
	dest := make(map[Point][]rune)

	for p, blizz := range from {
		for _, r := range blizz {
			next := vm.moveBlizzard(r, p)
			dest[next] = append(dest[next], r)
		}
	}

	return dest
}

func (vm *valleyMap) isOutOfBounds(p Point) bool {
	if p.x < vm.minX || p.x > vm.maxX {
		return true
	}
	if p.y < vm.minY || p.y > vm.maxY {
		return true
	}

	return false
}

var dirs = [5]Point{{0, 1}, {0, -1}, {-1, 0}, {1, 0}, {0, 0}}

func prioritizeNextMoves(curr Point, end Point) []Point {
	next := make([]Point, 5)

	for i, dir := range dirs {
		next[i] = Point{curr.x + dir.x, curr.y + dir.y}
	}

	sort.Slice(next, func(i, j int) bool {
		return getManhattanDistance(next[i], end) < getManhattanDistance(next[j], end)
	})

	return next
}

var blizzardMaps = make(map[int]map[Point][]rune) // blizzard maps for each minute

var seen = make(map[int]map[Point]bool) // decisions taken each minute
var minSteps = math.MaxInt

func (vm *valleyMap) crossValley(curr Point, step int, m map[Point][]rune) int {
	if curr == vm.end { // reached end!
		return step
	}

	if step > minSteps { // longer than fastest path
		return 0
	}

	if _, ok := m[curr]; ok { // hit a blizzard
		return 0
	}

	if vm.isOutOfBounds(curr) && curr != vm.start { // out of bounds
		return 0
	}

	if seen[step][curr] { // been here at this point in time
		return 0
	}

	if getManhattanDistance(curr, vm.end)+step > minSteps { // would take longer than fastest path
		return 0
	}

	if seen[step] == nil {
		seen[step] = make(map[Point]bool)
	}
	seen[step][curr] = true

	for _, next := range prioritizeNextMoves(curr, vm.end) {
		// Optimization: get map from cache. This alone takes it down from ~24m to ~1s wow!
		nextMap := blizzardMaps[step+1]
		if nextMap == nil {
			nextMap = vm.moveBlizzards(m)
			blizzardMaps[step+1] = nextMap
		}

		steps := vm.crossValley(next, step+1, nextMap)
		if steps != 0 && minSteps > steps {
			minSteps = steps
		}
	}

	return minSteps
}

func (d *day24) Part1(vm valleyMap) (string, error) {
	steps := vm.crossValley(vm.start, 0, vm.blizzards)

	return fmt.Sprint(steps), nil
}

func (d *day24) Part2(input valleyMap) (string, error) {
	return "TODO", nil
}

func (d *day24) Exec(input string) (*DayResult, error) {
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
