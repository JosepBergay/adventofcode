package aoc2022

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type monkeyMap struct {
	points map[Point]rune
	// leftmost and rightmost point for each line
	leftRight [][2]*Point
	// top and bottom point for each column
	topBot [][2]*Point
}

type movement struct {
	move int
	face string
}

type day22 struct {
	monkeyMap    monkeyMap
	path         string
	start        *Point
	instructions []movement
}

func init() {
	Days[22] = &day22{
		monkeyMap: monkeyMap{
			points:    make(map[Point]rune),
			leftRight: make([][2]*Point, 0),
			topBot:    make([][2]*Point, 0),
		},
		instructions: make([]movement, 0),
	}
}

func (d *day22) updateTopBot(p *Point) {
	for x := len(d.monkeyMap.topBot); x <= p.x; x++ {
		d.monkeyMap.topBot = append(d.monkeyMap.topBot, [2]*Point{nil, nil})
	}

	if d.monkeyMap.topBot[p.x][0] == nil {
		d.monkeyMap.topBot[p.x][0] = p
	}

	d.monkeyMap.topBot[p.x][1] = p
}

var pathRe = regexp.MustCompile(`\d+[A-Z]{0,1}`)

func (d *day22) Parse(input string) (string, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))

	scanner.Split(bufio.ScanLines)

	var x, y int
	var nextLineIsPath bool

	for scanner.Scan() {
		line := scanner.Text()

		var min, max *Point

		if line == "" {
			nextLineIsPath = true
			continue
		}
		if nextLineIsPath {
			d.path = line
			matches := pathRe.FindAllString(line, -1)
			for _, s := range matches {
				i := len(s) - 1
				if s[len(s)-1] != 'R' && s[len(s)-1] != 'L' {
					i = len(s)
				}
				move, err := strconv.Atoi(s[:i])
				if err != nil {
					fmt.Println(s)
					return "", err
				}
				d.instructions = append(d.instructions, movement{
					move, s[i:],
				})
			}
			break
		}

		for _, v := range line {
			switch v {
			case ' ':
			default: // '.' || '#'
				p := &Point{x, y}

				if d.start == nil {
					d.start = p
				}

				if min == nil {
					min = p
				}
				max = p

				d.updateTopBot(p)

				d.monkeyMap.points[*p] = v
			}

			x++
		}

		d.monkeyMap.leftRight = append(d.monkeyMap.leftRight, [2]*Point{min, max})
		x = 0
		y++
	}

	return "", nil
}

func (d *day22) Part1(input string) (string, error) {
	allFaces := [4]rune{'>', 'v', '<', '^'}

	// State
	currPoint := *d.start
	currFacing := 0 // Index of `allFaces`
	currInstIdx := 0

	getNext := func() Point {
		n := Point{}
		switch allFaces[currFacing] {
		case '>':
			n = Point{currPoint.x + 1, currPoint.y}
		case 'v':
			n = Point{currPoint.x, currPoint.y + 1}
		case '<':
			n = Point{currPoint.x - 1, currPoint.y}
		default: // case '^':
			n = Point{currPoint.x, currPoint.y - 1}
		}

		_, ok := d.monkeyMap.points[n]
		if !ok {
			// Wrap around
			switch allFaces[currFacing] {
			case '>':
				n = *d.monkeyMap.leftRight[currPoint.y][0]
			case 'v':
				n = *d.monkeyMap.topBot[currPoint.x][0]
			case '<':
				n = *d.monkeyMap.leftRight[currPoint.y][1]
			default: // case '^':
				n = *d.monkeyMap.topBot[currPoint.x][1]
			}
		}

		return n
	}

	rotate := func() {
		switch d.instructions[currInstIdx].face {
		case "R":
			currFacing = (currFacing + 1) % 4
		case "L":
			currFacing = (currFacing - 1 + 4) % 4
		}
	}

	for currInstIdx < len(d.instructions) {
		instr := d.instructions[currInstIdx]

		// Move to tile
		for i := 0; i < instr.move; i++ {
			// Compute next
			next := getNext()

			if d.monkeyMap.points[next] == '#' {
				// Hit a wall so stop.
				break
			}

			// Tile is '.' so move there.
			currPoint = next
		}

		// Turn left or right
		rotate()

		// Update state
		currInstIdx++
	}

	// Compute password
	pass := 1000*(currPoint.y+1) + 4*(currPoint.x+1) + currFacing

	return fmt.Sprint(pass), nil
}

func (d *day22) Part2(input string) (string, error) {
	return "TODO", nil
}

func (d *day22) Exec(input string) (*DayResult, error) {
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
