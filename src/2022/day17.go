package aoc2022

import (
	"fmt"
	"strings"
)

type Chamber [][7]int

type day17 struct {
	rocks [5]Chamber
}

func init() {
	Days[17] = &day17{}
}

func (d *day17) Parse(input string) (string, error) {
	d.rocks = [5]Chamber{
		{
			// ..####
			{0, 0, 1, 1, 1, 1},
		}, {
			// ...#.
			// ..###
			// ...#.
			{0, 0, 0, 1},
			{0, 0, 1, 1, 1},
			{0, 0, 0, 1},
		}, {
			// ....#
			// ....#
			// ..###
			{0, 0, 0, 0, 1},
			{0, 0, 0, 0, 1},
			{0, 0, 1, 1, 1},
		}, {
			// ..#
			// ..#
			// ..#
			// ..#
			{0, 0, 1},
			{0, 0, 1},
			{0, 0, 1},
			{0, 0, 1},
		}, {
			// ..##
			// ..##
			{0, 0, 1, 1},
			{0, 0, 1, 1},
		},
	}

	return strings.ReplaceAll(input, "\n", ""), nil
}

// func (r Chamber) String() string {
// 	out := ""
// 	for i := 0; i < len(r); i++ {
// 		out += fmt.Sprint(r[i], "\n")
// 	}
// 	return out
// }

func moveRockLeft(r Chamber) Chamber {
	// Check
	for i := 0; i < len(r); i++ {
		if r[i][0] == 1 {
			// Can't go left. Return rock as is.
			return r
		}
	}

	// Move
	moved := make(Chamber, len(r))
	for i := 0; i < len(r); i++ {
		moved[i] = [7]int{}
		for j := 0; j < 6; j++ {
			if r[i][j+1] == 1 {
				moved[i][j] = 1
			}
		}
	}

	return moved
}

func moveRockRight(r Chamber) Chamber {
	// Check
	for i := 0; i < len(r); i++ {
		if r[i][6] == 1 {
			// Can't go right. Return rock as is.
			return r
		}
	}

	// Move
	moved := make(Chamber, len(r))
	for i := 0; i < len(r); i++ {
		moved[i] = [7]int{}
		for j := 6; j > 0; j-- {
			if r[i][j-1] == 1 {
				moved[i][j] = 1
			}
		}
	}

	return moved
}

func moveRockHorizontally(r Chamber, getJet func() byte) Chamber {
	jet := getJet()
	if jet == '>' {
		return moveRockRight(r)
	} else {
		return moveRockLeft(r)
	}
}

func addLines(l1, l2 [7]int) (*[7]int, error) {
	out := [7]int{}
	for i := 0; i < 7; i++ {
		if l1[i]+l2[i] > 1 {
			return nil, fmt.Errorf("overlap")
		}
		out[i] = l1[i] + l2[i]
	}
	return &out, nil
}

func moveRock(rock Chamber, chamber *Chamber, getJet func() byte) {
	// Each rock appears 3 units above the highest rock. We can skip downward movement for now.
	for i := 0; i < 3; i++ {
		rock = moveRockHorizontally(rock, getJet)
	}

	// Now we are just on top of the highest rock. Start cycle.
	l := len(*chamber)
	for {
		next := moveRockHorizontally(rock, getJet)
		// Check overlap
		overlap := false
		for i := len(next) - 1; i >= 0; i-- {
			j := len(rock) - 1 - i + l
			if j >= len(*chamber) {
				break
			}
			_, err := addLines(next[i], (*chamber)[j])
			if err != nil {
				overlap = true
				break
			}
		}

		// If there's no horizontal overlap update r.
		if !overlap {
			rock = next
		}

		overlap = false
		for i := len(rock) - 1; i >= 0; i-- {
			j := len(rock) - 1 - i + l - 1
			if j < 0 || j >= len(*chamber) {
				break
			}
			_, err := addLines(rock[i], (*chamber)[j])
			if err != nil {
				overlap = true
				break
			}
		}

		if !overlap && l > 0 {
			l--
			continue
		}

		// If there's overlap going down, rock comes to rest.
		// Add rock in reverse so it stacks correctly.
		for i := len(rock) - 1; i >= 0; i-- {
			j := len(rock) - 1 - i + l
			if j < 0 || j >= len(*chamber) {
				*chamber = append(*chamber, rock[i])
			} else {
				line, _ := addLines((*chamber)[j], rock[i])
				(*chamber)[j] = *line
			}
		}
		break
	}
}

func (d *day17) runChamberSimulation(input string, rockCount int) int {
	chamber := make(Chamber, 0)

	jetIdx := 0
	getJet := func() (jet byte) {
		jet = input[jetIdx%len(input)]
		jetIdx++
		return
	}

	// 2282 was found empirically. Started checking cycles well into the cycle (10_000) then reduced
	// that number. Turns out there are some (minor) cycles in the first part of the loop.
	magicNumber := 2282 // 10_000

	// [jetIdx, rockModulo]: [height, rockIdx]
	cache := make(map[[2]int][2]int)

	// Tracking virtually added height
	h := 0

	// For each rock
	for rockIdx := 0; rockIdx < rockCount; rockIdx++ {
		rock := d.rocks[rockIdx%5]

		moveRock(rock, &chamber, getJet)

		hash := [2]int{jetIdx % len(input), rockIdx % len(d.rocks)}
		if h == 0 && rockIdx > magicNumber {
			if v, ok := cache[hash]; ok {
				// Found a cycle !?
				cycleHeight := len(chamber) - v[0]
				cycleLength := rockIdx - v[1]
				cyclesLeft := (rockCount - rockIdx) / cycleLength
				h += cyclesLeft * cycleHeight
				rocksLeft := (rockCount - rockIdx) % cycleLength
				// Skip looping until last partial cycle.
				rockIdx = rockCount - rocksLeft
			}
		}

		cache[hash] = [2]int{len(chamber), rockIdx}
	}

	return len(chamber) + h
}

// Len 77-25=52;
func (d *day17) Part1(input string) (string, error) {
	rockCount := 2022

	height := d.runChamberSimulation(input, rockCount)

	return fmt.Sprint(height), nil
}

func (d *day17) Part2(input string) (string, error) {
	rockCount := 1_000_000_000_000

	height := d.runChamberSimulation(input, rockCount)

	return fmt.Sprint(height), nil
}

func (d *day17) Exec(input string) (*DayResult, error) {
	parsed, err := d.Parse(input)

	if err != nil {
		return nil, err
	}

	part1, err := d.Part1(parsed)

	if err != nil {
		return nil, err
	}

	part2, err := d.Part2(parsed)

	if err != nil {
		return nil, err
	}

	result := &DayResult{part1, part2}

	return result, nil
}
