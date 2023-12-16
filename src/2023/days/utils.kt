package aoc2023.days

import kotlin.math.*

data class Point(val x: Int, val y: Int)

operator fun Point.plus(other: Point) = Point(x + other.x, y + other.y)

operator fun Point.minus(other: Point) = Point(x - other.x, y - other.y)

fun Point.manhattan(other: Point) = abs(x - other.x) + abs(y - other.y)

fun List<Point>.filterOutOfBounds(width: Int, height: Int = width): List<Point> {
    return this.filter { it.x in 0..width && it.y in 0..height }
}

fun Point.getAdjacents(
        width: Int,
        height: Int = width,
        withDiagonals: Boolean = false
): List<Point> {
    val directions = mutableListOf(Point(1, 0), Point(-1, 0), Point(0, 1), Point(0, -1))

    if (withDiagonals) {
        val diagonals = listOf(Point(1, 1), Point(-1, -1), Point(-1, 1), Point(1, -1))
        directions.addAll(diagonals)
    }

    return directions.map { this + it }.filterOutOfBounds(width, height)
}
