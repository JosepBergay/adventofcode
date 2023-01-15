package aoc2022

type day25 struct{}

func init() {
	Days[25] = &day25{}
}

func (d *day25) Parse(input string) (string, error) {
	return "TODO", nil
}

func (d *day25) Part1(input string) (string, error) {
	return "TODO", nil
}

func (d *day25) Part2(input string) (string, error) {
	return "TODO", nil
}

func (d *day25) Exec(input string) (*DayResult, error) {
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
