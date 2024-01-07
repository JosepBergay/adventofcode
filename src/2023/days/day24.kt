package aoc2023.days

import kotlin.io.path.readLines

class Day24 : BaseDay(24) {
    var input = mutableListOf<Pair<Point3, Point3>>()
    val testArea = 200000000000000.0..400000000000000.0
    // val testArea = 7.0..27.0

    override fun parse() {
        val lines = inputPath.readLines()
        // val lines = testInputD24.reader().readLines()

        for (line in lines) {
            val (pos, vel) =
                    line.split(" @ ").map {
                        val (x, y, z) = it.split(", ")
                        Point3(x.toLong(), y.toLong(), z.toLong())
                    }
            input.add(pos to vel)
        }
    }

    override fun part1(): Int {
        var count = 0

        for ((i, a) in input.withIndex()) {
            val (p1, v1) = a
            for ((_, b) in input.withIndex().drop(i + 1)) {
                val (p2, v2) = b

                // y = vy/vx*x + c
                // vy1/vx1 * x + c1 = vy2/vx2 * x + c2

                val dydx1 = v1.y.toDouble() / v1.x
                val dydx2 = v2.y.toDouble() / v2.x

                if (dydx1 == dydx2) continue // Parallel

                // c = y - dx
                val c1 = p1.y - dydx1 * p1.x
                val c2 = p2.y - dydx2 * p2.x

                val x = (c2 - c1) / (dydx1 - dydx2)
                val y = dydx1 * x + c1

                if (x !in testArea || y !in testArea) continue

                val isPast1 =
                        if (v1.x < 0) (x - p1.x) > 0
                        else (x - p1.x) < 0 || if (v1.y < 0) (y - p1.y) > 0 else (y - p1.y) < 0
                val isPast2 =
                        if (v2.x < 0) (x - p2.x) > 0
                        else (x - p2.x) < 0 || if (v2.y < 0) (y - p2.y) > 0 else (y - p2.y) < 0

                if (isPast1 || isPast2) continue

                count++
            }
        }

        return count
    }

    override fun part2(): Any {

        return "TODO"
    }
}

val testInputD24 =
        """19, 13, 30 @ -2, 1, -2
18, 19, 22 @ -1, -1, -2
20, 25, 34 @ -2, -2, -4
12, 31, 28 @ -1, -2, -1
20, 19, 15 @ 1, -5, -3
"""
