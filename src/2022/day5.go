package aoc2022

import (
	"aoc2022/utils"
	"bufio"
	"strconv"
	"strings"
)

type day5 struct{}

func init() {
	Days[5] = &day5{}
}

type instruction struct {
	move, from, to int
}

type day5state struct {
	stacks       []*utils.Stack[string]
	instructions []*instruction
}

func parseInstruction(line string) (*instruction, error) {
	// splitted has the shape of: ["move", "X", "from", "Y", "to", "Z"]
	splitted := strings.Split(line, " ")

	move, err := strconv.Atoi(splitted[1])
	if err != nil {
		return nil, err
	}

	from, err := strconv.Atoi(splitted[3])
	if err != nil {
		return nil, err
	}

	to, err := strconv.Atoi(splitted[5])
	if err != nil {
		return nil, err
	}

	return &instruction{move, from, to}, nil
}

type stackLine struct {
	value   string
	stackId int
}

func parseStackLine(line string) ([]stackLine, error) {
	// Each stack column is 4 runes wide `[X]·`
	// So the second rune in a column is our value.

	out := make([]stackLine, 0)
	for i := 0; i < len(line); i = i + 4 {
		value := string(line[i+1])

		if value == " " {
			// Empty group
			continue
		}

		stackId := int(i / 4)
		out = append(out, stackLine{value, stackId})
	}

	return out, nil
}

func (d *day5) Parse(input string) (*day5state, error) {
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)

	scanner.Split(bufio.ScanLines)

	state := &day5state{}

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 || line[1] == '1' {
			// The offending lines
			continue
		}

		if strings.Index(line, "move") == 0 {
			// It's an instruction let's add it
			inst, err := parseInstruction(line)
			if err != nil {
				return nil, err
			}

			state.instructions = append(state.instructions, inst)

			continue
		}

		// Else it's a stack line and each item should be added to it's stack
		values, err := parseStackLine(line)
		if err != nil {
			return nil, err
		}

		for _, v := range values {
			// Init stacks if needed
			if len(state.stacks) <= v.stackId {
				for i := len(state.stacks); i <= v.stackId; i++ {
					state.stacks = append(state.stacks, &utils.Stack[string]{})
				}
			}

			// Adding it on top, will have to reverse it afterwards :/
			state.stacks[v.stackId].Push(v.value)
		}
	}

	for _, stack := range state.stacks {
		stack.Reverse()
	}

	return state, nil
}

func (d *day5) Part1(input *day5state) (string, error) {
	for _, inst := range input.instructions {
		origin := input.stacks[inst.from-1]
		destination := input.stacks[inst.to-1]

		for i := 0; i < inst.move; i++ {
			value, err := origin.Pop()
			if err != nil {
				return "", err
			}

			destination.Push(value)
		}
	}

	out := make([]string, len(input.stacks))

	for _, stack := range input.stacks {
		value, err := stack.Pop()
		if err != nil {
			return "", err
		}

		out = append(out, value)
	}

	return strings.Join(out, ""), nil
}

func (d *day5) Part2(input *day5state) (string, error) {
	return "TODO", nil
}

func (d *day5) Exec(input string) (*DayResult, error) {
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
