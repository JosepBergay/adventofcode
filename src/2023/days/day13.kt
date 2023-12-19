package aoc2023.days

import kotlin.io.path.readText
import kotlin.math.*

class Day13 : BaseDay(13) {
    var input = listOf<String>()

    override fun parse() {
        input = inputPath.readText().split("\n\n")
        // input = testInputD13.reader().readText().split("\n\n")
    }

    private fun getColumns(rows: List<String>): List<String> {
        if (rows.isEmpty()) return emptyList()
        return (0..rows[0].length - 1).map { i -> rows.map { it[i] }.joinToString("") }
    }

    private fun validateReflection(lines: List<String>): Boolean {
        if (lines.isEmpty()) return true

        val first = lines.first()
        val last = lines.last()

        return if (first == last) validateReflection(lines.drop(1).dropLast(1)) else false
    }

    private fun getMirrorIndexes(idx: Int, lines: List<String>): List<Int> {
        return lines.withIndex()
                .filter { (i, s) -> lines[idx] == s && i != idx && i % 2 == 1 }
                .map { it.index }
    }

    private fun findLineOfReflection(idx: Int, lines: List<String>): List<Int> {
        return getMirrorIndexes(idx, lines).filter {
            validateReflection(lines.slice(if (idx == 0) 0..it else it..lines.size - 1))
        }
    }

    private fun findLineOfReflection(pattern: String): Pair<Boolean, Int> {
        val rows = pattern.reader().readLines()

        var idx = findLineOfReflection(0, rows).firstOrNull()

        if (idx != null) return (false to 1 + idx / 2)

        idx = findLineOfReflection(rows.size - 1, rows).firstOrNull()

        if (idx != null) return (false to idx + (rows.size - idx) / 2)

        val columns = getColumns(rows)

        idx = findLineOfReflection(0, columns).firstOrNull()

        if (idx != null) return (true to 1 + idx / 2)

        idx = findLineOfReflection(columns.size - 1, columns).firstOrNull()

        if (idx != null) return (true to idx + (columns.size - idx) / 2)

        error("woot")
    }

    override fun part1(): Int {
        return input.sumOf {
            findLineOfReflection(it).let { (vertical, num) -> if (vertical) num else num * 100 }
        }
    }

    override fun part2(): Any {

        return "TODO"
    }
}

val testInputD13 =
        """###....##
##......#
####..###
..#....#.
###....##
..######.
##......#
##......#
###.##.##
##.#.##.#
###.##.##
..######.
..#....#.
"""
// val testInputD13 =
//         """#.##..##.
// ..#.##.#.
// ##......#
// ##......#
// ..#.##.#.
// ..##..##.
// #.#.##.#.

// #...##..#
// #....#..#
// ..##..###
// #####.##.
// #####.##.
// ..##..###
// #....#..#
// """
