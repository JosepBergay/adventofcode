package aoc2022

import (
	"testing"
)

const inputD9 = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2
`

const expectedD9P1 = "13"

func TestDay9Part1(t *testing.T) {
	day := &day9{}

	parsed, err := day.Parse(inputD9)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part1(parsed)

	if err != nil {
		t.Errorf("[Part1]: %v", err.Error())
		return
	}

	if res != expectedD9P1 {
		t.Errorf("Expected: %v \nBut got: %v", expectedD9P1, res)
	}
}

const expectedD9P2 = ""

func TestDay9Part2(t *testing.T) {
	day := &day9{}

	parsed, err := day.Parse(inputD9)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part2(parsed)

	if err != nil {
		t.Errorf("[Part2]: %v", err.Error())
		return
	}

	if res != expectedD9P2 {
		t.Errorf("Expected: %v \nBut got: %v", expectedD9P2, res)
	}
}
