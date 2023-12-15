package aoc2023.days

import kotlin.io.path.*

class Day10 : BaseDay(10) {
    var maze = listOf<String>()
    var height = 0
    var width = 0
    var start = Point(0, 0)
    var mainLoop = hashSetOf<Point>()

    override fun parse() {
        // maze = testInputD10.reader().readLines()
        maze = inputPath.readLines()
        height = maze.size - 1
        width = maze[0].length - 1

        for ((y, line) in maze.withIndex()) {
            for ((x, c) in line.withIndex()) {
                if (c == 'S') {
                    start = Point(x, y)
                    return
                }
            }
        }
    }

    private fun getDirections(c: Char): List<Point> {
        return when (c) {
            '|' -> listOf(Point(0, 1), Point(0, -1)) // N-S
            '-' -> listOf(Point(1, 0), Point(-1, 0)) // E-W
            'L' -> listOf(Point(0, -1), Point(1, 0)) // N-E
            'J' -> listOf(Point(0, -1), Point(-1, 0)) // N-W
            '7' -> listOf(Point(0, 1), Point(-1, 0)) // S-W
            'F' -> listOf(Point(0, 1), Point(1, 0)) // S-E
            '.' -> listOf() // No pipe
            'S' -> { // Start
                // Find initial paths
                start.getAdjacents(width, height)
                        .filter {
                            val dirs = getNextPoints(it)
                            dirs.any { it == start }
                        }
                        .map { it - start }
            }
            else -> error("wooot")
        }
    }

    private fun getNextPoints(curr: Point): List<Point> {
        val c = maze[curr.y][curr.x]
        return getDirections(c).map { it + curr }
    }

    override fun part1(): Int {
        var curr = start
        val visited = hashSetOf<Point>(start)
        var steps = 0

        while (true) {
            val next =
                    getNextPoints(curr).filterOutOfBounds(width, height).filter { it !in visited }

            steps++

            if (next.isEmpty()) {
                mainLoop = visited // Assign main loop for part2
                return steps / 2
            }

            curr = next.first()
            visited.add(curr)
        }
    }

    override fun part2(): Any {
        val expanded = hashMapOf<Point, Char>()
        val expandedFactor = 3

        // Add +2 padding
        height += 2
        width += 2

        for ((x, y) in mainLoop) {
            val newPoint = Point((x + 1) * expandedFactor, (y + 1) * expandedFactor)

            expanded += newPoint to '#'
            expanded += getDirections(maze[y][x]).map { newPoint + it to '#' }
        }

        val queue = ArrayDeque<Point>()
        queue.add(Point(0, 0)) // Assume (0,0) is outside the loop

        // Flood fill from outside
        while (queue.isNotEmpty()) {
            val p = queue.removeFirst()

            expanded[p] = '~'

            queue +=
                    p.getAdjacents(width * expandedFactor, height * expandedFactor, true).filter {
                        it !in expanded && it !in queue
                    }
        }

        var enclosed = 0
        for ((y, line) in maze.withIndex()) {
            for ((x, _) in line.withIndex()) {
                enclosed += if (Point(x * 3, y * 3) !in expanded) 1 else 0
            }
        }

        // Print maze
        // Path("output10.txt").bufferedWriter().use {
        //     for (y in 0..height * expandedFactor) {
        //         for (x in 0..width * expandedFactor) {
        //             val c = if (Point(x, y) in expanded) expanded[Point(x, y)] else ' '
        //             it.write(c.toString())
        //         }
        //         it.newLine()
        //     }
        // }

        return enclosed
    }
}

val testInputD10 =
        """.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...
"""
