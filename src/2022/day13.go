package aoc2022

import (
	"bufio"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type day13 struct{}

func init() {
	Days[13] = &day13{}
}

type packetData struct {
	value   *int
	packets []*packetData
	parent  *packetData
}

func (p *packetData) String() string {
	if p.value != nil {
		return fmt.Sprint(*p.value)
	}

	return fmt.Sprint(p.packets)
}

type day13Input struct {
	pairs [][2]*packetData
}

func createPacketData(value *int, parent *packetData) *packetData {
	return &packetData{parent: parent, value: value}
}

func parsePackedData(line string) (*packetData, error) {
	curr := &packetData{}

	for i := 0; i < len(line); i++ {
		r := line[i]
		switch r {
		case '[':
			if curr.packets == nil {
				curr.packets = make([]*packetData, 0)
			}
			newPacketData := createPacketData(nil, curr)
			curr.packets = append(curr.packets, newPacketData)
			curr = newPacketData
		case ']':
			curr = curr.parent
		case ',':
			// Do nothing
		default:
			buf := make([]byte, 0)

			for {
				// Looping in case we have multiple digit integers
				_, err := strconv.Atoi(string(line[i]))
				if err != nil {
					i--
					break
				}
				buf = append(buf, line[i])
				i++
			}

			n, err := strconv.Atoi(string(buf))
			if err != nil {
				return nil, err
			}

			curr.packets = append(curr.packets, createPacketData(&n, curr))
		}
	}

	return curr, nil
}

func (d *day13) Parse(input string) (*day13Input, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)

	pairs := make([][2]*packetData, 0)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		left, err := parsePackedData(line)
		if err != nil {
			return nil, err
		}

		scanner.Scan()
		line = scanner.Text()

		right, err := parsePackedData(line)
		if err != nil {
			return nil, err
		}

		pairs = append(pairs, [2]*packetData{left, right})
	}

	return &day13Input{pairs}, nil
}

func convertToList(p *packetData) {
	p.packets = make([]*packetData, 0)
	p.packets = append(p.packets, createPacketData(p.value, p.parent))
	p.value = nil
}

func areInOrder(left, right packetData) int {
	if left.value == nil && right.value != nil {
		convertToList(&right)
	}

	if right.value == nil && left.value != nil {
		convertToList(&left)
	}

	if left.value != nil && right.value != nil {
		return *left.value - *right.value
	}

	minLen := math.Min(float64(len(left.packets)), float64(len(right.packets)))

	for i := 0; i < int(minLen); i++ {
		diff := areInOrder(*left.packets[i], *right.packets[i])
		if diff != 0 {
			return diff
		}
	}

	return len(left.packets) - len(right.packets)
}

func (d *day13) Part1(input *day13Input) (string, error) {
	sum := 0

	for i, p := range input.pairs {
		diff := areInOrder(*p[0].packets[0], *p[1].packets[0])
		if diff < 0 {
			sum += i + 1
		}
	}

	return fmt.Sprint(sum), nil
}

type packets []*packetData

func (a packets) Len() int           { return len(a) }
func (a packets) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a packets) Less(i, j int) bool { return areInOrder(*a[i], *a[j]) < 0 }

func (d *day13) Part2(input *day13Input) (string, error) {
	dividersStr := "[[2]]\n[[6]]"

	dividers, err := d.Parse(dividersStr)
	if err != nil {
		return "", err
	}

	flat := make(packets, 0)
	for _, p := range input.pairs {
		flat = append(flat, p[0], p[1])
	}
	flat = append(flat, dividers.pairs[0][0], dividers.pairs[0][1])

	// TODO: implement our own sorting
	sort.Sort(flat) // Sorts in place using 'https://github.com/orlp/pdqsort' from std lib

	decoderKey := 1
	for i, v := range flat {
		s := fmt.Sprint(v)
		if s == "[[[2]]]" || s == "[[[6]]]" {
			decoderKey *= i + 1
		}
	}

	return fmt.Sprint(decoderKey), nil
}

func (d *day13) Exec(input string) (*DayResult, error) {
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
