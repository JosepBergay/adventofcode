package aoc2023.days

import kotlin.io.path.*
import kotlin.math.*

data class DigInstruction(val dir: Direction, val count: Int, val color: String)

class Day18 : BaseDay(18) {
    var input = listOf<DigInstruction>()

    override fun parse() {
        val lines = inputPath.readLines()
        // val lines = testInputD18.reader().readLines()

        input =
                lines.map {
                    val (dir, count, color) = it.split(" ")

                    val direction =
                            when (dir) {
                                "U" -> Direction.North
                                "R" -> Direction.EAST
                                "D" -> Direction.South
                                "L" -> Direction.WEST
                                else -> error("lies!")
                            }

                    DigInstruction(direction, count.toInt(), color.drop(1).dropLast(1))
                }
    }

    override fun part1(): Any {
        var minX = 0
        var minY = 0
        var maxX = 0
        var maxY = 0
        val trench = hashSetOf<Point>()

        var curr = Point(0, 0)
        trench += curr

        for (instr in input) {
            for (i in (1..instr.count)) {
                curr = instr.dir.p + curr
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

    override fun part2(): Any {
        return "TODO"
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
