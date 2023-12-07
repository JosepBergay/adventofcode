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

    override fun parse() {
        // lines = testInputD7.reader().readLines().dropLast(1)
        lines = inputPath.readLines()
    }

    override fun part1(): Any {
        val allCardsMap = mutableMapOf<String, CamelCardHand>()

        val camelCardTypeComparator =
                compareBy<String> {
                    allCardsMap
                            .getOrPut(it) {
                                val (cards, bid) = it.split(" ")

                                val count =
                                        buildMap<Char, Int> {
                                            cards.forEach { merge(it, 1) { a, b -> a + b } }
                                        }

                                val type =
                                        when (count.size) {
                                            1 -> HandType.REPOKER
                                            2 ->
                                                    if (count.values.any { it == 4 }) HandType.POKER
                                                    else HandType.FULL
                                            3 ->
                                                    if (count.values.any { it == 3 }) HandType.TRIO
                                                    else HandType.DOUBLE_PAIR
                                            4 -> HandType.PAIR
                                            5 -> HandType.HIGH_CARD
                                            else -> error("woot")
                                        }

                                CamelCardHand(type, cards, bid, count)
                            }
                            .type
                            .ordinal
                }

        val cardStrengthComparator =
                compareBy<String> {
                    allCardsMap.get(it)!!
                            .cards
                            .fold("") { acc, curr ->
                                acc +
                                        when (curr) {
                                            'A' -> 'F'
                                            'K' -> 'E'
                                            'Q' -> 'D'
                                            'J' -> 'C'
                                            'T' -> 'B'
                                            else -> curr
                                        }
                            }
                            .toInt(16)
                }

        val sorted = lines.sortedWith(camelCardTypeComparator.then(cardStrengthComparator))

        return sorted.withIndex().sumOf { (idx, curr) ->
            allCardsMap.get(curr)!!.bid.toInt() * (idx + 1)
        }
    }

    override fun part2(): Any {

        return "TODO"
    }
}

val testInputD7 = """32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
"""
