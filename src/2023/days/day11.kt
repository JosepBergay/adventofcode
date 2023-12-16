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

    private fun sumGalaxiesDistances(expandFactor: Int): Long {
        val emptyRows =
                (0..height).filter { y -> (0..width).all { Point(it, y) !in galaxies } }.toHashSet()
        val emptyCols =
                (0..width).filter { x -> (0..height).all { Point(x, it) !in galaxies } }.toHashSet()

        var out = 0L
        val galaxies = galaxies.toHashSet() // Copy

        while (galaxies.isNotEmpty()) {
            val gal = galaxies.first()
            galaxies.remove(gal)

            val dist =
                    galaxies.sumOf { p ->
                        gal.manhattan(p).toLong() + // <- THIS!
                        (min(gal.x, p.x)..max(gal.x, p.x)).count { it in emptyCols } *
                                        expandFactor +
                                (min(gal.y, p.y)..max(gal.y, p.y)).count { it in emptyRows } *
                                        expandFactor
                    }
            out += dist
        }

        return out
    }

    override fun part1(): Long {
        return sumGalaxiesDistances(1)
    }

    override fun part2(): Long {
        return sumGalaxiesDistances(1_000_000 - 1)
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
