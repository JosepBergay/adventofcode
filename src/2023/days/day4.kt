package aoc2023.days

import kotlin.io.path.readLines
import kotlin.math.*

val testInputD4 =
        """
Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
"""

class Day4 : BaseDay(4) {
    var input = listOf<String>()

    override fun parse() {
        input = inputPath.readLines()
        // input = testInputD4.reader().readLines()
    }

    override fun part1(): Int {
        val regex = """\d+""".toRegex()

        return input.filter { it.isNotEmpty() }.sumOf {
            it.split("|", ":")
                    .drop(1)
                    .let {
                        regex.findAll(it.first()).map { it.value } to
                                regex.findAll(it.last()).map { it.value }
                    }
                    .let { (winning, nums) -> winning.filter { nums.contains(it) }.toList() }
                    .let { if (it.size > 0) 2.0.pow(it.size - 1) else 0.0 }
                    .toInt()
        }
    }

    override fun part2(): String {

        return "TODO"
    }
}
