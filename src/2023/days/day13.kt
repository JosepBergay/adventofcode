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

    private fun findLineOfReflection(pattern: String): Pair<Boolean, Int> {
        val rows = pattern.reader().readLines()

        var idx =
                getMirrorIndexes(0, rows).asReversed().find {
                    validateReflection(rows.slice(0..it))
                }

        if (idx != null) return (false to 1 + idx / 2)

        idx =
                getMirrorIndexes(rows.size - 1, rows).find {
                    validateReflection(rows.slice(it..rows.size - 1))
                }

        if (idx != null) return (false to idx + (rows.size - idx) / 2)

        val columns = getColumns(rows)

        idx =
                getMirrorIndexes(0, columns).asReversed().find {
                    validateReflection(columns.slice(0..it))
                }

        if (idx != null) return (true to 1 + idx / 2)

        idx =
                getMirrorIndexes(columns.size - 1, columns).find {
                    validateReflection(columns.slice(it..columns.size - 1))
                }

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
