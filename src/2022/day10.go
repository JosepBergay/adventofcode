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

func executeSignals(input string, duringEachCycle func(cycle, registerX int)) error {
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	registerX := 1
	cycleCount := 0

	pending := int(math.Inf(1))

	for {
		cycleCount++

		duringEachCycle(cycleCount, registerX)

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
			return err
		}

		pending = v
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func (d *day10) Part1(input string) (string, error) {
	sum := 0

	duringEachCycle := func(cycle, registerX int) {
		if isInterestingCycle(cycle) {
			sum += cycle * registerX
		}
	}

	err := executeSignals(input, duringEachCycle)

	if err != nil {
		return "", err
	}

	return fmt.Sprint(sum), nil
}

func (d *day10) Part2(input string) (string, error) {
	// Fancy (and efficient) way of building a string. Could also use += or fmt.Sprintf.
	var crt strings.Builder

	duringEachCycle := func(cycle, registerX int) {
		if cycle > 40*6 {
			// The way executeSignals is written makes duringEachCycle execute once more just
			// before EOF, so we skip this pixel.
			return
		}

		pixelPos := (cycle - 1) % 40
		if pixelPos == 0 {
			crt.WriteString("\n")
		}

		pixel := "."
		// Sprite is 3px wide
		if registerX-1 <= pixelPos && pixelPos <= registerX+1 {
			pixel = "#"
		}

		crt.WriteString(pixel)
	}

	err := executeSignals(input, duringEachCycle)

	if err != nil {
		return "", err
	}

	/** Pixel perfect!
	 * ###..####.####.####.#..#.###..####..##..
	 * #..#.#.......#.#....#.#..#..#.#....#..#.
	 * #..#.###....#..###..##...###..###..#..#.
	 * ###..#.....#...#....#.#..#..#.#....####.
	 * #.#..#....#....#....#.#..#..#.#....#..#.
	 * #..#.#....####.####.#..#.###..#....#..#.
	 */
	return crt.String(), nil
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
