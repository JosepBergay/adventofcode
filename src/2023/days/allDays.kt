package aoc2023.days

fun getDay(day: Int): BaseDay? {
    return when (day) {
        1 -> Day1()
        2 -> Day2()
        3 -> Day3()
        4 -> Day4()
        else -> null
    }
}
