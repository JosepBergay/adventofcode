package aoc2022

import "testing"

const inputD5 = `    [D]     
[N] [C]     
[Z] [M] [P] 
 1   2   3  

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

const expectedD5P1 = "CMZ"

func TestDay5Part1(t *testing.T) {
	day := &day5{}

	parsed, err := day.Parse(inputD5)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part1(parsed)

	if err != nil {
		t.Errorf("[Part1]: %v", err.Error())
		return
	}

	if res != expectedD5P1 {
		t.Errorf("Expected: %v \nBut got: %v", expectedD5P1, res)
	}
}

const expectedD5P2 = ""

func TestDay5Part2(t *testing.T) {
	day := &day5{}

	parsed, err := day.Parse(inputD5)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part2(parsed)

	if err != nil {
		t.Errorf("[Part2]: %v", err.Error())
		return
	}

	if res != expectedD5P2 {
		t.Errorf("Expected: %v \nBut got: %v", expectedD5P2, res)
	}
}
