package aoc2022

import "testing"

const inputD7 = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k
`

const expectedD7P1 = "95437"

func TestDay7Part1(t *testing.T) {
	day := &day7{}

	parsed, err := day.Parse(inputD7)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part1(parsed)

	if err != nil {
		t.Errorf("[Part1]: %v", err.Error())
		return
	}

	if res != expectedD7P1 {
		t.Errorf("Expected: %v \nBut got: %v", expectedD7P1, res)
	}
}

const expectedD7P2 = ""

func TestDay7Part2(t *testing.T) {
	day := &day7{}

	parsed, err := day.Parse(inputD7)

	if err != nil {
		t.Errorf("[Parse]: %v", err.Error())
		return
	}

	res, err := day.Part2(parsed)

	if err != nil {
		t.Errorf("[Part2]: %v", err.Error())
		return
	}

	if res != expectedD7P2 {
		t.Errorf("Expected: %v \nBut got: %v", expectedD7P2, res)
	}
}
