package aoc2022

import "testing"

const inputD21 = `root: pppw + sjmn
dbpl: 5
cczh: sllz + lgvd
zczc: 2
ptdq: humn - dvpt
dvpt: 3
lfqf: 4
humn: 5
ljgn: 2
sjmn: drzm * dbpl
sllz: 4
pppw: cczh / lfqf
lgvd: ljgn * ptdq
drzm: hmdt - zczc
hmdt: 32
`

const expectedD21P1 = "152"

func TestDay21Part1(t *testing.T) {
	day := &day21{}

	parsed, err := day.Parse(inputD21)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part1(parsed)

	if err != nil {
		t.Errorf("[Part1]: %v", err.Error())
		return
	}

	if res != expectedD21P1 {
		t.Errorf("Expected: %v \nBut got: %v", expectedD21P1, res)
	}
}

const expectedD21P2 = "301"

func TestDay21Part2(t *testing.T) {
	day := &day21{}

	parsed, err := day.Parse(inputD21)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part2(parsed)

	if err != nil {
		t.Errorf("[Part2]: %v", err.Error())
		return
	}

	if res != expectedD21P2 {
		t.Errorf("Expected: %v \nBut got: %v", expectedD21P2, res)
	}
}
