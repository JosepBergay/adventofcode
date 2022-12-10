package aoc2022

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type day10 struct{}

func init() {
	Days[10] = &day10{}
}

func (d *day10) Parse(input string) (string, error) {
	return input, nil
}

func isInterestingCycle(cycle int) bool {
	interestingCycles := [...]int{20, 60, 100, 140, 180, 220}

	for _, v := range interestingCycles {
		if v == cycle {
			return true
		}
	}

	return false
}

func (d *day10) Part1(input string) (string, error) {
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	sum := 0

	registerX := 1
	cycleCount := 0

	pending := int(math.Inf(1))

	for {
		cycleCount++

		if isInterestingCycle(cycleCount) {
			sum += cycleCount * registerX // Compute signal strength *DURING* cycle
		}

		// We fake a cycle when we have a pending value from previous cycle.
		if pending != int(math.Inf(1)) {
			registerX += pending
			// Reset pending
			pending = int(math.Inf(1))
			continue
		}

		hasMore := scanner.Scan()
		if !hasMore {
			break
		}

		instr := scanner.Text()

		if instr == "noop" {
			continue
		}

		strs := strings.Split(instr, " ")
		v, err := strconv.Atoi(strs[1])
		if err != nil {
			return "", err
		}

		pending = v
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return fmt.Sprint(sum), nil
}

func (d *day10) Part2(input string) (string, error) {
	return "TODO", nil
}

func (d *day10) Exec(input string) (*DayResult, error) {
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
