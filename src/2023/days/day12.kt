package aoc2023.days

import kotlin.io.path.readLines
import kotlin.math.*

class Day12 : BaseDay(12) {
    var input = listOf<Pair<String, List<Int>>>()
    val cache = hashMapOf<Pair<String, List<Int>>, Long>()

    override fun parse() {
        // val lines = testInputD12.reader().readLines()
        val lines = inputPath.readLines()
        input =
                lines.map {
                    val (springs, groups) = it.split(" ")
                    springs to groups.split(",").map { it.toInt() }
                }
    }

    private fun validateRecord(springs: String, groups: List<Int>): Boolean {
        val newGroups = springs.split(".").filter { it.isNotEmpty() }

        return newGroups.size == groups.size &&
                newGroups.withIndex().all { (i, s) -> groups.size > i && s.length == groups[i] }
    }

    private fun countArrangementsBruteForce(springs: String, groups: List<Int>): Int {
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

    private fun countArrangements(springs: String, groups: List<Int>): Long {
        if (springs.length == 0 && groups.size == 0) return 1 // Nothing left so it fits.
        else if (springs.length == 0) return 0 // No springs but still got groups left.
        else if (groups.size == 0) {
            return if ('#' in springs) 0 else 1 // If there're no damaged springs it fits.
        }

        if (springs.length < groups.sum() + groups.size - 1) return 0 // Line is not long enough.

        return cache.getOrPut(springs to groups) {
            when (springs.first()) {
                '.' -> countArrangements(springs.drop(1), groups)
                '?' ->
                        countArrangements("." + springs.drop(1), groups) +
                                countArrangements("#" + springs.drop(1), groups)
                '#' -> {
                    val group = groups.first()

                    // Check if group length contains undamaged spring.
                    if (springs.slice(0 ..< group).any { it == '.' }) return 0

                    // Check if at the end of the group there's another group.
                    if (group < springs.length && springs[group] == '#') return 0

                    return countArrangements(springs.drop(group + 1), groups.drop(1))
                }
                else -> error("woot")
            }
        }
    }

    override fun part1(): Long {
        return input.sumOf { countArrangements(it.first, it.second) }
    }

    override fun part2(): Long {
        val unfolded =
                input.map { (s, g) ->
                    val springs = (1..5).map { s }.joinToString("?")
                    val groups = (1..5).flatMap { g }
                    springs to groups
                }

        return unfolded.sumOf { countArrangements(it.first, it.second) }
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
