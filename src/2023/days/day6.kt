package aoc2023.days

import kotlin.io.path.readLines

data class RaceInfo(var time: Int, var distance: Int)

class Day6 : BaseDay(6) {
    val input = mutableListOf<RaceInfo>()

    override fun parse() {
        val lines = inputPath.readLines()

        input.addAll(digitsRegex.findAll(lines[0]).map { RaceInfo(it.value.toInt(), 0) })

        for ((i, it) in digitsRegex.findAll(lines[1]).withIndex()) {
            input[i].distance = it.value.toInt()
        }
    }

    override fun part1(): Int {
        // d < n(t - n) (where n is time holding the button)

        return input
                .map { raceInfo ->
                    (1..raceInfo.time - 1).count { raceInfo.distance < it * (raceInfo.time - it) }
                }
                .reduce { prev, curr -> prev * curr }
    }

    override fun part2(): String {

        return "TODO"
    }
}

val testInputD6 = """
Time:      7  15   30
Distance:  9  40  200
"""
