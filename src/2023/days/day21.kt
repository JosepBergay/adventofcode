package aoc2023.days

import kotlin.io.path.readLines

class Day21 : BaseDay(21) {
    val input = hashMapOf<Point, Char>()
    var start = Point(0, 0)
    var maxX = 0
    var maxY = 0

    override fun parse() {
        val lines = inputPath.readLines()
        // val lines = testInputD21.reader().readLines()

        maxY = lines.size - 1
        maxX = lines[0].length - 1

        for ((y, line) in lines.withIndex()) {
            for ((x, c) in line.withIndex()) {
                input[Point(x, y)] = c
                if (c == 'S') {
                    start = Point(x, y)
                }
            }
        }
    }

    override fun part1(): Any {
        val maxSteps = 64

        var curr = hashSetOf(start)

        repeat(maxSteps) {
            curr =
                    curr
                            .flatMap { it.getAdjacents(maxX, maxY).filter { input[it] != '#' } }
                            .toHashSet()
        }

        return curr.size
    }

    override fun part2(): Any {

        return "TODO"
    }
}

val testInputD21 =
        """...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........
"""
