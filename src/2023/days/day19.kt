package aoc2023.days

import kotlin.io.path.readLines

data class PartRating(val x: Int, val m: Int, val a: Int, val s: Int) {
    fun getRating(): Int {
        return x + m + a + s
    }
}

data class Workflow(
        val name: String,
        val rules: List<Pair<((PartRating) -> Boolean), String>> // Condition to Next workflow name
)

class Day19 : BaseDay(19) {
    val workflows = mutableMapOf<String, Workflow>()
    val ratings = mutableListOf<PartRating>()

    private fun parseCondition(cond: String): (PartRating) -> Boolean {
        val firstLetter = cond.first()
        val op = cond.drop(1).first()
        val operand = cond.drop(2).toInt()

        val func =
                fun(r: PartRating): Boolean {
                    val v =
                            when (firstLetter) {
                                'x' -> r.x
                                'm' -> r.m
                                'a' -> r.a
                                's' -> r.s
                                else -> error("no xmas!")
                            }

                    return if (op == '>') v > operand else v < operand
                }

        return func
    }

    override fun parse() {
        val lines = inputPath.readLines()
        // val lines = testInputD19.reader().readLines()

        for (l in lines) {
            if (l.startsWith('{')) {
                // Rating
                val (x, m, a, s) = digitsRegex.findAll(l).map { it.value.toInt() }.toList()
                ratings.add(PartRating(x, m, a, s))
            } else if (l.isNotEmpty()) {
                // Workflow
                val (name, rawRules) = l.split("{")

                val rules =
                        rawRules.dropLast(1).split(",").map {
                            var split = it.split(":")

                            if (split.size == 1) {
                                val condition =
                                        (fun(_: PartRating): Boolean {
                                            return true
                                        })
                                condition to split[0]
                            } else {
                                parseCondition(split[0]) to split[1]
                            }
                        }

                workflows[name] = Workflow(name, rules)
            }
        }
    }

    private fun execWorkflows(rating: PartRating): String {
        var curr = "in"

        while (curr != "A" && curr != "R") {
            val wf = workflows[curr]!!

            for ((condition, nextWf) in wf.rules) {
                if (condition(rating)) {
                    curr = nextWf
                    break
                }
            }
        }

        return curr
    }

    override fun part1(): Int {
        return ratings.groupBy { execWorkflows(it) }["A"]?.sumOf { it.getRating() } ?: 0
    }

    override fun part2(): Any {

        return "TODO"
    }
}

val testInputD19 =
        """px{a<2006:qkq,m>2090:A,rfg}
pv{a>1716:R,A}
lnx{m>1548:A,A}
rfg{s<537:gd,x>2440:R,A}
qs{s>3448:A,lnx}
qkq{x<1416:A,crn}
crn{x>2662:A,R}
in{s<1351:px,qqz}
qqz{s>2770:qs,m<1801:hdj,R}
gd{a>3333:R,R}
hdj{m>838:A,pv}

{x=787,m=2655,a=1222,s=2876}
{x=1679,m=44,a=2067,s=496}
{x=2036,m=264,a=79,s=2244}
{x=2461,m=1339,a=466,s=291}
{x=2127,m=1623,a=2188,s=1013
"""
