package aoc2023.days

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

    private fun getValidAdjacents(p: Point): List<Point> {
        return Direction.entries.map { it.p + p }.filter { it in input }
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
                for (adj in getValidAdjacents(curr)) {
                    if (distances[adj]!! < distances[curr]!! + 1) {
                        distances[adj] = distances[curr]!! + 1
                    }
                }
            }
        }

        return distances[end]!!
    }

    private fun makeGraph(): HashMap<Point, MutableList<Pair<Point, Int>>> {
        val vertices = hashMapOf(start to mutableListOf<Pair<Point, Int>>(), end to mutableListOf())

        for (p in input.keys) {
            if (getValidAdjacents(p).count() > 2) {
                vertices[p] = mutableListOf()
            }
        }

        for (v in vertices.keys.toList()) {
            var curr = setOf(v)
            val visited = hashSetOf(v)
            var dist = 0

            while (curr.isNotEmpty()) {
                dist++

                curr = buildSet {
                    for (n in curr) {
                        for (adj in getValidAdjacents(n).filter { it !in visited }) {
                            if (adj in vertices) {
                                vertices[v]!! += adj to dist
                            } else {
                                add(adj)
                                visited += adj
                            }
                        }
                    }
                }
            }
        }

        return vertices
    }

    override fun part2(): Any {
        val graph = makeGraph()

        var acc = 0

        if (graph[start]!!.size == 1) {
            val (n, d) = graph[start]!!.single()
            graph.remove(start)
            for (adj in graph.values) {
                adj.removeIf { start == it.first }
            }
            start = n
            acc += d
        }

        if (graph[end]!!.size == 1) {
            val (n, d) = graph[end]!!.single()
            graph.remove(end)
            for (adj in graph.values) {
                adj.removeIf { end == it.first }
            }
            end = n
            acc += d
        }

        visited.clear()

        fun findMax(p: Point, curr: Int): Int {
            if (p == end) return curr

            visited += p

            val max =
                    graph[p]!!.filter { it.first !in visited }.maxOfOrNull { (n, w) ->
                        findMax(n, curr + w)
                    }

            visited -= p

            return max ?: 0
        }

        return findMax(start, 0) + acc
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
