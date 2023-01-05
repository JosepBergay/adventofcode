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

func (m monkeyNumber) String() string {
	if m.value != nil {
		return fmt.Sprintf("%v: %v,", m.name, *m.value)
	}
	return fmt.Sprintf("%v: %v %v %v,", m.name, m.monkey1, m.operator, m.monkey2)
}

type day21 struct {
	m map[string]*monkeyNumber
}

const rootMonkeyName = "root"

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
	n := d.findMonkeyNumber(rootMonkeyName)

	return fmt.Sprint(n), nil
}

func (d *day21) Part2(input string) (string, error) {
	var getHumanBranch func(*monkeyNumber) []*monkeyNumber

	getHumanBranch = func(monkey *monkeyNumber) []*monkeyNumber {
		if monkey.name == "humn" {
			// Found ourselves!
			branch := make([]*monkeyNumber, 1)
			branch[0] = monkey
			return branch
		}

		// monkey.value may be filled if we run Part1 first.
		if monkey.monkey1 == "" && monkey.monkey2 == "" {
			// Dead end
			return nil
		}

		branchNames := [2]string{monkey.monkey1, monkey.monkey2}
		for _, branchName := range branchNames {
			if branch := getHumanBranch(d.m[branchName]); branch != nil {
				// Found branch
				// humanBranch = append(humanBranch, *input[branchName])
				return append(branch, monkey)
			}
		}

		return nil
	}

	rootMonkey := d.m[rootMonkeyName]

	rootMonkey.operator = "="

	humanBranch := getHumanBranch(rootMonkey)

	acc := 0

	// Human branch starts with `humn` and ends with `root` so we loop in reverse order.
	for i := len(humanBranch) - 1; i > 0; i-- {
		curr := humanBranch[i]
		branches := [2]string{curr.monkey1, curr.monkey2}

		// Get the branch that does not contain human.
		for j, b := range branches {
			if humanBranch[i-1].name != b {
				// Compute number for the other branche.
				num := d.findMonkeyNumber(b)

				// Isolate x
				switch curr.operator {
				case "=":
					acc = num
				case "+":
					acc = acc - num
				case "-":
					if j == 0 {
						// num - x = curr
						acc = num - acc
					} else {
						// x - num = curr
						acc = acc + num
					}
				case "*":
					acc = acc / num
				case "/":
					if j == 0 {
						// num / x = curr
						acc = num / acc
					} else {
						// x / num = curr
						acc = acc * num
					}
				}
			}
		}
	}

	return fmt.Sprint(acc), nil
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
