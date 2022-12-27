package aoc2022

import "testing"

const inputD17 = `>>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>
`

const expectedD17P1 = "3068"

func TestAddLines(t *testing.T) {
	l1 := [7]int{0, 0, 0, 1, 0, 0, 0}
	l2 := [7]int{0, 0, 0, 0, 1, 1, 1}
	l, err := addLines(l1, l2)
	if err != nil {
		t.Error("Error adding lines", err)
		return
	}
	expect := [7]int{0, 0, 0, 1, 1, 1, 1}
	if *l != expect {
		t.Errorf("lines not equal -> Expected: %v; But got %v", expect, *l)
	}
}

func TestDay17Part1(t *testing.T) {
	day := &day17{}

	parsed, err := day.Parse(inputD17)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part1(parsed)

	if err != nil {
		t.Errorf("[Part1]: %v", err.Error())
		return
	}

	if res != expectedD17P1 {
		t.Errorf("Expected: %v \nBut got: %v", expectedD17P1, res)
	}
}

const expectedD17P2 = ""

func TestDay17Part2(t *testing.T) {
	day := &day17{}

	parsed, err := day.Parse(inputD17)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part2(parsed)

	if err != nil {
		t.Errorf("[Part2]: %v", err.Error())
		return
	}

	if res != expectedD17P2 {
		t.Errorf("Expected: %v \nBut got: %v", expectedD17P2, res)
	}
}
