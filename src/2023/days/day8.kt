package aoc2023.days

import kotlin.io.path.readLines

class Day8 : BaseDay(8) {
    var instructions: String = ""
    var network = emptyMap<String, Pair<String, String>>()

    override fun parse() {
        // val lines = testInputD8.reader().readLines()
        val lines = inputPath.readLines()

        instructions = lines[0]
        network =
                lines.drop(2)
                        .map {
                            val split = it.split(" = ")
                            val (left, right) = split[1].drop(1).dropLast(1).split(", ")
                            split[0] to (left to right)
                        }
                        .toMap()
    }

    override fun part1(): Int {
        var curr = "AAA"
        var step = 0

        while (curr != "ZZZ") {
            val dir = instructions[step % instructions.length]

            curr = if (dir == 'L') network[curr]!!.first else network[curr]!!.second

            step++
        }

        return step
    }

    override fun part2(): Long {
        var currentNodes = network.keys.filter { it.endsWith('A') }
        var acc = 1L

        for (i in currentNodes) {
            var curr = i
            var step = 0

            while (!curr.endsWith('Z')) {
                val dir = instructions[step % instructions.length]

                curr = if (dir == 'L') network[curr]!!.first else network[curr]!!.second

                step++
            }

            acc *= step / instructions.length
        }

        return acc * instructions.length
    }
}

val testInputD8 =
        """LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)
"""
