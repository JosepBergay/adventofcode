package aoc2023.days

import kotlin.io.path.readLines

class Day10 : BaseDay(10) {
    var maze = mutableListOf<MutableList<Char>>()
    var height = 0
    var width = 0
    var start = Point(0, 0)

    override fun parse() {
        // val lines = testInputD10.reader().readLines()
        val lines = inputPath.readLines()

        height = lines.size - 1

        for ((y, line) in lines.withIndex()) {
            width = line.length - 1
            val aux = mutableListOf<Char>()

            for ((x, c) in line.withIndex()) {
                if (c == 'S') {
                    start = Point(x, y)
                }

                aux.add(c)
            }

            maze.add(aux)
        }
    }

    private fun getDirections(curr: Point): List<Point> {
        val c = maze[curr.y][curr.x]
        return when (c) {
            '|' -> listOf(Point(0, 1), Point(0, -1)).map { it + curr } // N-S
            '-' -> listOf(Point(1, 0), Point(-1, 0)).map { it + curr } // E-W
            'L' -> listOf(Point(0, -1), Point(1, 0)).map { it + curr } // N-E
            'J' -> listOf(Point(0, -1), Point(-1, 0)).map { it + curr } // N-W
            '7' -> listOf(Point(0, 1), Point(-1, 0)).map { it + curr } // S-W
            'F' -> listOf(Point(0, 1), Point(1, 0)).map { it + curr } // S-E
            '.' -> listOf() // No pipe
            'S' -> { // Start
                // Find initial paths
                start.getAdjacents(width, height).filter { p ->
                    val dirs = getDirections(p)
                    dirs.any { it == start }
                }
            }
            else -> error("wooot")
        }
    }

    override fun part1(): Any {
        var curr = start
        val visited = hashMapOf<Point, Boolean>(start to true)
        var steps = 0

        while (true) {
            val next =
                    getDirections(curr).filterOutOfBounds(width, height).filter { it !in visited }

            steps++
            // println("$steps -> $curr -> $next")

            if (next.isEmpty()) return steps / 2

            curr = next.first()
            visited[curr] = true
        }
    }

    override fun part2(): Any {

        return "TODO"
    }
}

val testInputD10 = """..F7.
.FJ|.
SJ.L7
|F--J
LJ...
"""
