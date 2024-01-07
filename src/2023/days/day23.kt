package aoc2023.days

import kotlin.collections.hashMapOf
import kotlin.io.path.readLines

class Day23 : BaseDay(23) {
    var input = mutableMapOf<Point, Char>()
    var start = Point(0, 0)
    var end = Point(0, 0)

    override fun parse() {
        val lines = inputPath.readLines()
        // val lines = testInputD23.reader().readLines()

        for ((y, line) in lines.withIndex()) {
            for ((x, c) in line.withIndex()) {
                if (c == '#') continue

                input[Point(x, y)] = c
            }
        }
        start = input.keys.find { it.y == 0 }!!
        end = input.keys.maxBy { it.y }
    }

    val stack = ArrayDeque<Point>()
    val visited = hashSetOf<Point>()

    private fun topologicalSort(curr: Point) {
        visited += curr

        for (dir in Direction.entries) {
            val p = curr + dir.p

            if (p !in input || p in visited) continue

            val c = input[p]

            if ((c == '>' && dir == Direction.EAST) ||
                            (c == '^' && dir == Direction.North) ||
                            (c == 'v' && dir == Direction.South) ||
                            (c == '<' && dir == Direction.WEST) ||
                            c == '.'
            )
                    topologicalSort(p)
        }

        stack.addLast(curr)
    }

    override fun part1(): Any {
        val distances = hashMapOf<Point, Int>()

        for (p in input.keys) {
            distances[p] = Int.MIN_VALUE

            if (p in visited) continue

            topologicalSort(p)
        }

        distances[start] = 0

        while (stack.isNotEmpty()) {
            val curr = stack.removeLast()

            if (distances[curr]!! != Int.MIN_VALUE) {
                for (adj in Direction.entries.map { it.p + curr }.filter { it in input }) {
                    if (distances[adj]!! < distances[curr]!! + 1) {
                        distances[adj] = distances[curr]!! + 1
                    }
                }
            }
        }

        return distances[end]!!
    }

    override fun part2(): Any {

        return "TODO"
    }
}

val testInputD23 =
        """#.#####################
#.......#########...###
#######.#########.#.###
###.....#.>.>.###.#.###
###v#####.#v#.###.#.###
###.>...#.#.#.....#...#
###v###.#.#.#########.#
###...#.#.#.......#...#
#####.#.#.#######.#.###
#.....#.#.#.......#...#
#.#####.#.#.#########v#
#.#...#...#...###...>.#
#.#.#v#######v###.###v#
#...#.>.#...>.>.#.###.#
#####v#.#.###v#.#.###.#
#.....#...#...#.#.#...#
#.#########.###.#.#.###
#...###...#...#...#.###
###.###.#.###v#####v###
#...#...#.#.>.>.#.>.###
#.###.###.#.###.#.#v###
#.....###...###...#...#
#####################.#
"""
