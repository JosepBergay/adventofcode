package aoc2023.days

import java.util.PriorityQueue
import kotlin.io.path.readLines

data class CityBlock(val p: Point, val lastDir: Direction?, val dirCount: Int, val heatLoss: Int) {
    fun getDirections(width: Int, height: Int, dirCountRange: IntRange): List<Direction> {
        return Direction.entries
                .filter { lastDir == null || it.p != Point(-lastDir.p.x, -lastDir.p.y) }
                .filter {
                    lastDir == null ||
                            if (it == lastDir) dirCount < dirCountRange.last
                            else dirCount >= dirCountRange.first
                }
                .filter { (it.p + p).isNotOutOfBounds(width, height) }
    }
}

class Day17 : BaseDay(17) {
    val input = hashMapOf<Point, Int>()
    var height = 0
    var width = 0

    override fun parse() {
        val lines = inputPath.readLines()
        // val lines = testInputD17.reader().readLines()

        height = lines.size - 1
        width = lines[0].length - 1

        for ((y, line) in lines.withIndex()) {
            for ((x, n) in line.withIndex()) {
                input[Point(x, y)] = n.digitToInt()
            }
        }
    }

    private fun dijsktra(dirCountRange: IntRange): Int {
        val start = Point(0, 0)
        val end = Point(width, height)

        val queue = java.util.PriorityQueue<CityBlock>(compareBy { it.heatLoss })
        val visited = hashSetOf<Triple<Point, Direction?, Int>>()

        queue += CityBlock(start, null, 0, 0)

        while (queue.isNotEmpty()) {
            val curr = queue.poll()

            if (curr.p == end && curr.dirCount >= dirCountRange.first) return curr.heatLoss

            val hash = Triple(curr.p, curr.lastDir, curr.dirCount)
            if (hash in visited) continue
            visited += hash

            for (dir in curr.getDirections(width, height, dirCountRange)) {
                val nextPoint = dir.p + curr.p
                val dirCount = if (dir == curr.lastDir) curr.dirCount + 1 else 1
                val heatLoss = curr.heatLoss + input[nextPoint]!!
                queue += CityBlock(nextPoint, dir, dirCount, heatLoss)
            }
        }

        return -1
    }

    override fun part1(): Int {
        return dijsktra(0..3)
    }

    override fun part2(): Int {
        return dijsktra(4..10)
    }
}

val testInputD17 =
        """2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533
"""
