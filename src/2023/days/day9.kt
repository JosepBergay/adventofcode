package aoc2023.days

import kotlin.io.path.readLines

class Day9 : BaseDay(9) {

    override fun parse() {}

    override fun part1(): Int {
        return inputPath.readLines().sumOf {
            val nums = it.split(" ").map { it.toInt() }

            val seqs = mutableListOf<List<Int>>(nums)

            while (!seqs.last().all { it == 0 }) {
                seqs.add(seqs.last().zipWithNext() { a, b -> b - a })
            }

            seqs.sumOf { it.last() }
        }
    }

    override fun part2(): Int {
        return inputPath.readLines().sumOf {
            val nums = it.split(" ").map { it.toInt() }

            val seqs = mutableListOf(nums)

            while (!seqs.last().all { it == 0 }) {
                seqs.add(seqs.last().zipWithNext() { a, b -> b - a })
            }

            val initial = 0 // Fixes overload ambiguity error ?!
            seqs.foldRight(initial) { a, acc -> a.first() - acc }
        }
    }
}

val testInputD9 = """0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
"""
