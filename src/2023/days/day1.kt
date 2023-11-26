package aoc2023.days

import java.net.http.*
import kotlin.io.path.readText

class Day1 : BaseDay(1) {
    var input: Int = 0

    override fun parse() {
        input = inputPath.readText().length
    }

    override fun part1(): String {
        return input.toString()
    }

    override fun part2(): Int {
        return "TODO".length
    }
}
