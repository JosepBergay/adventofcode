package aoc2023.days

import java.net.http.*
import kotlin.io.path.*

val testPart1 = """1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet

"""

val testPart2 =
        """two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
43fourthreesjldzsonesix3ndm
4zxrtfz
"""

class Day1 : BaseDay(1) {
    var input1 = ""
    var input2 = ""

    override fun parse() {
        // input1 = inputPath.readText()
        // input2 = testPart2
    }

    override fun part1(): Int {
        var out = 0

        for (l in inputPath.readLines()) {
            if (l.isNullOrEmpty()) continue

            var firstIdx = -1

            for (i in 0..l.length - 1) {
                val n = l[i].digitToIntOrNull()

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
        var out = 0

        for (line in inputPath.readLines()) {
            // println("line $l out $out")
            if (line.isNullOrEmpty()) continue

            for (window in line.windowed(5, partialWindows = true)) {
                if (window[0].isDigit()) {
                    out += window[0].digitToInt() * 10
                    break
                }

                val v = window.startsWithDigitStr()
                if (v != -1) {
                    out += (v + 1) * 10
                    break
                }
            }

            for (window in line.reversed().windowed(5, partialWindows = true)) {
                if (window.first().isDigit()) {
                    out += window.first().digitToInt()
                    break
                }

                val v = window.endsWithDigitStr()
                if (v != -1) {
                    out += (v + 1)
                    break
                }
            }
        }

        return out
    }
}

val digitsAsStrings =
        arrayOf(
                "one",
                "two",
                "three",
                "four",
                "five",
                "six",
                "seven",
                "eight",
                "nine",
        )

/** It is only guaranteed to work if String is less than 6 characters long. */
fun String.startsWithDigitStr(): Int {
    for ((idx, str) in digitsAsStrings.withIndex()) {
        if (this.startsWith(str)) {
            return idx
        }
    }

    return -1
}

/** It is only guaranteed to work if String is less than 6 characters long. */
fun String.endsWithDigitStr(): Int {
    for ((idx, str) in digitsAsStrings.withIndex()) {
        if (this.startsWith(str.reversed())) {
            return idx
        }
    }

    return -1
}
