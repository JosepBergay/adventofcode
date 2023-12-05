package aoc2023.days

import kotlin.io.path.*

data class SeedMappingState(
        var mapped: MutableList<Long>,
        var remaining: MutableList<Long>,
        var stage: Int
)

class Day5 : BaseDay(5) {
    var input = listOf<String>()

    override fun parse() {
        input = inputPath.readLines()
        // input = testInputD5.reader().readLines()
    }

    override fun part1(): Long {
        val digitsRegex = """\d+""".toRegex()

        val seeds = digitsRegex.findAll(input[0]).map { it.value.toLong() }.toMutableList()

        val initialState = SeedMappingState(mutableListOf(), seeds, 0)

        val out =
                input.drop(1).fold(initialState) { state, curr ->
                    if (curr.isEmpty()) {
                        // Next stage
                        state.stage++
                        state.remaining += state.mapped
                        state.mapped = mutableListOf()

                        return@fold state
                    } else if (!curr[0].isDigit()) {
                        // source-to-destination map:
                        return@fold state
                    }

                    val range = digitsRegex.findAll(curr).map { it.value.toLong() }.toList()

                    val destinationStart = range[0]
                    val sourceStart = range[1]
                    val len = range[2]
                    val offset = sourceStart - destinationStart

                    val groups = state.remaining.groupBy { it in sourceStart..(sourceStart + len) }

                    val inRange = groups.get(true) ?: mutableListOf()

                    state.mapped.addAll(inRange.toMutableList().map { it - offset }.toMutableList())
                    state.remaining.removeAll { it in inRange }

                    state
                }

        return (out.remaining + out.mapped).min()
    }

    override fun part2(): String {

        return "TODO"
    }
}

val testInputD5 =
        """seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4

"""
