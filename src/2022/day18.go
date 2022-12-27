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
}

func init() {
	Days[18] = &day18{}
}

func (d *day18) addCube(x, y, z int) {
	if _, ok := d.matrix[x]; !ok {
		d.matrix[x] = make(map[int]map[int]int, 0)
	}

	if _, ok := d.matrix[x][y]; !ok {
		d.matrix[x][y] = make(map[int]int, 0)
	}

	d.matrix[x][y][z] = 1
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
	return "TODO", nil
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
