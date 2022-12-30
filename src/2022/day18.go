package aoc2022

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Cube [3]int

type Matrix map[int]map[int]map[int]int

type day18 struct {
	matrix Matrix
	maxX, maxY, maxZ,
	minX, minY, minZ int
}

func init() {
	Days[18] = &day18{}
}

func (d *day18) addMinMax(x, y, z int) {
	if len(d.matrix) == 0 {
		d.maxX, d.minX = x, x
		d.maxY, d.minY = y, y
		d.maxZ, d.minZ = z, z
		return
	}

	if x > d.maxX {
		d.maxX = x
	} else if x < d.minX {
		d.minX = x
	}

	if y > d.maxY {
		d.maxY = y
	} else if y < d.minY {
		d.minY = y
	}

	if z > d.maxZ {
		d.maxZ = z
	} else if z < d.minZ {
		d.minZ = z
	}
}

func addCube(m *Matrix, x, y, z int) {
	if _, ok := (*m)[x]; !ok {
		(*m)[x] = make(map[int]map[int]int, 0)
	}

	if _, ok := (*m)[x][y]; !ok {
		(*m)[x][y] = make(map[int]int, 0)
	}

	(*m)[x][y][z] = 1
}

func (d *day18) addCube(x, y, z int) {
	d.addMinMax(x, y, z)
	addCube(&d.matrix, x, y, z)
}

func (d *day18) Parse(input string) (Matrix, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)

	d.matrix = make(Matrix, 0)

	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, ",")

		coords := [3]int{}

		for i, v := range split {
			c, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			coords[i] = c
		}

		d.addCube(coords[0], coords[1], coords[2])
	}

	return d.matrix, nil
}

func getExposedSidesCount(m Matrix, x, y, z int) int {
	out := 0

	if _, ok := m[x+1][y][z]; !ok {
		out++
	}
	if _, ok := m[x-1][y][z]; !ok {
		out++
	}
	if _, ok := m[x][y+1][z]; !ok {
		out++
	}
	if _, ok := m[x][y-1][z]; !ok {
		out++
	}
	if _, ok := m[x][y][z+1]; !ok {
		out++
	}
	if _, ok := m[x][y][z-1]; !ok {
		out++
	}

	return out
}

func (d *day18) Part1(input Matrix) (string, error) {
	out := 0
	for x := range input {
		for y := range input[x] {
			for z := range input[x][y] {
				out += getExposedSidesCount(input, x, y, z)
			}
		}
	}

	return fmt.Sprint(out), nil
}

func (d *day18) Part2(input Matrix) (string, error) {
	// Imagine a bounding box that fully contains the whole input.
	// Start at any corner of this box. Then walk all points, if curr point exists in input treat it
	// like a wall and stop, else add it to the new matrix. This new matrix will have a hole inside
	// with the shape of the input.

	directions := [6][3]int{
		{1, 0, 0},
		{-1, 0, 0},
		{0, 1, 0},
		{0, -1, 0},
		{0, 0, 1},
		{0, 0, -1},
	}

	box := make(Matrix)
	seen := make(map[[3]int]bool)

	isOutOfBounds := func(point [3]int) bool {
		if point[0] < d.minX-1 || point[0] > d.maxX+1 {
			return true
		}
		if point[1] < d.minY-1 || point[1] > d.maxY+1 {
			return true
		}
		if point[2] < d.minZ-1 || point[2] > d.maxZ+1 {
			return true
		}
		return false
	}

	var walkRec func(curr [3]int)
	walkRec = func(curr [3]int) {
		if _, ok := seen[curr]; ok {
			return
		}
		if isOutOfBounds(curr) {
			return
		}
		if _, ok := d.matrix[curr[0]][curr[1]][curr[2]]; ok {
			return
		}

		seen[curr] = true
		addCube(&box, curr[0], curr[1], curr[2])

		for _, dir := range directions {
			next := [3]int{curr[0] + dir[0], curr[1] + dir[1], curr[2] + dir[2]}
			walkRec(next)
		}
	}

	start := [3]int{d.minX - 1, d.minY - 1, d.minZ - 1}
	walkRec(start)

	surfaceStr, err := d.Part1(box)
	if err != nil {
		return "", err
	}
	surface, err := strconv.Atoi(surfaceStr)
	if err != nil {
		return "", err
	}

	// If input min-max is 1-4 then bounding box min-max is 0-5 so box delta must be 6.
	deltaZ := d.maxZ - d.minZ + 3
	deltaX := d.maxX - d.minX + 3
	deltaY := d.maxY - d.minY + 3
	boxOutterSurface := (deltaZ)*(deltaX)*2 + (deltaZ)*(deltaY)*2 + (deltaX)*(deltaY)*2

	return fmt.Sprint(surface - boxOutterSurface), nil
}

func (d *day18) Exec(input string) (*DayResult, error) {
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
