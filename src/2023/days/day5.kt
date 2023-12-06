package aoc2023.days

import kotlin.io.path.*

data class SeedMappingState<T>(
        var mapped: MutableList<T>,
        var remaining: MutableList<T>,
        var stage: Int
)

val digitsRegex = """\d+""".toRegex()

class Day5 : BaseDay(5) {
    var input = listOf<String>()
    var initialSeeds = emptySequence<Long>()

    override fun parse() {
        input = inputPath.readLines()
        // input = testInputD5.reader().readLines()

        initialSeeds = digitsRegex.findAll(input[0]).map { it.value.toLong() }
    }

    private fun <T> runAlmanac(
            initialState: SeedMappingState<T>,
            groupFn: (state: SeedMappingState<T>, validRange: LongRange) -> Pair<List<T>, List<T>>,
            applyOffset: (item: T, offset: Long) -> T
    ): SeedMappingState<T> {
        return input.drop(1).fold(initialState) { state, curr ->
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

            val (destinationStart, sourceStart, len) =
                    digitsRegex.findAll(curr).map { it.value.toLong() }.toList()

            val offset = sourceStart - destinationStart
            val validRange = sourceStart..sourceStart + len

            val (inRange, outOfRange) = groupFn(state, validRange)

            state.mapped.addAll(inRange.map { applyOffset(it, offset) })
            state.remaining = outOfRange.toMutableList()

            state
        }
    }

    override fun part1(): Long {
        val initialState = SeedMappingState(mutableListOf(), initialSeeds.toMutableList(), 0)

        return runAlmanac(
                        initialState,
                        groupFn = { s, validRange -> s.remaining.partition { it in validRange } },
                        applyOffset = { it, offset -> it - offset }
                )
                .let { (it.remaining + it.mapped).min() }
    }

    override fun part2(): Long {
        val seedRanges = initialSeeds.chunked(2).map { it[0]..(it[0] + it[1]) }.toMutableList()

        val initialState = SeedMappingState(mutableListOf(), seedRanges, 0)

        return runAlmanac(
                        initialState,
                        groupFn = { s, validRange ->
                            val initial = mutableListOf<LongRange>() to mutableListOf<LongRange>()

                            s.remaining.fold(initial) { (inR, outR), curr ->
                                when {
                                    // Totally out of range (no overlap)
                                    curr.last < validRange.first || validRange.last < curr.first ->
                                            inR to outR.apply { add(curr) }
                                    // Fully in range (valid contains current)
                                    validRange.first <= curr.first &&
                                            curr.last <= validRange.last ->
                                            inR.apply { add(curr) } to outR
                                    // Partially in range (current range contains valid)
                                    curr.first < validRange.first && validRange.last < curr.last ->
                                            inR.apply { add(validRange) } to
                                                    outR.apply {
                                                        add(curr.first ..< validRange.first)
                                                        add(validRange.last + 1..curr.last)
                                                    }
                                    // Partially in range (current range higher bound is in valid)
                                    curr.first < validRange.first ->
                                            inR.apply { add(validRange.first..curr.last) } to
                                                    outR.apply { curr.first ..< validRange.first }
                                    // Partially in range (current range lower bound is in valid)
                                    else ->
                                            inR.apply { curr.first..validRange.last } to
                                                    outR.apply {
                                                        add(validRange.last + 1..curr.last)
                                                    }
                                }
                            }
                        },
                        applyOffset = { it, offset -> (it.first - offset)..(it.last - offset) }
                )
                .let { (it.remaining + it.mapped).map { it.first }.min() }
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
