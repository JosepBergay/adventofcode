package aoc2022

type day2 struct{}

func init() {
	Days[2] = &day2{}
}

func (d *day2) Parse(input string) (string, error) {
	return "TODO", nil
}

func (d *day2) Part1(input string) (string, error) {
	return "TODO", nil
}

func (d *day2) Part2(input string) (string, error) {
	return "TODO", nil
}

func (d *day2) Exec(input string) (*DayResult, error) {
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
