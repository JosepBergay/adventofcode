package aoc2022

import "testing"

const inputD4 = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
`

const expectedD4P1 = "2"

func TestDay4Part1(t *testing.T) {
	day := &day4{}

	parsed, err := day.Parse(inputD4)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part1(parsed)

	if err != nil {
		t.Errorf("[Part1]: %v", err.Error())
		return
	}

	if res != expectedD4P1 {
		t.Errorf("Expected: %v \nBut got: %v", expectedD4P1, res)
	}
}

const expectedD4P2 = ""

func TestDay4Part2(t *testing.T) {
	day := &day4{}

	parsed, err := day.Parse(inputD4)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part2(parsed)

	if err != nil {
		t.Errorf("[Part2]: %v", err.Error())
		return
	}

	if res != expectedD4P2 {
		t.Errorf("Expected: %v \nBut got: %v", expectedD4P2, res)
	}
}
