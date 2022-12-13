package aoc2022

import "testing"

const inputD12 = `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi
`

const expectedD12P1 = "31"

func TestDay12Part1(t *testing.T) {
	day := &day12{}

	parsed, err := day.Parse(inputD12)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part1(parsed)

	if err != nil {
		t.Errorf("[Part1]: %v", err.Error())
		return
	}

	if res != expectedD12P1 {
		t.Errorf("Expected: %v \nBut got: %v", expectedD12P1, res)
	}
}

const expectedD12P2 = ""

func TestDay12Part2(t *testing.T) {
	day := &day12{}

	parsed, err := day.Parse(inputD12)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part2(parsed)

	if err != nil {
		t.Errorf("[Part2]: %v", err.Error())
		return
	}

	if res != expectedD12P2 {
		t.Errorf("Expected: %v \nBut got: %v", expectedD12P2, res)
	}
}
