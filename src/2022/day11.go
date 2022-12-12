package aoc2022

import (
	"aoc2022/utils"
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type day11 struct{}

func init() {
	Days[11] = &day11{}
}

type monkey struct {
	items       utils.Queue[int]
	op          func(int) int
	divisibleBy int
	ifTrue      int
	ifFalse     int
}

func (d *day11) Parse(input string) ([]monkey, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)

	monkeys := make([]monkey, 0)

	for scanner.Scan() {
		line := scanner.Text()

		switch 0 {
		case strings.Index(line, "Monkey "):
			monkeys = append(monkeys, monkey{})

		case strings.Index(line, "  Starting items: "):
			s := strings.Split(line, "  Starting items: ")
			strs := strings.Split(s[1], ", ")
			for _, str := range strs {
				v, err := strconv.Atoi(str)
				if err != nil {
					return nil, err
				}
				monkeys[len(monkeys)-1].items.Enqueue(v)
			}

		case strings.Index(line, "  Operation: new = old "):
			strs := strings.Split(line, "  Operation: new = old ")
			aux := strings.Split(strs[1], " ")
			operator := aux[0]
			operand := aux[1]

			getOperand := func(n int) int {
				if operand == "old" {
					return n
				}
				v, _ := strconv.Atoi(operand)
				return v
			}

			compute := func(i1, i2 int) int {
				if operator == "+" {
					return i1 + i2
				} else {
					return i1 * i2
				}
			}

			monkeys[len(monkeys)-1].op = func(old int) int {
				return compute(old, getOperand(old))
			}

		case strings.Index(line, "  Test: divisible by "):
			strs := strings.Split(line, "  Test: divisible by ")
			v, err := strconv.Atoi(strs[1])
			if err != nil {
				return nil, err
			}
			monkeys[len(monkeys)-1].divisibleBy = v

		case strings.Index(line, "    If true: throw to monkey "):
			strs := strings.Split(line, "    If true: throw to monkey ")
			v, err := strconv.Atoi(strs[1])
			if err != nil {
				return nil, err
			}
			monkeys[len(monkeys)-1].ifTrue = v

		case strings.Index(line, "    If false: throw to monkey "):
			strs := strings.Split(line, "    If false: throw to monkey ")
			v, err := strconv.Atoi(strs[1])
			if err != nil {
				return nil, err
			}
			monkeys[len(monkeys)-1].ifFalse = v
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return monkeys, nil
}

func computeMonkeyBusiness(monkeys []monkey, rounds int, getWorryLvl func(int, monkey) int) int {
	inspectCount := make([]int, len(monkeys))

	for i := 0; i < rounds; i++ {
		// Gotcha! The range expression is evaluated once, before beginning the loop and a copy
		// of the array is used to generate the iteration values.
		for m, monkey := range monkeys {
			for {
				// Because we are mutating items, we need to grab the reference from `monkeys` slice
				// and not from the `monkey` copy.
				item, err := monkeys[m].items.Deque()
				if err != nil {
					break
				}
				inspectCount[m]++
				worryLvl := getWorryLvl(item, monkey)
				throwTo := monkey.ifTrue
				if worryLvl%monkey.divisibleBy != 0 {
					throwTo = monkey.ifFalse
				}
				monkeys[throwTo].items.Enqueue(worryLvl)
			}
		}
	}

	sort.Ints(inspectCount)

	monkeyBusiness := inspectCount[len(inspectCount)-1] * inspectCount[len(inspectCount)-2]

	return monkeyBusiness
}

func (d *day11) Part1(monkeys []monkey) (string, error) {
	rounds := 20

	getWorryLvl := func(item int, m monkey) int { return m.op(item) / 3 }

	monkeyBusiness := computeMonkeyBusiness(monkeys, rounds, getWorryLvl)

	return fmt.Sprint(monkeyBusiness), nil
}

/**
 * Part2 - Given the example:
 * During round 1:
 * monkey 0: All items must go to monkey 3 (FALSE)
 * monkey 1: All items must go to monkey 0 (FALSE)
 * monkey 2: All items must go to monkey 3 (FALSE)
 * ...
 * 💡 Trick is to keep worry level low by computing the minimum common multiple between all monkey
 * tests. That way we maintain the divisibility checks they were gona take and thus the if/else
 * condition.
 */

func (d *day11) Part2(monkeys []monkey) (string, error) {
	rounds := 10_000

	mods := 1
	for _, m := range monkeys {
		mods = m.divisibleBy * mods
	}

	getWorryLvl := func(item int, m monkey) int {
		x := m.op(item)
		x = x % mods
		return x
	}

	monkeyBusiness := computeMonkeyBusiness(monkeys, rounds, getWorryLvl)

	return fmt.Sprint(monkeyBusiness), nil
}

func (d *day11) Exec(input string) (*DayResult, error) {
	parsed, err := d.Parse(input)

	if err != nil {
		return nil, err
	}

	part1, err := d.Part1(parsed)

	if err != nil {
		return nil, err
	}

	parsed2, _ := d.Parse(input)

	part2, err := d.Part2(parsed2)

	if err != nil {
		return nil, err
	}

	result := &DayResult{part1, part2}

	return result, nil
}
