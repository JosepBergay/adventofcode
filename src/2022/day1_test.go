package aoc2022

import "testing"

const input = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`

const expectedP1 = "24000"

func TestDay1Part1(t *testing.T) {
	day := &day{}

	parsed, err := day.Parse(input)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part1(parsed)

	if err != nil {
		t.Errorf("[Part1]: %v", err.Error())
		return
	}

	if res != expectedP1 {
		t.Errorf("Expected: %v \nBut got: %v", expectedP1, res)
	}
}

const expectedP2 = "45000"

func TestDay1Part2(t *testing.T) {
	day := &day{}

	parsed, err := day.Parse(input)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part2(parsed)

	if err != nil {
		t.Errorf("[Part2]: %v", err.Error())
		return
	}

	if res != expectedP2 {
		t.Errorf("Expected: %v \nBut got: %v", expectedP2, res)
	}
}
