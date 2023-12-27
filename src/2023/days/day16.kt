package aoc2023.days

import kotlin.io.path.readLines

class Day16 : BaseDay(16) {
    val cave = hashMapOf<Point, Char>()
    var height = 0
    var width = 0

    override fun parse() {
        // val lines = testInputD16.reader().readLines()
        val lines = inputPath.readLines()

        height = lines.size - 1
        width = lines[0].length - 1

        for ((y, line) in lines.withIndex()) {
            for ((x, c) in line.withIndex()) {
                cave[Point(x, y)] = c
            }
        }
    }

    private fun getMoves(position: Point, direction: Point): List<Pair<Point, Point>> {
        return when (cave[position]) {
            '.' -> listOf(position to direction)
            '/' ->
                    when (direction) {
                        Point(1, 0) -> listOf(position to Point(0, -1))
                        Point(-1, 0) -> listOf(position to Point(0, 1))
                        Point(0, 1) -> listOf(position to Point(-1, 0))
                        Point(0, -1) -> listOf(position to Point(1, 0))
                        else -> error("lies / -> $direction")
                    }
            '\\' ->
                    when (direction) {
                        Point(1, 0) -> listOf(position to Point(0, 1))
                        Point(-1, 0) -> listOf(position to Point(0, -1))
                        Point(0, 1) -> listOf(position to Point(1, 0))
                        Point(0, -1) -> listOf(position to Point(-1, 0))
                        else -> error("lies \\ -> $direction")
                    }
            '|' ->
                    when (direction) {
                        Point(1, 0), Point(-1, 0) ->
                                listOf(position to Point(0, 1), position to Point(0, -1))
                        Point(0, 1), Point(0, -1) -> listOf(position to direction)
                        else -> error("lies | -> $direction")
                    }
            '-' ->
                    when (direction) {
                        Point(1, 0), Point(-1, 0) -> listOf(position to direction)
                        Point(0, 1), Point(0, -1) ->
                                listOf(position to Point(1, 0), position to Point(-1, 0))
                        else -> error("lies - -> $direction")
                    }
            else -> error("lies ${cave[position]}")
        }
    }

    override fun part1(): Int {
        val seen = hashSetOf<Pair<Point, Point>>() // (Position, Direction)

        val queue = ArrayDeque<Pair<Point, Point>>()
        queue.add(Point(0, 0) to Point(1, 0))

        while (queue.isNotEmpty()) {
            val curr = queue.removeFirst()

            if (curr in seen) continue

            seen.add(curr)

            getMoves(curr.first, curr.second)
                    .map { it.first + it.second to it.second }
                    // .also { println("pos $curr.first dir $curr.second: moves -> $it") }
                    .filter { it.first.isNotOutOfBounds(width) && it !in seen }
                    .forEach { queue.addLast(it) }
        }

        return seen.distinctBy { it.first }.size
    }

    override fun part2(): Any {

        return "TODO"
    }
}

val testInputD16 =
        """.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....
"""
