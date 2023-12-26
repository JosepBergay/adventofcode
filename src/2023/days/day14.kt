package aoc2023.days

import kotlin.io.path.readLines
import kotlin.text.reversed

class Day14 : BaseDay(14) {
    var input = listOf<String>()

    override fun parse() {
        input = inputPath.readLines()
        // input = testInputD14.reader().readLines()
    }

    override fun part1(): Int {
        return (0..input[0].length - 1).sumOf {
            var sum = 0
            var emptyIdx = -1

            for (i in 0..input.size - 1) {
                when (input[i][it]) {
                    '#' -> {
                        emptyIdx = -1
                    }
                    '.' -> {
                        if (emptyIdx == -1) {
                            emptyIdx = i
                        }
                    }
                    'O' -> {
                        if (emptyIdx == -1) {
                            sum += input.size - i
                        } else {
                            sum += input.size - emptyIdx
                            emptyIdx++
                        }
                    }
                }
            }

            sum
        }
    }

    private fun rollWest(lines: List<String>): List<String> {
        var out = mutableListOf<String>()

        for (y in 0..lines.size - 1) {
            var str = lines[y]
            var emptyIdx = -1

            for (x in 0..lines[y].length - 1) {
                when (lines[y][x]) {
                    '#' -> {
                        emptyIdx = -1
                    }
                    '.' -> {
                        if (emptyIdx == -1) {
                            emptyIdx = x
                        }
                    }
                    'O' -> {
                        if (emptyIdx != -1) {
                            // Swap chars
                            str =
                                    str.replaceRange(emptyIdx, emptyIdx + 1, "O")
                                            .replaceRange(x, x + 1, ".")

                            emptyIdx++
                        }
                    }
                }
            }

            out.add(str)
        }

        return out
    }

    private fun spinCycle(lines: List<String>): List<String> {
        var tmp = rollWest(lines.getColumns()).getColumns() // Roll North

        tmp = rollWest(tmp)

        tmp =
                rollWest(tmp.getColumns().map { it.reversed() })
                        .map { it.reversed() }
                        .getColumns() // Roll South

        tmp = rollWest(tmp.map { it.reversed() }).map { it.reversed() } // Roll East

        return tmp
    }

    private fun getNorthSupportLoad(lines: List<String>): Int {
        return lines.withIndex().sumOf { (i, line) -> line.count { it == 'O' } * (lines.size - i) }
    }

    override fun part2(): Int {
        val cache = hashMapOf<List<String>, Int>()
        var curr = input
        cache.put(input, 0)

        val max = 1_000_000_000

        for (i in 1..max) {
            curr = spinCycle(curr)

            if (curr in cache) {
                val cycleStart = cache.get(curr)!!
                val cycleLength = i - cycleStart
                val maxIdx = ((max - cycleStart) % cycleLength) + cycleStart

                val key = cache.keys.find { cache[it] == maxIdx }

                return getNorthSupportLoad(key!!)
            } else {
                cache.put(curr, i)
            }
        }

        // No cycle?!
        // return getNorthSupportLoad(curr)
        error("No cycle found, uncomment line above!")
    }
}

val testInputD14 =
        """O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....
"""
