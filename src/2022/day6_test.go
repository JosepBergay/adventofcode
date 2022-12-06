package aoc2022

import (
	"fmt"
	"testing"
)

const inputD6 = `mjqjpqmgbljsphdztnvjfqwrcgsmlb
`

const expectedD6P1 = "7"

func TestDay6Part1(t *testing.T) {
	testCases := []struct {
		input, want string
	}{
		{
			"mjqjpqmgbljsphdztnvjfqwrcgsmlb", "7",
		},
		{
			"bvwbjplbgvbhsrlpgdmjqwftvncz", "5",
		},
		{
			"nppdvjthqldpwncqszvftbrmjlhg", "6",
		},
		{
			"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", "10",
		},
		{
			"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", "12",
		},
	}
	day := &day6{}

	for i, tC := range testCases {
		t.Run(fmt.Sprintf("Testing P1 input %v", i), func(t *testing.T) {
			if res, err := day.Part1(tC.input); err != nil || res != tC.want {
				t.Errorf("[Part1]: expected %v, nil but got %v, %v", tC.want, res, err)
			}
		})
	}
}

const expectedD6P2 = ""

func TestDay6Part2(t *testing.T) {
	day := &day6{}

	parsed, err := day.Parse(inputD6)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part2(parsed)

	if err != nil {
		t.Errorf("[Part2]: %v", err.Error())
		return
	}

	if res != expectedD6P2 {
		t.Errorf("Expected: %v \nBut got: %v", expectedD6P2, res)
	}
}
