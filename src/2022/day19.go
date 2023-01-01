package aoc2022

import (
	"bufio"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type day19 struct{}

type materials struct {
	ore      int
	clay     int
	obsidian int
	geode    int
}

type blueprint struct {
	id            int
	oreRobot      materials
	clayRobot     materials
	obsidianRobot materials
	geodeRobot    materials
}

func init() {
	Days[19] = &day19{}
}

var re = regexp.MustCompile(`\d+`)

func (d *day19) Parse(input string) ([]blueprint, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)

	blueprints := make([]blueprint, 0)

	for scanner.Scan() {
		line := scanner.Text()

		strs := re.FindAllString(line, -1)

		nums := [7]int{}
		for i, s := range strs {
			n, err := strconv.Atoi(s)
			if err != nil {
				return nil, err
			}
			nums[i] = n
		}

		bp := blueprint{
			id:            nums[0],
			oreRobot:      materials{ore: nums[1]},
			clayRobot:     materials{ore: nums[2]},
			obsidianRobot: materials{ore: nums[3], clay: nums[4]},
			geodeRobot:    materials{ore: nums[5], obsidian: nums[6]},
		}

		blueprints = append(blueprints, bp)
	}

	return blueprints, nil
}

func canBuildRobot(robotBlueprint materials, resources materials) bool {
	return robotBlueprint.ore <= resources.ore &&
		robotBlueprint.clay <= resources.clay &&
		robotBlueprint.obsidian <= resources.obsidian
}

func buildRobot(robotBlueprint materials, resources *materials) {
	resources.clay -= robotBlueprint.clay
	resources.obsidian -= robotBlueprint.obsidian
	resources.ore -= robotBlueprint.ore
}

type robotFactoryState struct {
	robotCount [4]int
	resources  materials
}

func (s *robotFactoryState) gatherResources() {
	s.resources.ore += s.robotCount[0]
	s.resources.clay += s.robotCount[1]
	s.resources.obsidian += s.robotCount[2]
	s.resources.geode += s.robotCount[3]
}

func (s *robotFactoryState) canBuildRobot(robotBlueprint materials) bool {
	return canBuildRobot(robotBlueprint, s.resources)
}

func (s *robotFactoryState) buildRobot(robotBlueprint materials) {
	buildRobot(robotBlueprint, &s.resources)
}

func getMaxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func computeBluePrintMaxGeodes(bp blueprint, minutes int) int {
	initialState := robotFactoryState{}
	initialState.robotCount[0] = 1

	var runFactory func(s robotFactoryState, timeLeft, buildNext int) int
	runFactory = func(s robotFactoryState, timeLeft, buildNext int) int {
		if timeLeft == 0 {
			return s.resources.geode
		}

		// Stop building this robot if we have more resources than what's needed.
		// This keeps our resources low and makes sure we follow a path that alternates between the
		// robots we build.
		if buildNext == 0 && s.resources.ore >= getMaxInt(bp.oreRobot.ore,
			getMaxInt(bp.clayRobot.ore,
				getMaxInt(bp.obsidianRobot.ore, bp.geodeRobot.ore))) {
			return 0
		}
		if buildNext == 1 && s.resources.clay >= bp.obsidianRobot.clay {
			return 0
		}
		// Although our input does not fail we got to add `s.resources.clay == 0` for tests to pass.
		// Case is when we don't have enough obsidian but what we really need is clay
		if buildNext == 2 &&
			(s.resources.obsidian >= bp.geodeRobot.obsidian || s.resources.clay == 0) {
			return 0
		}
		// Always build geodeRobots
		// if buildNext == 3 && s.resources.geode >= bp.g...

		maxGeodes := 0

		for timeLeft > 0 {
			if buildNext == 0 && s.canBuildRobot(bp.oreRobot) {
				s.buildRobot(bp.oreRobot)
				s.gatherResources()
				s.robotCount[0]++
				for i := 0; i < 4; i++ {
					maxGeodes = getMaxInt(maxGeodes, runFactory(s, timeLeft-1, i))
				}
				return maxGeodes
			} else if buildNext == 1 && s.canBuildRobot(bp.clayRobot) {
				s.buildRobot(bp.clayRobot)
				s.gatherResources()
				s.robotCount[1]++
				for i := 0; i < 4; i++ {
					maxGeodes = getMaxInt(maxGeodes, runFactory(s, timeLeft-1, i))
				}
				return maxGeodes
			} else if buildNext == 2 && s.canBuildRobot(bp.obsidianRobot) {
				s.buildRobot(bp.obsidianRobot)
				s.gatherResources()
				s.robotCount[2]++
				for i := 0; i < 4; i++ {
					maxGeodes = getMaxInt(maxGeodes, runFactory(s, timeLeft-1, i))
				}
				return maxGeodes
			} else if buildNext == 3 && s.canBuildRobot(bp.geodeRobot) {
				s.buildRobot(bp.geodeRobot)
				s.gatherResources()
				s.robotCount[3]++

				for i := 0; i < 4; i++ {
					maxGeodes = getMaxInt(maxGeodes, runFactory(s, timeLeft-1, i))
				}
				return maxGeodes
			}

			// Can't build next robot so just gather and carry on
			s.gatherResources()
			maxGeodes = getMaxInt(maxGeodes, s.resources.geode)
			timeLeft--
		}

		return maxGeodes
	}

	maxGeodes := 0
	for i := 0; i < 4; i++ {
		maxGeodes = getMaxInt(maxGeodes, runFactory(initialState, minutes, i))
	}

	return maxGeodes
}

func (d *day19) Part1(input []blueprint) (string, error) {
	minutes := 24
	out := 0

	for _, bp := range input {
		out += bp.id * computeBluePrintMaxGeodes(bp, minutes)
	}

	return fmt.Sprint(out), nil
}

func (d *day19) Part2(input []blueprint) (string, error) {
	minutes := 32
	out := 1

	for i, bp := range input {
		if i == 3 {
			break
		}
		out *= computeBluePrintMaxGeodes(bp, minutes)
	}

	return fmt.Sprint(out), nil
}

func (d *day19) Exec(input string) (*DayResult, error) {
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
