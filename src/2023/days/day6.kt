package aoc2023.days

import kotlin.io.path.readLines

data class RaceInfo(var time: Int, var distance: Int)

class Day6 : BaseDay(6) {
    val times = mutableListOf<String>()
    val distances = mutableListOf<String>()

    override fun parse() {
        val lines = inputPath.readLines()

        times.addAll(digitsRegex.findAll(lines[0]).map { it.value })
        distances.addAll(digitsRegex.findAll(lines[1]).map { it.value })
    }

    private fun countWinningRaces(time: Long, distance: Long): Int {
        // d < n(t - n) (where n is time holding the button)
        return (1..time - 1).count { distance < it * (time - it) }
    }

    override fun part1(): Int {
        return times
                .mapIndexed { i, it -> countWinningRaces(it.toLong(), distances[i].toLong()) }
                .reduce { prev, curr -> prev * curr }
    }

    override fun part2(): Int {
        val time = times.reduce { prev, curr -> prev + curr }.let { it.toLong() }
        val distance = distances.reduce { prev, curr -> prev + curr }.let { it.toLong() }

        return countWinningRaces(time, distance)
    }
}

val testInputD6 = """
Time:      7  15   30
Distance:  9  40  200
"""
