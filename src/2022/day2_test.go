package aoc2022

import "testing"

const inputD2 = `A Y
B X
C Z
`

const expectedD2P1 = "15"

func TestDay2Part1(t *testing.T) {
	day := &day2{}

	parsed, err := day.Parse(inputD2)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part1(parsed)

	if err != nil {
		t.Errorf("[Part1]: %v", err.Error())
		return
	}

	if res != expectedD2P1 {
		t.Errorf("Expected: %v \nBut got: %v", expectedD2P1, res)
	}
}

const expectedD2P2 = ""

func TestDay2Part2(t *testing.T) {
	day := &day2{}

	parsed, err := day.Parse(inputD2)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part2(parsed)

	if err != nil {
		t.Errorf("[Part2]: %v", err.Error())
		return
	}

	if res != expectedD1P2 {
		t.Errorf("Expected: %v \nBut got: %v", expectedD2P2, res)
	}
}
