package aoc2023

import aoc2023.days.*
import java.net.http.*
import kotlin.io.path.*
import kotlin.time.measureTimedValue
import kotlinx.coroutines.*
import kotlinx.coroutines.channels.*

val url = "https://adventofcode.com"
// "https://adventofcode.com/2022/day/%v/input"

val year = 2023

fun main(args: Array<String>) {
    val cookie = System.getenv("SESSION_COOKIE")

    if (cookie.isNullOrEmpty()) {
        println("No cookie :< - You better have inputs in place")
    }

    var days = buildList {
        for (arg in args) {
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

            add(day)
        }
    }

    if (days.isEmpty()) {
        // Run them all
        days = (1..25).mapNotNull { getDay(it) }
    }

    println("Running days ${days.map { it.day }}")

    runBlocking {
        val channel = Channel<String>()

        for (day in days) {
            launch { runDay(day, cookie, channel) }
        }

        var count = 0
        for (msg in channel) {
            println(msg)

            if (days.size <= ++count) {
                channel.close() // or break
            }
        }
    }
}

suspend fun runDay(day: BaseDay, cookie: String, c: SendChannel<String>) {
    if (day.inputPath.notExists()) {
        day.fetchInput(url, cookie, year)
    }

    val (res, elapsed) = measureTimedValue { day.exec() }

    val msg = ("Day ${day.day}:\t[Part1]: ${res.part1} [Part2]: ${res.part2} ($elapsed)")

    c.send(msg)
}
