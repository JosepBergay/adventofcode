package aoc2022

import (
	"fmt"
	"strings"
)

type day6 struct{}

func init() {
	Days[6] = &day6{}
}

func (d *day6) Parse(input string) (string, error) {
	return input, nil
}

// isMarker returns true if candidate has no repeated values
func isMarker(candidate []byte) bool {
	m := make(map[byte]any)

	for _, v := range candidate {
		m[v] = true
	}

	return len(m) == len(candidate)
}

func findMarker(input string, markerLength int) (string, error) {
	reader := strings.NewReader(input)

	msg := make([]byte, markerLength)

	i := 0
	for {
		_, err := reader.ReadAt(msg, int64(i))
		if err != nil {
			return "", err
		}

		if isMarker(msg) {
			return fmt.Sprint(i + markerLength), nil
		}

		i++
	}
}

func (d *day6) Part1(input string) (string, error) {
	return findMarker(input, 4)
}

func (d *day6) Part2(input string) (string, error) {
	return findMarker(input, 14)
}

func (d *day6) Exec(input string) (*DayResult, error) {
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
