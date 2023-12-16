package aoc2023.days

import kotlin.io.path.readLines
import kotlin.math.*

class Day12 : BaseDay(12) {
    var input = listOf<String>()

    override fun parse() {
        // input = testInputD12.reader().readLines()
        input = inputPath.readLines()
    }

    private fun validateRecord(springs: String, groups: List<Int>): Boolean {
        val newGroups = springs.split(".").filter { it.isNotEmpty() }

        return newGroups.size == groups.size &&
                newGroups.withIndex().all { (i, s) -> groups.size > i && s.length == groups[i] }
    }

    private fun countArrangements(springs: String, groups: List<Int>): Int {
        val unknownIdxs = springs.mapIndexedNotNull { i, c -> if (c == '?') i else null }

        val sb = StringBuilder(springs)

        val maxArrangements = 2.toDouble().pow(unknownIdxs.size).toInt()

        // Brute forcing :<
        return (0 ..< maxArrangements)
                .map {
                    it.toString(2).padStart(unknownIdxs.size, '0').map {
                        if (it == '1') '#' else '.'
                    }
                }
                .map {
                    it.forEachIndexed { i, c -> sb.setCharAt(unknownIdxs[i], c) }
                    sb.toString()
                }
                .count { validateRecord(it, groups) }
    }

    override fun part1(): Int {
        return input.sumOf {
            val (springs, groups) = it.split(" ")
            countArrangements(springs, groups.split(",").map { it.toInt() })
        }
    }

    override fun part2(): Any {

        return "TODO"
    }
}

val testInputD12 =
        """???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1
"""
