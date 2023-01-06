package aoc2022

import "testing"

const inputD22 = `        ...#
        .#..
        #...
        ....
...#.......#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.

10R5L5R10L4R5L5
`

const expectedD22P1 = "6032"

func TestDay22Part1(t *testing.T) {
	day := Days[22].(*day22)

	parsed, err := day.Parse(inputD22)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part1(parsed)

	if err != nil {
		t.Errorf("[Part1]: %v", err.Error())
		return
	}

	if res != expectedD22P1 {
		t.Errorf("Expected: %v \nBut got: %v", expectedD22P1, res)
	}
}

const expectedD22P2 = "5031"

func TestDay22Part2(t *testing.T) {
	day := Days[22].(*day22)

	parsed, err := day.Parse(inputD22)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part2(parsed)

	if err != nil {
		t.Errorf("[Part2]: %v", err.Error())
		return
	}

	if res != expectedD22P2 {
		t.Errorf("Expected: %v \nBut got: %v", expectedD22P2, res)
	}
}
