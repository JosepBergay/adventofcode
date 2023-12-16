package aoc2023.days

import kotlin.io.path.readLines
import kotlin.math.*

class Day11 : BaseDay(11) {
    val galaxies = hashSetOf<Point>()
    var height = 0
    var width = 0

    override fun parse() {
        // val lines = testInputD11.reader().readLines()
        val lines = inputPath.readLines()
        height = lines.size - 1
        width = lines[0].length - 1
        galaxies +=
                lines.flatMapIndexed { y, line ->
                    line.flatMapIndexed { x, c ->
                        if (c == '#') listOf(Point(x, y)) else emptyList()
                    }
                }
    }

    override fun part1(): Any {
        val emptyRows = hashSetOf<Int>()
        emptyRows += (0..height).filter { y -> (0..width).all { Point(it, y) !in galaxies } }
        val emptyCols = hashSetOf<Int>()
        emptyCols += (0..width).filter { x -> (0..height).all { Point(x, it) !in galaxies } }

        var out = 0

        while (galaxies.isNotEmpty()) {
            val gal = galaxies.first()
            galaxies.remove(gal)

            val dist =
                    galaxies.sumOf { p ->
                        gal.manhattan(p) +
                                (0..abs(gal.x - p.x)).count {
                                    emptyCols.contains(it + min(gal.x, p.x))
                                } +
                                (0..abs(gal.y - p.y)).count {
                                    emptyRows.contains(it + min(gal.y, p.y))
                                }
                    }
            out += dist
        }

        return out
    }

    override fun part2(): Any {

        return "TODO"
    }
}

val testInputD11 =
        """...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....
"""
