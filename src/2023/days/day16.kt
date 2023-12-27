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

    private fun energizeTiles(initial: Pair<Point, Point>): Int {
        val seen = hashSetOf<Pair<Point, Point>>() // (Position, Direction)

        val queue = ArrayDeque<Pair<Point, Point>>()
        queue.add(initial)

        while (queue.isNotEmpty()) {
            val curr = queue.removeFirst()

            if (curr in seen) continue

            seen.add(curr)

            getMoves(curr.first, curr.second)
                    .map { it.first + it.second to it.second }
                    .filter { it.first.isNotOutOfBounds(width) }
                    .forEach { queue.addLast(it) }
        }

        return seen.distinctBy { it.first }.size
    }

    override fun part1(): Int {
        return energizeTiles(Point(0, 0) to Point(1, 0))
    }

    override fun part2(): Int {
        val initial = buildList {
            addAll((0..width).map { Point(it, 0) to Point(0, 1) }) // topRow
            addAll((0..width).map { Point(it, height) to Point(0, -1) }) // botRow
            addAll((0..height).map { Point(0, it) to Point(1, 0) }) // leftCol
            addAll((0..height).map { Point(width, it) to Point(-1, 0) }) // rightCol
        }

        return initial.maxOf { energizeTiles(it) }
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
