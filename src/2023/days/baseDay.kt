package aoc2023.days

import java.net.URI
import java.net.http.*
import kotlin.io.path.Path

interface IBaseDay {
    fun parse(): Unit

    fun part1(): Any

    fun part2(): Any
}

abstract class BaseDay(val day: Int) : IBaseDay {
    init {
        if (day <= 0) {
            throw ExceptionInInitializerError("Day must be greater than 0")
        }
    }

    val inputPath = Path("./days/inputs/day" + day + ".txt")

    fun fetchInput(url: String, cookie: String, year: Int) {
        val segments = listOf(url, year, "day", day, "input")

        val fullUrl = segments.joinToString("/")

        val uri = URI.create(fullUrl)

        val request =
                HttpRequest.newBuilder().uri(uri).setHeader("cookie", "session=" + cookie).build()

        val client = HttpClient.newBuilder().build()

        val response = client.send(request, HttpResponse.BodyHandlers.ofFile(inputPath))

        if (response.statusCode() != 200) {
            throw Exception("HTTP ERROR" + response.toString())
        }
    }

    fun exec(): DayResult {
        parse()

        return DayResult(part1(), part2())
    }
}

data class DayResult(val part1: Any, val part2: Any)
