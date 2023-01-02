package aoc2022

import (
	"aoc2022/utils"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type day20 struct{}

func init() {
	Days[20] = &day20{}
}

func (d *day20) Parse(input string) ([]int, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)

	out := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()

		i, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}

		out = append(out, i)
	}

	return out, nil
}

type node struct {
	value     int
	processed bool
}

func (d *day20) Part1(input []int) (string, error) {
	// Using a doubly linked list to avoid shifting values each time.
	list := utils.DoublyLinkedList[node]{}

	// Init list
	for _, v := range input {
		list.Append(node{v, false})
	}

	for _, v := range input {
		// 1.- Remove from current position and find index
		i, err := list.Remove(node{v, false})
		if err != nil {
			return "", err
		}

		// 2.- Determine new index
		newIdx := (i + v) % (len(input) - 1)
		if newIdx < 0 {
			newIdx += len(input) - 1
		}

		// 3.- Insert at new index
		err = list.InsertAt(node{v, true}, newIdx)
		if err != nil {
			fmt.Println(v, newIdx, i)
			return "", err
		}
	}

	// Get index of value 0
	i, err := list.GetIndex(node{0, true})
	if err != nil {
		return "", err
	}

	out := 0

	for _, c := range [3]int{i + 1000, i + 2000, i + 3000} {
		n, err := list.Get(c % len(input))
		if err != nil {
			return "", err
		}
		out += n.value
	}

	return fmt.Sprint(out), nil
}

func (d *day20) Part2(input []int) (string, error) {
	return "TODO", nil
}

func (d *day20) Exec(input string) (*DayResult, error) {
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
