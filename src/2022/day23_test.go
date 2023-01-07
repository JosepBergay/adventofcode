package aoc2022

import "testing"

const inputD23 = `....#..
..###.#
#...#.#
.#...##
#.###..
##.#.##
.#..#..
`

const expectedD23P1 = "110"

func TestDay23Part1(t *testing.T) {
	day := Days[23].(*day23)

	parsed, err := day.Parse(inputD23)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part1(parsed)

	if err != nil {
		t.Errorf("[Part1]: %v", err.Error())
		return
	}

	if res != expectedD23P1 {
		t.Errorf("Expected: %v \nBut got: %v", expectedD23P1, res)
	}
}

const expectedD23P2 = ""

func TestDay23Part2(t *testing.T) {
	day := Days[23].(*day23)

	parsed, err := day.Parse(inputD23)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part2(parsed)

	if err != nil {
		t.Errorf("[Part2]: %v", err.Error())
		return
	}

	if res != expectedD23P2 {
		t.Errorf("Expected: %v \nBut got: %v", expectedD23P2, res)
	}
}
