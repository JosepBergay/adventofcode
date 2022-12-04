package aoc2022

import "testing"

const inputD3 = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`

const expectedD3P1 = "157"

func TestDay3Part1(t *testing.T) {
	day := &day3{}

	parsed, err := day.Parse(inputD3)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part1(parsed)

	if err != nil {
		t.Errorf("[Part1]: %v", err.Error())
		return
	}

	if res != expectedD3P1 {
		t.Errorf("Expected: %v \nBut got: %v", expectedD3P1, res)
	}
}

const expectedD3P2 = "70"

func TestDay3Part2(t *testing.T) {
	day := &day3{}

	parsed, err := day.Parse(inputD3)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part2(parsed)

	if err != nil {
		t.Errorf("[Part2]: %v", err.Error())
		return
	}

	if res != expectedD3P2 {
		t.Errorf("Expected: %v \nBut got: %v", expectedD3P2, res)
	}
}
