package aoc2023.days

import java.math.BigInteger
import kotlin.io.path.readLines
import kotlin.toBigInteger

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

    private fun determinant(m: List<List<BigInteger>>): BigInteger {
        if (m.isEmpty()) return BigInteger.ONE

        val r =
                m.first().mapIndexed { i, it ->
                    it * determinant(m.drop(1).map { it.filterIndexed { j, _ -> j != i } })
                }

        return r.foldIndexed(BigInteger.ZERO) { i, a, b -> if (i % 2 == 0) a + b else a - b }
    }

    override fun part2(): BigInteger {
        // Inspired by:
        // shahata5 @ https://www.reddit.com/r/adventofcode/comments/18pnycy/comment/khlrstp/

        val A = mutableListOf<List<BigInteger>>()
        val B = mutableListOf<BigInteger>()

        val (p0, v0) = input[0]

        repeat(3) {
            val (pN, vN) = input[it + 1]

            A.add(
                    listOf(
                            v0.y.toBigInteger() - vN.y.toBigInteger(),
                            vN.x.toBigInteger() - v0.x.toBigInteger(),
                            BigInteger.ZERO,
                            pN.y.toBigInteger() - p0.y.toBigInteger(),
                            p0.x.toBigInteger() - pN.x.toBigInteger(),
                            BigInteger.ZERO
                    )
            )
            B.add(
                    p0.x.toBigInteger() * v0.y.toBigInteger() -
                            p0.y.toBigInteger() * v0.x.toBigInteger() -
                            pN.x.toBigInteger() * vN.y.toBigInteger() +
                            pN.y.toBigInteger() * vN.x.toBigInteger()
            )

            A.add(
                    listOf(
                            v0.z.toBigInteger() - vN.z.toBigInteger(),
                            BigInteger.ZERO,
                            vN.x.toBigInteger() - v0.x.toBigInteger(),
                            pN.z.toBigInteger() - p0.z.toBigInteger(),
                            BigInteger.ZERO,
                            p0.x.toBigInteger() - pN.x.toBigInteger()
                    )
            )
            B.add(
                    p0.x.toBigInteger() * v0.z.toBigInteger() -
                            p0.z.toBigInteger() * v0.x.toBigInteger() -
                            pN.x.toBigInteger() * vN.z.toBigInteger() +
                            pN.z.toBigInteger() * vN.x.toBigInteger()
            )
        }

        // Cramer
        val detA = determinant(A)

        val (prx, pry, prz) =
                A.mapIndexed { i, _ ->
                    determinant(
                            A.mapIndexed { j, r ->
                                val l = r.toMutableList()
                                l[i] = B[j]
                                l
                            }
                    ) / detA
                }

        return prx + pry + prz
    }
}

val testInputD24 =
        """19, 13, 30 @ -2, 1, -2
18, 19, 22 @ -1, -1, -2
20, 25, 34 @ -2, -2, -4
12, 31, 28 @ -1, -2, -1
20, 19, 15 @ 1, -5, -3
"""
