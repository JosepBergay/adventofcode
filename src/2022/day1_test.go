package aoc2022

import "testing"

const inputD1 = `1000
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

const expectedD1P1 = "24000"

func TestDay1Part1(t *testing.T) {
	day := &day1{}

	parsed, err := day.Parse(inputD1)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part1(parsed)

	if err != nil {
		t.Errorf("[Part1]: %v", err.Error())
		return
	}

	if res != expectedD1P1 {
		t.Errorf("Expected: %v \nBut got: %v", expectedD1P1, res)
	}
}

const expectedD1P2 = "45000"

func TestDay1Part2(t *testing.T) {
	day := &day1{}

	parsed, err := day.Parse(inputD1)

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
		t.Errorf("Expected: %v \nBut got: %v", expectedD1P2, res)
	}
}
