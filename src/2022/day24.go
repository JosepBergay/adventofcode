package aoc2022

import (
	"aoc2022/utils"
	"fmt"
	"math"
	"sort"
	"strings"
)

type blizzardMap map[Point][]rune

type valleyMap struct {
	start, end             Point
	blizzards              blizzardMap // Blizzard motions at this point. > | v | < | ^
	minX, maxX, minY, maxY int
	cycle                  int // when blizzard map is repeated
	curr                   Point
}

type day24 struct{}

func init() {
	Days[24] = &day24{}
}

func (d *day24) Parse(input string) (valleyMap, error) {
	lines := strings.Split(strings.TrimRight(input, "\n"), "\n")

	vm := valleyMap{}
	vm.blizzards = make(blizzardMap)

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

	// Find cycle in blizzard map (when we repeat the same blizzard config)
	curr := vm.blizzards
	i := 0
	for {
		blizzardMaps[i] = curr // update cache
		next := vm.moveBlizzards(curr)
		if vm.blizzards.isEqual(next) {
			break
		}
		i++
		curr = next
	}

	vm.cycle = len(blizzardMaps)

	return vm, nil
}

var blizzardMaps = make(map[int]blizzardMap) // blizzard maps for each minute, acts as cache

func (m blizzardMap) isEqual(m2 blizzardMap) bool {
	if len(m) != len(m2) {
		return false
	}

	for k, v := range m {
		_, ok := m2[k]
		if !ok {
			return false
		}
		if len(v) != len(m2[k]) {
			return false
		}
		for i, r := range v {
			if r != m2[k][i] {
				return false
			}
		}
	}

	return true
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

func (vm *valleyMap) moveBlizzards(from blizzardMap) blizzardMap {
	dest := make(blizzardMap)

	for p, blizz := range from {
		for _, r := range blizz {
			next := vm.moveBlizzard(r, p)
			dest[next] = append(dest[next], r)
		}
	}

	return dest
}

func (vm valleyMap) getBlizzardMapAt(minute int) blizzardMap {
	// Optimization: get map from cache. This alone takes it down from ~24m to ~1s (on part1) wow!
	nextMap := blizzardMaps[minute%vm.cycle]
	if nextMap == nil {
		panic("should be init in parse")
	}
	return nextMap
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

var seen = make(map[int]map[Point]bool) // decisions taken each minute
var minSteps = math.MaxInt

func (vm *valleyMap) crossValley(step int) int {
	curr := vm.curr
	if curr == vm.end { // reached end!
		return step
	}

	if step > minSteps { // longer than fastest path
		return 0
	}

	currMap := vm.getBlizzardMapAt(step)
	if _, ok := currMap[curr]; ok { // hit a blizzard
		return 0
	}

	if vm.isOutOfBounds(curr) && curr != vm.start { // out of bounds
		return 0
	}

	if seen[step][curr] && curr != vm.start { // been here at this point in time
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
		// nextMap := vm.getBlizzardMapAt(step+1)
		vm.curr = next
		steps := vm.crossValley(step + 1)
		if steps != 0 && minSteps > steps {
			minSteps = steps
		}
	}

	return minSteps
}

func (vm *valleyMap) crossValleyEfficiently(startAt int) int {
	type PointMinute struct {
		p   Point
		min int
	}

	visited := make(map[PointMinute]bool)

	queue := utils.PriorityQueue[PointMinute]{}
	queue.Insert(&PointMinute{vm.start, startAt}, 0)

	for queue.Len() > 0 {
		pm := queue.Delete()

		if pm.p == vm.end {
			return pm.min + 1
		}

		nextMap := vm.getBlizzardMapAt(pm.min + 1)
		for _, nextMove := range prioritizeNextMoves(pm.p, vm.end) {
			if _, ok := nextMap[pm.p]; ok {
				continue // hit a blizzard
			}
			if vm.isOutOfBounds(nextMove) && nextMove != vm.start && nextMove != vm.end {
				continue
			}
			nextPm := PointMinute{p: nextMove, min: pm.min + 1}
			if visited[nextPm] {
				continue
			}
			visited[nextPm] = true

			queue.Insert(&nextPm, getManhattanDistance(nextMove, vm.end)+nextPm.min)
		}
	}

	return -1
}

func (d *day24) Part1(vm valleyMap) (string, error) {
	vm.curr = vm.start
	steps := vm.crossValley(0)

	return fmt.Sprint(steps), nil
}

func (d *day24) Part2(vm valleyMap) (string, error) {
	// If we cross valley like P1 we run into stack overflow so let's do it "efficiently"
	steps := vm.crossValleyEfficiently(0)

	vm.start, vm.end = vm.end, vm.start // switch directions

	steps = vm.crossValleyEfficiently(steps)

	vm.start, vm.end = vm.end, vm.start // switch directions

	steps = vm.crossValleyEfficiently(steps)

	return fmt.Sprint(steps), nil
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
