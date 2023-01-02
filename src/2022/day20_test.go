package aoc2022

import "testing"

const inputD20 = `1
2
-3
3
-2
0
4
`

const expectedD20P1 = "3"

func TestDay20Part1(t *testing.T) {
	day := &day20{}

	parsed, err := day.Parse(inputD20)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part1(parsed)

	if err != nil {
		t.Errorf("[Part1]: %v", err.Error())
		return
	}

	if res != expectedD20P1 {
		t.Errorf("Expected: %v \nBut got: %v", expectedD20P1, res)
	}
}

const expectedD20P2 = ""

func TestDay20Part2(t *testing.T) {
	day := &day20{}

	parsed, err := day.Parse(inputD20)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part2(parsed)

	if err != nil {
		t.Errorf("[Part2]: %v", err.Error())
		return
	}

	if res != expectedD20P2 {
		t.Errorf("Expected: %v \nBut got: %v", expectedD20P2, res)
	}
}
