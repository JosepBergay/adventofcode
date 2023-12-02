package aoc2023.days

import kotlin.io.path.readLines

typealias Cube = String

data class GameSet(val cubes: Map<Cube, Int>)

data class Game(val id: Int, val sets: Collection<GameSet>)

class Day2 : BaseDay(2) {
    var games: List<Game> = listOf()

    override fun parse() {
        games =
                inputPath.readLines().filter { !it.isNullOrEmpty() }.mapIndexed { idx, it ->
                    val sets =
                            it.split(":")[1].split(";").map {
                                val pairs =
                                        it.split(",").map {
                                            val cubesStr = it.trim().split(" ")

                                            cubesStr[1] to cubesStr[0].toInt()
                                        }

                                GameSet(pairs.toMap())
                            }

                    Game(idx + 1, sets)
                }
    }

    override fun part1(): Int {
        val bag = mapOf("red" to 12, "green" to 13, "blue" to 14)

        return games.sumOf {
            val valid = it.sets.all { it.cubes.all { it.key in bag && it.value <= bag[it.key]!! } }

            if (valid) it.id else 0
        }
    }

    override fun part2(): Int {
        return games.sumOf {
            val setPower = mutableMapOf("red" to 0, "green" to 0, "blue" to 0)

            for (s in it.sets) {
                for (p in setPower) {
                    if (p.key in s.cubes && s.cubes[p.key]!! > p.value) {
                        setPower[p.key] = s.cubes[p.key]!!
                    }
                }
            }

            setPower.values.reduce { acc, v -> acc * v }
        }
    }
}
