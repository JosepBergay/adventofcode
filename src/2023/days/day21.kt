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

    override fun part1(): Int {
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

    override fun part2(): Long {
        // Inspired by brilliant posts:
        // https://github.com/villuna/aoc23/wiki/A-Geometric-solution-to-advent-of-code-2023,-day-21
        // https://aoc.csokavar.hu/?day=21

        val maxSize = maxX + 1 // 131x131

        // val maxSteps = 26501365
        val maxSteps = 65 + 2 * maxSize // Enough with 2.5x (Half tile width + 2*Tile width)

        val edges = mutableListOf<Long>()

        var curr = hashSetOf(start)

        repeat(maxSteps + 1) {
            if (it % maxSize == 65) {
                // Entered new big tile
                edges += curr.size.toLong()
            }

            curr =
                    curr
                            .flatMap { p ->
                                Direction.entries.map { it.p + p }.filter {
                                    input[it.mod(maxSize)] != '#'
                                }
                            }
                            .toHashSet()
        }

        // Quadratic interpolation
        // a * n^2 + b * n + c      if n = k * 131 + 65

        val n = (26501365 / maxSize).toLong()
        val (a, b, c) = edges

        val out = a + n * (b - a + (n - 1) * (c - b - b + a) / 2)

        return out
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
