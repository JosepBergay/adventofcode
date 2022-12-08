package aoc2022

import "testing"

const inputD8 = `30373
25512
65332
33549
35390
`

const expectedD8P1 = "21"

func TestDay8Part1(t *testing.T) {
	day := &day8{}

	parsed, err := day.Parse(inputD8)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part1(parsed)

	if err != nil {
		t.Errorf("[Part1]: %v", err.Error())
		return
	}

	if res != expectedD8P1 {
		t.Errorf("Expected: %v \nBut got: %v", expectedD8P1, res)
	}
}

const expectedD8P2 = "8"

func TestDay8Part2(t *testing.T) {
	day := &day8{}

	parsed, err := day.Parse(inputD8)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part2(parsed)

	if err != nil {
		t.Errorf("[Part2]: %v", err.Error())
		return
	}

	if res != expectedD8P2 {
		t.Errorf("Expected: %v \nBut got: %v", expectedD8P2, res)
	}
}
