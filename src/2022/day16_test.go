package aoc2022

import "testing"

const inputD16 = `Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
Valve BB has flow rate=13; tunnels lead to valves CC, AA
Valve CC has flow rate=2; tunnels lead to valves DD, BB
Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
Valve EE has flow rate=3; tunnels lead to valves FF, DD
Valve FF has flow rate=0; tunnels lead to valves EE, GG
Valve GG has flow rate=0; tunnels lead to valves FF, HH
Valve HH has flow rate=22; tunnel leads to valve GG
Valve II has flow rate=0; tunnels lead to valves AA, JJ
Valve JJ has flow rate=21; tunnel leads to valve II
`

const expectedD16P1 = "1651"

func TestDay16Part1(t *testing.T) {
	day := &day16{}

	parsed, err := day.Parse(inputD16)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part1(parsed)

	if err != nil {
		t.Errorf("[Part1]: %v", err.Error())
		return
	}

	if res != expectedD16P1 {
		t.Errorf("Expected: %v \nBut got: %v", expectedD16P1, res)
	}
}

const expectedD16P2 = ""

func TestDay16Part2(t *testing.T) {
	day := &day16{}

	parsed, err := day.Parse(inputD16)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part2(parsed)

	if err != nil {
		t.Errorf("[Part2]: %v", err.Error())
		return
	}

	if res != expectedD16P2 {
		t.Errorf("Expected: %v \nBut got: %v", expectedD16P2, res)
	}
}
