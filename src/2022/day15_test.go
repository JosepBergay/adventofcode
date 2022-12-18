package aoc2022

import "testing"

const inputD15 = `Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3`

const expectedD15P1 = "26"

func TestDay15Part1(t *testing.T) {
	day := &day15{}

	day.row = 10

	parsed, err := day.Parse(inputD15)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part1(parsed)

	if err != nil {
		t.Errorf("[Part1]: %v", err.Error())
		return
	}

	if res != expectedD15P1 {
		t.Errorf("Expected: %v \nBut got: %v", expectedD15P1, res)
	}
}

const expectedD15P2 = "56000011"

func TestDay15Part2(t *testing.T) {
	day := &day15{}

	parsed, err := day.Parse(inputD15)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	day.minX, day.minY = 0, 0
	day.maxX, day.maxY = 20, 20
	res, err := day.Part2(parsed)

	if err != nil {
		t.Errorf("[Part2]: %v", err.Error())
		return
	}

	if res != expectedD15P2 {
		t.Errorf("Expected: %v \nBut got: %v", expectedD15P2, res)
	}
}
