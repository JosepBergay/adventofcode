package aoc2023.days

import kotlin.io.path.readLines

class Day14 : BaseDay(14) {
    val input = mutableListOf<Int>()

    override fun parse() {}

    override fun part1(): Int {
        val lines = inputPath.readLines()

        return (0..lines[0].length - 1).sumOf {
            var sum = 0
            var emptyIdx = -1
            var emptyCapacity = 0

            for (i in 0..lines.size - 1) {
                when (lines[i][it]) {
                    '#' -> {
                        emptyIdx = -1
                        emptyCapacity = 0
                    }
                    '.' -> {
                        if (emptyIdx == -1) {
                            emptyIdx = i
                        }
                        emptyCapacity++
                    }
                    'O' -> {
                        if (emptyIdx == -1) {
                            sum += lines.size - i
                        } else {
                            sum += lines.size - emptyIdx
                            emptyIdx++
                        }
                    }
                }
            }

            sum
        }
    }

    override fun part2(): Any {

        return "TODO"
    }
}

val testInputD14 =
        """O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....
"""
