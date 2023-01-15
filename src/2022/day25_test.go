package aoc2022

import "testing"

const inputD25 = `1=-0-2
12111
2=0=
21
2=01
111
20012
112
1=-1=
1-12
12
1=
122
`

const expectedD25P1 = "2=-1=0"

func TestDay25Part1(t *testing.T) {
	day := &day25{}

	parsed, err := day.Parse(inputD25)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part1(parsed)

	if err != nil {
		t.Errorf("[Part1]: %v", err.Error())
		return
	}

	if res != expectedD25P1 {
		t.Errorf("Expected: %v \nBut got: %v", expectedD25P1, res)
	}
}

const expectedD25P2 = ""

func TestDay25Part2(t *testing.T) {
	day := &day25{}

	parsed, err := day.Parse(inputD25)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part2(parsed)

	if err != nil {
		t.Errorf("[Part2]: %v", err.Error())
		return
	}

	if res != expectedD25P2 {
		t.Errorf("Expected: %v \nBut got: %v", expectedD25P2, res)
	}
}
