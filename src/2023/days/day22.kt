package aoc2023.days

import kotlin.collections.hashMapOf
import kotlin.io.path.readLines

class Day22 : BaseDay(22) {
    var input = listOf<Volume>()
    // Brick -> bricks above
    val above = hashMapOf<Volume, HashSet<Volume>>()
    // Brick -> bricks below
    val below = hashMapOf<Volume, HashSet<Volume>>()

    override fun parse() {
        val lines = inputPath.readLines()
        // val lines = testInputD22.reader().readLines()

        input =
                lines
                        .map {
                            val (start, end) =
                                    it.split("~").map {
                                        it.split(",").map { it.toInt() }.let { (x, y, z) ->
                                            Triple(x, y, z)
                                        }
                                    }

                            Volume(
                                    start.first..end.first,
                                    start.second..end.second,
                                    start.third..end.third
                            )
                        }
                        .sortedBy { it.z.first }

        val stack = mutableListOf<Volume>()

        // Let it sand!!
        for (brick in input) {
            val overlap = stack.filter { it.hasIntersectionXY(brick) }.maxByOrNull { it.z.last }

            val moved =
                    if (overlap == null) Volume(brick.x, brick.y, 1..brick.z.count())
                    else
                            Volume(
                                    brick.x,
                                    brick.y,
                                    (overlap.z.last + 1)..(overlap.z.last + brick.z.count())
                            )

            stack.add(moved)
        }

        input = stack

        for (brick in input) {
            above[brick] =
                    input
                            .filter {
                                brick.z.last + 1 == it.z.first && it.hasIntersectionXY(brick)
                            }
                            .toHashSet()

            below[brick] =
                    input
                            .filter {
                                brick.z.first - 1 == it.z.last && it.hasIntersectionXY(brick)
                            }
                            .toHashSet()
        }
    }

    override fun part1(): Int {
        return input.count { brick -> above[brick]!!.all { below[it]!!.any { it != brick } } }
    }

    override fun part2(): Int {
        var wouldFallCount = 0

        for (brick in input) {
            val falling = hashSetOf<Volume>()

            val q = ArrayDeque<Volume>()
            q += brick

            while (q.isNotEmpty()) {
                val curr = q.removeFirst()

                falling += curr

                val wouldFall = above[curr]!!.filter { below[it]!!.all { it in falling } }

                q += wouldFall
            }

            wouldFallCount += falling.size - 1
        }

        return wouldFallCount
    }
}

val testInputD22 =
        """1,0,1~1,2,1
0,0,2~2,0,2
0,2,3~2,2,3
0,0,4~0,2,4
2,0,5~2,2,5
0,1,6~2,1,6
1,1,8~1,1,9
5,4,6~5,4,9
5,4,12~5,4,15
"""
