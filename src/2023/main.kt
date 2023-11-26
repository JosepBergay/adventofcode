package aoc2023

import aoc2023.days.*
import java.net.http.*
import kotlin.io.path.*
import kotlin.time.measureTimedValue

val url = "https://adventofcode.com"
// "https://adventofcode.com/2022/day/%v/input"

val year = 2022

fun main(args: Array<String>) {
    val cookie = System.getenv("SESSION_COOKIE")

    if (cookie.isNullOrEmpty()) {
        println("No cookie :<")
        return
    }

    for (arg in args) {
        // var res = arg.toIntOrNull()?.let { getDay(it) }?.run { exec(url, cookie, year) }
        val i = arg.toIntOrNull()

        if (i == null || i < 1 || i > 25) {
            println("Invalid arg: $arg")
            continue
        }

        val day = getDay(i)

        if (day == null) {
            println("Day $i not implemented")
            continue
        }

        if (day.inputPath.notExists()) {
            day.fetchInput(url, cookie, year)
        }

        val (res, elapsed) = measureTimedValue { day.exec() }

        println("Day $i:\t[Part1]: ${res.part1} [Part2]: ${res.part2} ($elapsed)")
    }
}
