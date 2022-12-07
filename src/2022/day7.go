package aoc2022

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type day7 struct{}

func init() {
	Days[7] = &day7{}
}

type fileTree struct {
	children []*fileTree
	parent   *fileTree
	isDir    bool
	name     string
	size     int
}

func createNode(line string) (*fileTree, error) {
	cmd := strings.Split(line, " ")

	isDir := cmd[0] == "dir"

	node := &fileTree{isDir: isDir, children: make([]*fileTree, 0), name: cmd[1]}

	if !isDir {
		size, err := strconv.Atoi(cmd[0])
		if err != nil {
			return nil, err
		}

		node.size = size
	}

	return node, nil
}

func (d *day7) Parse(input string) (*fileTree, error) {
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	rootTree := &fileTree{
		isDir:    true,
		name:     "/",
		children: make([]*fileTree, 0),
	}

	curr := rootTree

	for scanner.Scan() {
		line := scanner.Text()

		if line == "$ cd /" {
			// Set current to root
			curr = rootTree
		} else if line == "$ cd .." {
			if curr.parent == nil {
				return nil, fmt.Errorf("can't go to parent from %v", curr.name)
			}
			curr = curr.parent
		} else if strings.Index(line, "$ cd ") == 0 {
			// Step into to child
			dirName := line[len("$ cd "):]
			for _, child := range curr.children {
				if child.name == dirName {
					curr = child
				}
			}
		} else if strings.Index(line, "$ ls") == 0 {
			// List, just hold it
		} else {
			// Add children to current
			newNode, err := createNode(line)
			if err != nil {
				return nil, err
			}

			newNode.parent = curr
			curr.children = append(curr.children, newNode)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}

	return rootTree, nil
}

func populateDirSize(node *fileTree) int {
	size := 0

	for _, child := range node.children {
		if child.isDir {
			size += populateDirSize(child)
		} else {
			size += child.size
		}
	}

	node.size = size

	return size
}

func addDirSize(node *fileTree, maxSize, acc int) int {
	for _, child := range node.children {
		if child.isDir {
			acc = addDirSize(child, maxSize, acc)
		}
	}

	if node.isDir && node.size <= maxSize {
		acc += node.size
	}

	return acc
}

func (d *day7) Part1(root *fileTree) (string, error) {
	populateDirSize(root)

	out := addDirSize(root, 100_000, 0)

	return fmt.Sprint(out), nil
}

func findSmallestDirSpace(node *fileTree, minSpace, currentSmallest int) int {
	if !node.isDir {
		return currentSmallest
	}

	for _, child := range node.children {
		currentSmallest = findSmallestDirSpace(child, minSpace, currentSmallest)
	}

	if node.size > minSpace && node.size < currentSmallest {
		return node.size
	}

	return currentSmallest
}

func (d *day7) Part2(root *fileTree) (string, error) {
	missingSpace := root.size + 30_000_000 - 70_000_000

	out := findSmallestDirSpace(root, missingSpace, root.size)

	return fmt.Sprint(out), nil
}

func (d *day7) Exec(input string) (*DayResult, error) {
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
