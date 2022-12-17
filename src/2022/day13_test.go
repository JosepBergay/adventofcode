package aoc2022

import "testing"

const inputD13 = `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`

const expectedD13P1 = "13"

func TestDay13Part1(t *testing.T) {
	day := &day13{}

	parsed, err := day.Parse(inputD13)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part1(parsed)

	if err != nil {
		t.Errorf("[Part1]: %v", err.Error())
		return
	}

	if res != expectedD13P1 {
		t.Errorf("Expected: %v \nBut got: %v", expectedD13P1, res)
	}
}

const expectedD13P2 = ""

func TestDay13Part2(t *testing.T) {
	day := &day13{}

	parsed, err := day.Parse(inputD13)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part2(parsed)

	if err != nil {
		t.Errorf("[Part2]: %v", err.Error())
		return
	}

	if res != expectedD13P2 {
		t.Errorf("Expected: %v \nBut got: %v", expectedD13P2, res)
	}
}
