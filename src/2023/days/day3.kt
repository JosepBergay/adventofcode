package aoc2023.days

import kotlin.collections.mutableListOf
import kotlin.io.path.readLines

val testInputD3 =
        """467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..

"""

data class Point(val x: Int, val y: Int)

fun Point.getAdjacents(
        width: Int,
        height: Int = width,
        withDiagonals: Boolean = false
): List<Point> {
    val directions = mutableListOf(Point(1, 0), Point(-1, 0), Point(0, 1), Point(0, -1))

    if (withDiagonals) {
        val diagonals = listOf(Point(1, 1), Point(-1, -1), Point(-1, 1), Point(1, -1))
        directions.addAll(diagonals)
    }

    return directions.map { Point(x + it.x, y + it.y) }.filterOutOfBounds(width, height)
}

fun List<Point>.filterOutOfBounds(width: Int, height: Int = width): List<Point> {
    return this.filter { it.x in 0..width && it.y in 0..height }
}

class Day3 : BaseDay(3) {
    var lines = listOf<String>()
    var width = 0
    var height = 0

    override fun parse() {
        lines = inputPath.readLines()
        // lines = testInputD3.reader().readLines()
        height = lines.filter { it.isNotEmpty() }.size - 1
        width = lines[0].length - 1
    }

    override fun part1(): Int {
        var sum = 0

        for ((y, line) in lines.withIndex()) {
            var num = ""
            var hasSymbol = false

            for ((x, c) in line.withIndex()) {
                val point = Point(x, y)

                if (c.isDigit()) {
                    num += c

                    hasSymbol =
                            hasSymbol ||
                                    point.getAdjacents(width, height, true).any {
                                        lines[it.y][it.x] != '.' && !lines[it.y][it.x].isDigit()
                                    }
                } else if (num != "") {
                    if (hasSymbol) {
                        sum += num.toInt()
                        hasSymbol = false
                    }
                    num = ""
                }
            }
        }

        return sum
    }

    private fun getNumFromPoint(p: Point): String {
        var num = "${lines[p.y][p.x]}"

        for (x in p.x - 1 downTo 0) {
            if (!lines[p.y][x].isDigit()) break

            num = lines[p.y][x] + num
        }

        for (x in p.x + 1..width) {
            if (!lines[p.y][x].isDigit()) break

            num += lines[p.y][x]
        }

        return num
    }

    override fun part2(): Int {
        val topSide = listOf(Point(-1, -1), Point(0, -1), Point(1, -1))
        val leftSide = listOf(Point(-1, 0))
        val rightSide = listOf(Point(1, 0))
        val botSide = listOf(Point(-1, 1), Point(0, 1), Point(1, 1))

        val sides = listOf(topSide, leftSide, rightSide, botSide)

        return lines.withIndex().sumOf { (y, line) ->
            line.withIndex()
                    .filter { (x, c) ->
                        c == '*' &&
                                sides.sumOf {
                                    it
                                            .map { Point(x + it.x, y + it.y) }
                                            .filterOutOfBounds(width, height)
                                            .map { lines[it.y][it.x] }
                                            .joinToString("")
                                            // Assume gears are not near any other Symbol
                                            .split('.')
                                            .filter { it.isNotEmpty() && it.all { it.isDigit() } }
                                            .size
                                } == 2
                    }
                    .sumOf { (x) ->
                        var out = 1

                        for (side in sides) {
                            var parsing = false
                            val neighbours =
                                    side
                                            .map { Point(x + it.x, y + it.y) }
                                            .filterOutOfBounds(width, height)

                            for (p in neighbours) {
                                if (!parsing && lines[p.y][p.x].isDigit()) {
                                    out *= getNumFromPoint(p).toInt()
                                    parsing = true
                                } else if (!lines[p.y][p.x].isDigit()) {
                                    parsing = false
                                }
                            }
                        }

                        out
                    }
        }
    }
}
