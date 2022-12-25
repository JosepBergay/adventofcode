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
		// if curr.connected == nil {
		// 	fmt.Println(from, to, curr)
		// }
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

func (d *day16) getPathPressure(path []valve) int {
	start := d.valves[d.start]
	minutes := 30
	total := 0

	curr := start
	pos := 0
	for i := minutes; i > 0; i-- {
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

func (d *day16) generateAllPaths(path []valve, unseen []valve, distance int) {
	if len(unseen) == 0 {
		allPathsMap[fmt.Sprint(path)] = path
		return
	}

	last := path[len(path)-1]

	for _, v := range unseen {
		dist := d.findShortestDistanceBetweenValves(last, v)

		if dist+1+distance >= 30 {
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
		d.generateAllPaths(copiedPath, newUnseen, dist+1+distance)
	}
}

func (d *day16) Part1(input string) (string, error) {
	startingPath := make([]valve, 0)
	startingPath = append(startingPath, d.valves[d.start])
	d.generateAllPaths(startingPath, d.getValvesWithFlowAndNotOpened(), 0)

	fmt.Println(len(allPathsMap))

	out := 0
	for _, path := range allPathsMap {
		if p := d.getPathPressure(path); p > out {
			out = p
		}
	}

	return fmt.Sprint(out), nil
}

func (d *day16) Part2(input string) (string, error) {
	return "TODO", nil
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
