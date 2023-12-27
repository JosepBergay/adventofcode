package aoc2023.days

fun getDay(day: Int): BaseDay? {
    return when (day) {
        1 -> Day1()
        2 -> Day2()
        3 -> Day3()
        4 -> Day4()
        5 -> Day5()
        6 -> Day6()
        7 -> Day7()
        8 -> Day8()
        9 -> Day9()
        10 -> Day10()
        11 -> Day11()
        12 -> Day12()
        13 -> Day13()
        14 -> Day14()
        15 -> Day15()
        16 -> Day16()
        else -> null
    }
}
