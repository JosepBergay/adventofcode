package aoc2022

import (
	"aoc2022/utils"
	"bufio"
	"fmt"
	"math"
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
	order     int
}

func (n node) String() string {
	return fmt.Sprint(n.value)
}

func findNodeIdxInOriginalOrder(list utils.DoublyLinkedList[node], value int, processed bool) int {
	minOrder := math.MaxInt
	idx := -1
	list.Traverse(func(item node, i int) bool {
		if item.value == value && item.processed == processed && item.order < minOrder {
			minOrder = item.order
			idx = i
		}
		return false
	})
	return idx
}

func getGrooveCoordinates(input []int, rounds int) (int, error) {
	// Using a doubly linked list to avoid shifting values each time.
	list := utils.DoublyLinkedList[node]{}

	// Init list
	for i, v := range input {
		list.Append(node{v, false, i})
	}

	processed := false
	maxIdx := len(input) - 1
	for round := 0; round < rounds; round++ {
		for _, v := range input {
			// 1.- Find next node in order that is not yet processed
			i := findNodeIdxInOriginalOrder(list, v, processed)

			// 2.- Remove from current position
			n, err := list.RemoveAt(i)
			if err != nil {
				return -1, err
			}

			// 3.- Mark it as processed
			n.processed = !n.processed

			if v == 0 {
				list.InsertAt(n, i)
				continue
			}

			// 4.- Determine new index
			newIdx := (i + v) % maxIdx
			if newIdx < 0 {
				newIdx += maxIdx
			}

			// 5.- Insert at new index
			err = list.InsertAt(n, newIdx)
			if err != nil {
				return -1, err
			}
		}
		processed = !processed
	}

	// Get index of value 0
	_, i, err := list.Find(func(item node, idx int) bool {
		return item.value == 0
	})
	if err != nil {
		return -1, err
	}

	out := 0

	for _, c := range [3]int{i + 1000, i + 2000, i + 3000} {
		n, err := list.Get(c % len(input))
		if err != nil {
			return -1, err
		}
		out += n.value
	}

	return out, nil
}

func (d *day20) Part1(input []int) (string, error) {
	coords, err := getGrooveCoordinates(input, 1)
	if err != nil {
		return "", err
	}

	return fmt.Sprint(coords), nil
}

func (d *day20) Part2(input []int) (string, error) {
	decryptionKey := 811589153

	for i := 0; i < len(input); i++ {
		v := input[i] * decryptionKey
		input[i] = v
	}

	coords, err := getGrooveCoordinates(input, 10)
	if err != nil {
		return "", err
	}

	return fmt.Sprint(coords), nil
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
