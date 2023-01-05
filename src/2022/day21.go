package aoc2022

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type monkeyNumber struct {
	name     string
	value    *int
	monkey1  string
	monkey2  string
	operator string
}

type day21 struct {
	m map[string]*monkeyNumber
}

const rootMonkey = "root"

func init() {
	Days[21] = &day21{}
}

func (d *day21) Parse(input string) (string, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)

	m := make(map[string]*monkeyNumber)

	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, ": ")

		monkey := monkeyNumber{
			name: s[0],
		}

		m[monkey.name] = &monkey

		value, err := strconv.Atoi(s[1])
		if err == nil { // if no error
			monkey.value = &value
			continue
		}

		s = strings.Split(s[1], " ")

		monkey.monkey1 = s[0]
		monkey.operator = s[1]
		monkey.monkey2 = s[2]
	}

	d.m = m

	return "", nil
}

func (d *day21) findMonkeyNumber(name string) int {
	monkey := d.m[name]
	if monkey.value != nil {
		return *monkey.value
	}

	m1 := d.findMonkeyNumber(monkey.monkey1)
	m2 := d.findMonkeyNumber(monkey.monkey2)

	var v int
	switch monkey.operator {
	case "+":
		v = m1 + m2
	case "-":
		v = m1 - m2
	case "*":
		v = m1 * m2
	case "/":
		v = m1 / m2
	}

	monkey.value = &v
	return v
}

func (d *day21) Part1(input string) (string, error) {
	n := d.findMonkeyNumber("root")

	return fmt.Sprint(n), nil
}

func (d *day21) Part2(input string) (string, error) {
	return "TODO", nil
}

func (d *day21) Exec(input string) (*DayResult, error) {
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
