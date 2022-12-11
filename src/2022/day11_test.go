package aoc2022

import "testing"

const inputD11 = `Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1
`

const expectedD11P1 = "10605"

func TestDay11Part1(t *testing.T) {
	day := &day11{}

	parsed, err := day.Parse(inputD11)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part1(parsed)

	if err != nil {
		t.Errorf("[Part1]: %v", err.Error())
		return
	}

	if res != expectedD11P1 {
		t.Errorf("Expected: %v \nBut got: %v", expectedD11P1, res)
	}
}

const expectedD11P2 = ""

func TestDay11Part2(t *testing.T) {
	day := &day11{}

	parsed, err := day.Parse(inputD11)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part2(parsed)

	if err != nil {
		t.Errorf("[Part2]: %v", err.Error())
		return
	}

	if res != expectedD11P2 {
		t.Errorf("Expected: %v \nBut got: %v", expectedD11P2, res)
	}
}
