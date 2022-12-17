package aoc2022

import "testing"

const inputD14 = `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`

const expectedD14P1 = "24"

func TestDay14Part1(t *testing.T) {
	day := &day14{}

	parsed, err := day.Parse(inputD14)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part1(parsed)

	if err != nil {
		t.Errorf("[Part1]: %v", err.Error())
		return
	}

	if res != expectedD14P1 {
		t.Errorf("Expected: %v \nBut got: %v", expectedD14P1, res)
	}
}

const expectedD14P2 = "93"

func TestDay14Part2(t *testing.T) {
	day := &day14{}

	parsed, err := day.Parse(inputD14)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part2(parsed)

	if err != nil {
		t.Errorf("[Part2]: %v", err.Error())
		return
	}

	if res != expectedD14P2 {
		t.Errorf("Expected: %v \nBut got: %v", expectedD14P2, res)
	}
}
