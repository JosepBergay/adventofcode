package aoc2022

import "testing"

const inputD19 = `Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.
Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.
`

const expectedD19P1 = "33"

func TestDay19Part1(t *testing.T) {
	day := &day19{}

	parsed, err := day.Parse(inputD19)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part1(parsed)

	if err != nil {
		t.Errorf("[Part1]: %v", err.Error())
		return
	}

	if res != expectedD19P1 {
		t.Errorf("Expected: %v \nBut got: %v", expectedD19P1, res)
	}
}

const expectedD19P2 = "3348"

func TestDay19Part2(t *testing.T) {
	day := &day19{}

	parsed, err := day.Parse(inputD19)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part2(parsed)

	if err != nil {
		t.Errorf("[Part2]: %v", err.Error())
		return
	}

	if res != expectedD19P2 {
		t.Errorf("Expected: %v \nBut got: %v", expectedD19P2, res)
	}
}
