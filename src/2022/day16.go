package aoc2022

import (
	"aoc2022/utils"
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type valve struct {
	id        string
	flowRate  int
	connected *[]string
}

type day16 struct {
	start     string
	valves    map[string]valve
	opened    map[string]bool
	maxFlow   int
	distances map[string]map[string]int
	all       []valve
}

func init() {
	Days[16] = &day16{}
}

var valvesRe = regexp.MustCompile(`[A-Z]{2}`)
var flowRateRe = regexp.MustCompile(`\d+`)

func (d *day16) Parse(input string) (string, error) {
	d.valves = make(map[string]valve)
	d.start = "AA"
	d.opened = make(map[string]bool)
	d.distances = make(map[string]map[string]int) // {"AA": {"DD": 1, "II": 1, ...}}
	d.all = make([]valve, 0)

	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()

		valves := valvesRe.FindAllString(line, -1)
		flowRateStr := flowRateRe.FindString(line)

		flowRate, err := strconv.Atoi(flowRateStr)
		if err != nil {
			return "", err
		}

		connected := valves[1:]
		v := valve{
			id:        valves[0],
			flowRate:  flowRate,
			connected: &connected,
		}

		if flowRate > d.maxFlow {
			d.maxFlow = flowRate
		}

		d.distances[v.id] = make(map[string]int)
		for _, c := range *v.connected {
			d.distances[v.id][c] = 1
		}

		d.valves[v.id] = v
		d.all = append(d.all, v)
	}

	if _, found := d.valves[d.start]; !found {
		return "", fmt.Errorf("starting valve not found :s")
	}

	return "", nil
}

func computeEventualTotalPressure(flowRate, distance, minutesLeft int) int {
	return flowRate * (minutesLeft - distance - 1) // -1 for opening
}

func (d *day16) getValvesWithFlowAndNotOpened() []valve {
	out := make([]valve, 0)
	for _, v := range d.valves {
		if _, found := d.opened[v.id]; !found && v.flowRate > 0 {
			out = append(out, v)
		}
	}
	return out
}

func (d *day16) findShortestDistanceBetweenValves(from, to valve) int {
	if d, ok := d.distances[from.id][to.id]; ok {
		return d
	}

	// find distance
	p := utils.FindShortestPath(from, to, d.all, func(curr valve) []valve {
		adjacents := make([]valve, 0)
		for _, v := range *curr.connected {
			adjacents = append(adjacents, d.valves[v])
		}
		return adjacents
	}, func(curr, adjacent valve) int { return 1 })

	// Now we know all distances from 'from' to the rest of valves.
	for i, v := range p {
		// update map for next time. Index is the distance because all paths weight 1
		d.distances[from.id][v.id] = i
	}

	return len(p) - 1
}

var allPathsMap = make(map[string][]valve)

// Used to create map keys
func (p valve) String() string {
	return p.id
}

func (d *day16) getPathPressure(path []valve, maxMinutes int) int {
	start := d.valves[d.start]
	total := 0

	curr := start
	pos := 0
	for i := maxMinutes; i > 0; i-- {
		if pos >= len(path)-1 {
			break
		}
		next := path[pos+1]
		dist := d.findShortestDistanceBetweenValves(curr, next)
		pres := computeEventualTotalPressure(next.flowRate, dist, i)

		i -= dist // - 1 will be done in for loop T.T
		total += pres
		curr = next
		pos++
	}

	return total
}

func (d *day16) generateAllPaths(path []valve, unseen []valve, distance, maxMinutes int) {
	if len(unseen) == 0 {
		allPathsMap[fmt.Sprint(path)] = path
		return
	}

	last := path[len(path)-1]

	for _, v := range unseen {
		dist := d.findShortestDistanceBetweenValves(last, v)

		if dist+1+distance >= maxMinutes {
			allPathsMap[fmt.Sprint(path)] = path
			continue
		}

		newUnseen := make([]valve, 0)
		for _, u := range unseen {
			if u.id != v.id {
				newUnseen = append(newUnseen, u)
			}
		}
		copiedPath := make([]valve, len(path))
		copy(copiedPath, path)
		copiedPath = append(copiedPath, v)
		d.generateAllPaths(copiedPath, newUnseen, dist+1+distance, maxMinutes)
	}
}

func (d *day16) Part1(input string) (string, error) {
	startingPath := make([]valve, 0)
	startingPath = append(startingPath, d.valves[d.start])
	maxMinutes := 30
	d.generateAllPaths(startingPath, d.getValvesWithFlowAndNotOpened(), 0, maxMinutes)

	out := 0
	for _, path := range allPathsMap {
		if p := d.getPathPressure(path, maxMinutes); p > out {
			out = p
		}
	}

	return fmt.Sprint(out), nil
}

type pathPressure struct {
	path     []valve
	pressure int
}

func (d *day16) Part2(input string) (string, error) {
	// Reset all paths in case part1 was run before.
	allPathsMap = make(map[string][]valve)

	// Generate again but with 26 minutes left
	startingPath := make([]valve, 0)
	startingPath = append(startingPath, d.valves[d.start])
	maxMinutes := 26
	d.generateAllPaths(startingPath, d.getValvesWithFlowAndNotOpened(), 0, maxMinutes)

	// Me:			[JJ, BB, CC]
	// Elephant:	[DD, HH, EE]

	pathsPressures := make([]pathPressure, 0)
	for _, path := range allPathsMap {
		p := pathPressure{
			path:     path[1:], // Skip start ('AA')
			pressure: d.getPathPressure(path, maxMinutes),
		}
		pathsPressures = append(pathsPressures, p)
	}

	// The idea is to find the two paths with highest pressure that don't overlap.
	out := 0
	// paths := [2][]valve{}
	//! This will take a while... ~40m :<
	for i, pp := range pathsPressures {
		for j := i + 1; j < len(pathsPressures); j++ {
			overlaps := false
			for _, v := range pathsPressures[j].path {
				overlaps = strings.Contains(fmt.Sprint(pp.path), v.id)
				if overlaps {
					break
				}
			}
			if !overlaps && out < pp.pressure+pathsPressures[j].pressure {
				out = pp.pressure + pathsPressures[j].pressure
				// paths[0], paths[1] = pp.path, pathsPressures[j].path
			}
		}
	}

	// fmt.Println(paths)

	return fmt.Sprint(out), nil
}

func (d *day16) Exec(input string) (*DayResult, error) {
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
