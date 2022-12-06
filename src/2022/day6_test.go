package aoc2022

import (
	"fmt"
	"testing"
)

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
			"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", "11",
		},
	}
	day := &day6{}

	for i, tC := range testCases {
		t.Run(fmt.Sprintf("Testing P1 input %v", i), func(t *testing.T) {
			if res, err := day.Part1(tC.input); err != nil || res != tC.want {
				t.Errorf("[Part1]: expected %v, <nil> but got %v, %v", tC.want, res, err)
			}
		})
	}
}

func TestDay6Part2(t *testing.T) {
	testCases := []struct {
		input, want string
	}{
		{
			"mjqjpqmgbljsphdztnvjfqwrcgsmlb", "19",
		},
		{
			"bvwbjplbgvbhsrlpgdmjqwftvncz", "23",
		},
		{
			"nppdvjthqldpwncqszvftbrmjlhg", "23",
		},
		{
			"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", "29",
		},
		{
			"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", "26",
		},
	}
	day := &day6{}

	for i, tC := range testCases {
		t.Run(fmt.Sprintf("Testing P1 input %v", i), func(t *testing.T) {
			if res, err := day.Part2(tC.input); err != nil || res != tC.want {
				t.Errorf("[Part1]: expected %v, <nil> but got %v, %v", tC.want, res, err)
			}
		})
	}
}
