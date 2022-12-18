package aoc2022

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type day15 struct {
	cave             map[Point]string
	sensorsToBeacons map[Point]Point
	row              int
	minX             int
	maxX             int
	minY             int
	maxY             int
}

func init() {
	Days[15] = &day15{}
}

func (d *day15) addToMap(p Point, mark string) error {
	if mark != "S" && mark != "B" && mark != "#" {
		return fmt.Errorf("%v is not valid mark", mark)
	}

	_, exists := d.cave[p]

	if exists {
		return nil
	}

	if p.x > d.maxX {
		d.maxX = p.x
	} else if p.x < d.minX {
		d.minX = p.x
	}

	if p.y > d.maxY {
		d.maxY = p.y
	} else if p.y < d.minY {
		d.minY = p.y
	}

	d.cave[p] = mark

	return nil
}

func (d *day15) Parse(input string) (string, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))

	re := regexp.MustCompile(`[-]?\d+`)

	d.cave = make(map[Point]string)
	d.sensorsToBeacons = make(map[Point]Point)

	for scanner.Scan() {
		line := scanner.Text()

		bytes := re.FindAll([]byte(line), -1)
		if len(bytes) != 4 {
			return "", errors.New("line has not 4 numbers")
		}

		coords := [4]int{}
		for i, b := range bytes {
			coord, err := strconv.Atoi(string(b))
			if err != nil {
				return "", err
			}
			coords[i] = coord
		}

		sensor := Point{coords[0], coords[1]}
		beacon := Point{coords[2], coords[3]}
		d.addToMap(sensor, "S")
		d.addToMap(beacon, "B")
		d.sensorsToBeacons[sensor] = beacon
	}

	return "", nil
}

// TRBL includes Top, Right, Bottom, Left directions in this order.
var TRBL = [4]Point{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

// Stack Overflow T.T
func (d *day15) walkSpace(curr Point, visited map[Point]bool, outOfBounds func(curr Point) bool) {
	if outOfBounds(curr) {
		return
	}

	if visited[curr] {
		return
	}

	d.addToMap(curr, "#")
	visited[curr] = true

	// Recurse
	for _, dir := range TRBL {
		next := Point{curr.x + dir.x, curr.y + dir.y}
		d.walkSpace(next, visited, outOfBounds)
	}
}

func getManhattanDistance(p1, p2 Point) int {
	x := p1.x - p2.x
	y := p1.y - p2.y
	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func (d *day15) Part1(input string) (string, error) {
	// Update map min/max X
	for sensor, beacon := range d.sensorsToBeacons {
		manhattan := getManhattanDistance(sensor, beacon)

		if sensor.x-manhattan < d.minX {
			d.minX = sensor.x - manhattan
		}

		if sensor.x+manhattan > d.maxX {
			d.maxX = sensor.x + manhattan
		}
	}

	out := 0
	for x := d.minX; x <= d.maxX; x++ {
		p := Point{x, d.row}

		if _, ok := d.cave[p]; ok {
			continue
		}

		for sensor, beacon := range d.sensorsToBeacons {
			if getManhattanDistance(sensor, p) <= getManhattanDistance(sensor, beacon) {
				out++
				break
			}
		}
	}

	return fmt.Sprint(out), nil
}

func (d *day15) Part2(input string) (string, error) {
	for x := 0; x <= d.maxX; x++ {
		for y := 0; y <= d.maxY; y++ {
			curr := Point{x, y}
			var sensor *Point
			sensorToBeaconDist := 0

			for s, b := range d.sensorsToBeacons {
				sensorToCurrDist := getManhattanDistance(s, curr)
				sensorToBeaconDist = getManhattanDistance(s, b)

				if sensorToCurrDist <= sensorToBeaconDist {
					sensor = &s
					break
				}
			}

			if sensor != nil {
				// Move y out of this sensor detection zone.
				aux := math.Abs(float64(curr.x - sensor.x))
				y = y + sensor.y - curr.y + sensorToBeaconDist - int(aux)
				continue
			}

			if _, ok := d.cave[curr]; !ok {
				out := curr.x*4000000 + curr.y
				return fmt.Sprint(out), nil
			}
		}
	}

	return "", errors.New("not found :(")
}

func (d *day15) Exec(input string) (*DayResult, error) {
	parsed, err := d.Parse(input)

	if err != nil {
		return nil, err
	}

	d.row = 2_000_000
	part1, err := d.Part1(parsed)

	if err != nil {
		return nil, err
	}

	d.minX, d.minY = 0, 0
	d.maxX, d.maxY = 4_000_000, 4_000_000
	part2, err := d.Part2(parsed)

	if err != nil {
		return nil, err
	}

	result := &DayResult{part1, part2}

	return result, nil
}
