package aoc2023.days

fun getDay(day: Int): BaseDay? {
    return when (day) {
        1 -> Day1()
        else -> null
    }
}
