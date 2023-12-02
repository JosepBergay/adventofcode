package aoc2023.days

fun getDay(day: Int): BaseDay? {
    return when (day) {
        1 -> Day1()
        2 -> Day2()
        else -> null
    }
}
