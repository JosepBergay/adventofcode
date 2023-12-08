package aoc2023.days

import kotlin.io.path.readLines

class Day8 : BaseDay(8) {
    val input = mutableListOf<Int>()

    override fun parse() {
        // for (line in inputPath.readLines()) {
        //     line.toIntOrNull()?.let { input.add(it) }
        // }
    }

    override fun part1(): Any {
        // val lines = testInputD8.reader().readLines()
        val lines = inputPath.readLines()

        val instructions = lines[0]
        val network =
                lines.drop(2)
                        .map {
                            val split = it.split(" = ")
                            val (left, right) = split[1].drop(1).dropLast(1).split(", ")
                            split[0] to (left to right)
                        }
                        .toMap()

        var curr = "AAA"
        var step = 0

        while (curr != "ZZZ") {
            val dir = instructions[step % instructions.length]

            curr = if (dir == 'L') network[curr]!!.first else network[curr]!!.second

            step++
        }

        return step
    }

    override fun part2(): Any {

        return "TODO"
    }
}

val testInputD8 = """LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)
"""
