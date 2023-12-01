package aoc2023.days

import java.net.http.*
import kotlin.io.path.readText

val testPart1 = """1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet

"""


class Day1 : BaseDay(1) {
    var input1 = ""

    override fun parse() {
        input1 = inputPath.readText()
    }

    override fun part1(): Int {
        val reader = input1.reader()

        var out = 0

        for (l in reader.readLines()) {
            if (l.isNullOrEmpty()) continue

            var firstIdx = -1

            for (i in 0..l.length - 1) {
                val n = l[i].digitToIntOrNull() // ?.let { first = it }

                if (n != null) {
                    firstIdx = i
                    out += n * 10
                    break
                }
            }
            for (i in (l.length - 1) downTo firstIdx) {
                val n = l[i].digitToIntOrNull()

                if (n != null) {
                    out += n
                    break
                }
            }
        }

        return out
    }

    override fun part2(): Int {
        return "TODO".length
    }
}
