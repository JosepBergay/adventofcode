package aoc2023.days

import kotlin.io.path.readLines

enum class HandType {
    HIGH_CARD, // 0
    PAIR,
    DOUBLE_PAIR,
    TRIO,
    FULL,
    POKER,
    REPOKER, // 6
}

data class CamelCardHand(
        val type: HandType,
        val cards: String,
        val bid: String,
        val count: Map<Char, Int>
)

class Day7 : BaseDay(7) {
    var lines = listOf<String>()
    val allCardsMap = mutableMapOf<String, CamelCardHand>()

    override fun parse() {
        // lines = testInputD7.reader().readLines()
        lines = inputPath.readLines()
    }

    private fun getCardStrengthComparator(isPart1: Boolean): Comparator<String> {
        return compareBy<String> {
            allCardsMap.get(it)!!
                    .cards
                    .fold("") { acc, curr ->
                        acc +
                                when (curr) {
                                    'A' -> 'F'
                                    'K' -> 'E'
                                    'Q' -> 'D'
                                    'J' -> if (isPart1) 'C' else '1'
                                    'T' -> 'B'
                                    else -> curr
                                }
                    }
                    .toInt(16)
        }
    }

    private fun computeWinnings(sorted: List<String>): Int {
        return sorted.withIndex().sumOf { (idx, curr) ->
            allCardsMap.get(curr)!!.bid.toInt() * (idx + 1)
        }
    }

    private fun getHandType(count: Map<Char, Int>): HandType {
        return when (count.size) {
            1 -> HandType.REPOKER
            2 -> if (count.values.any { it == 4 }) HandType.POKER else HandType.FULL
            3 -> if (count.values.any { it == 3 }) HandType.TRIO else HandType.DOUBLE_PAIR
            4 -> HandType.PAIR
            5 -> HandType.HIGH_CARD
            else -> error("woot")
        }
    }

    override fun part1(): Int {
        val camelCardTypeComparator =
                compareBy<String> {
                    allCardsMap
                            .getOrPut(it) {
                                val (cards, bid) = it.split(" ")

                                val count = cards.groupingBy { it }.eachCount()

                                val type = getHandType(count)

                                CamelCardHand(type, cards, bid, count)
                            }
                            .type
                            .ordinal
                }

        val sorted = lines.sortedWith(camelCardTypeComparator.then(getCardStrengthComparator(true)))

        return computeWinnings(sorted)
    }

    override fun part2(): Int {
        val camelCardTypeComparator =
                compareBy<String> {
                    val hand = allCardsMap.get(it)!! // Computed from part1

                    if ('J' in hand.count) {
                        var max =
                                hand.count.filter { it.key != 'J' }.maxByOrNull { it.value }?.let {
                                    it.toPair()
                                }

                        if (max == null) {
                            // JJJJJ
                            max = 'A' to 5
                        }

                        val count =
                                hand.count
                                        .entries
                                        .map { if (it.key == 'J') max else it.toPair() }
                                        .toMap()

                        getHandType(count).ordinal
                    } else {
                        hand.type.ordinal
                    }
                }

        val sorted =
                lines.sortedWith(camelCardTypeComparator.then(getCardStrengthComparator(false)))

        return computeWinnings(sorted)
    }
}

val testInputD7 = """32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
"""
