package aoc2023.days

import kotlin.io.path.*
import kotlin.math.*

data class DigInstruction(val part1: Pair<Direction, Int>, val part2: Pair<Direction, Int>)

class Day18 : BaseDay(18) {
    var input = listOf<DigInstruction>()

    override fun parse() {
        val lines = inputPath.readLines()
        // val lines = testInputD18.reader().readLines()

        input =
                lines.map {
                    val (dir, count, color) = it.split(" ")

                    val part1 =
                            when (dir) {
                                "U" -> Direction.North
                                "R" -> Direction.EAST
                                "D" -> Direction.South
                                "L" -> Direction.WEST
                                else -> error("lies!")
                            } to count.toInt()

                    val part2 =
                            when (color.dropLast(1).last()) {
                                '0' -> Direction.EAST
                                '1' -> Direction.South
                                '2' -> Direction.WEST
                                '3' -> Direction.North
                                else -> error("woot")
                            } to color.drop(2).dropLast(2).toInt(16)

                    DigInstruction(part1, part2)
                }
    }

    private fun part1FloodFill(): Int {
        // Innefficient
        var minX = 0
        var minY = 0
        var maxX = 0
        var maxY = 0
        val trench = hashSetOf<Point>()

        var curr = Point(0, 0)
        trench += curr

        for ((part1, _) in input) {
            val (dir, count) = part1
            for (i in (1..count)) {
                curr = dir.p + curr
                trench += curr

                maxX = max(maxX, curr.x)
                maxY = max(maxY, curr.y)
                minX = min(minX, curr.x)
                minY = min(minY, curr.y)
            }
        }

        // Add +1 padding
        minX--
        minY--
        maxX++
        maxY++

        val queue = ArrayDeque<Point>()
        val visited = hashSetOf<Point>()
        queue.add(Point(minX, minY))

        while (queue.isNotEmpty()) {
            val aux = queue.removeFirst()

            if (aux in visited) continue
            visited += aux

            queue +=
                    Direction.entries
                            .map { it.p + aux }
                            .filter { it.x in minX..maxX && it.y in minY..maxY }
                            .filter { it !in trench }
        }

        return (maxX - minX + 1) * (maxY - minY + 1) - visited.size
    }

    private fun shoelace(points: List<Point>, perimeter: Int): Long {
        val p = buildList {
            addAll(points)
            add(points.first())
        }

        return p.windowed(2).sumOf { (a, b) ->
            a.x.toLong() * b.y.toLong() - b.x.toLong() * a.y.toLong()
        } / 2 + perimeter / 2 + 1
    }

    private fun solve(input: List<Pair<Direction, Int>>): Long {
        val points = mutableListOf<Point>()
        var perimeter = 0

        var prev = Point(0, 0)
        points += prev

        for ((dir, count) in input) {
            prev = Point(dir.p.x * count, dir.p.y * count) + prev
            perimeter += count
            points += prev
        }

        return shoelace(points, perimeter)
    }

    override fun part1(): Long {
        return solve(input.map { it.part1 })
    }

    override fun part2(): Long {
        return solve(input.map { it.part2 })
    }
}

val testInputD18 =
        """R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)
"""
