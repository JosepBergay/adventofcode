package aoc2023.days

import kotlin.io.path.readLines

data class PartRating(val x: Int, val m: Int, val a: Int, val s: Int) {
    fun getRating(): Int {
        return x + m + a + s
    }
}

data class Rule(
        val condition: ((PartRating) -> Boolean),
        val nextWf: String,
        val op: Char,
        val letter: Char,
        val operand: Int
)

private typealias Workflow = List<Rule>

class Day19 : BaseDay(19) {
    private val workflows = mutableMapOf<String, Workflow>()
    private val ratings = mutableListOf<PartRating>()

    private fun parseRule(rawRule: String): Rule {
        var split = rawRule.split(":")

        if (split.size == 1) {
            val func =
                    (fun(_: PartRating): Boolean {
                        return true
                    })
            val nextWf = split[0]

            return Rule(func, nextWf, '*', '*', -1)
        }

        val cond = split[0]
        val nextWf = split[1]

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

        return Rule(func, nextWf, op, firstLetter, operand)
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

                workflows[name] = rawRules.dropLast(1).split(",").map { parseRule(it) }
            }
        }
    }

    private fun execWorkflows(rating: PartRating): String {
        var curr = "in"

        while (curr != "A" && curr != "R") {
            val wf = workflows[curr]!!

            for ((condition, nextWf) in wf) {
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

    private fun solveP2Rec(
            wfName: String,
            state: HashMap<Char, IntRange>
    ): List<HashMap<Char, IntRange>> {
        if (wfName == "A") return listOf(state)

        if (wfName == "R") return emptyList()

        val out = mutableListOf<HashMap<Char, IntRange>>()

        val wf = workflows[wfName]!!

        for (rule in wf) {
            if (rule.letter !in state) {
                // Terminal state '*'
                out.addAll(solveP2Rec(rule.nextWf, state))
                continue
            }

            val range = state[rule.letter]!!
            val newState = HashMap(state)

            when (rule.op) {
                '<' -> {
                    newState[rule.letter] = range.first ..< rule.operand

                    state[rule.letter] = rule.operand..range.last
                }
                '>' -> {
                    newState[rule.letter] = rule.operand + 1..range.last

                    state[rule.letter] = range.first..rule.operand
                }
            }

            out.addAll(solveP2Rec(rule.nextWf, newState))
        }

        return out.toList()
    }

    override fun part2(): Long {
        val initialWorkflow = "in"
        val initialState = hashMapOf('x' to 1..4000, 'm' to 1..4000, 'a' to 1..4000, 's' to 1..4000)

        val accepted = solveP2Rec(initialWorkflow, initialState)

        return accepted.sumOf { it.values.fold(1L) { acc, curr -> acc * curr.count().toLong() } }
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
