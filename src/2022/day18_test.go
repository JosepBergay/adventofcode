package aoc2022

import "testing"

const inputD18 = `2,2,2
1,2,2
3,2,2
2,1,2
2,3,2
2,2,1
2,2,3
2,2,4
2,2,6
1,2,5
3,2,5
2,1,5
2,3,5
`

const expectedD18P1 = "64"

func TestDay18Part1(t *testing.T) {
	day := &day18{}

	parsed, err := day.Parse(inputD18)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part1(parsed)

	if err != nil {
		t.Errorf("[Part1]: %v", err.Error())
		return
	}

	if res != expectedD18P1 {
		t.Errorf("Expected: %v \nBut got: %v", expectedD18P1, res)
	}
}

const expectedD18P2 = "58"

func TestDay18Part2(t *testing.T) {
	day := &day18{}

	parsed, err := day.Parse(inputD18)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part2(parsed)

	if err != nil {
		t.Errorf("[Part2]: %v", err.Error())
		return
	}

	if res != expectedD18P2 {
		t.Errorf("Expected: %v \nBut got: %v", expectedD18P2, res)
	}
}
