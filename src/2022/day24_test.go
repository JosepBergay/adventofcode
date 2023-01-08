package aoc2022

import "testing"

const inputD24 = `#.######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#
`

const expectedD24P1 = "18"

func TestDay24Part1(t *testing.T) {
	day := &day24{}

	parsed, err := day.Parse(inputD24)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part1(parsed)

	if err != nil {
		t.Errorf("[Part1]: %v", err.Error())
		return
	}

	if res != expectedD24P1 {
		t.Errorf("Expected: %v \nBut got: %v", expectedD24P1, res)
	}
}

const expectedD24P2 = ""

func TestDay24Part2(t *testing.T) {
	day := &day24{}

	parsed, err := day.Parse(inputD24)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part2(parsed)

	if err != nil {
		t.Errorf("[Part2]: %v", err.Error())
		return
	}

	if res != expectedD24P2 {
		t.Errorf("Expected: %v \nBut got: %v", expectedD24P2, res)
	}
}
