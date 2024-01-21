package aoc2023.days

import kotlin.collections.hashMapOf
import kotlin.io.path.readLines

class Day25 : BaseDay(25) {
    val input = hashMapOf<String, HashSet<String>>()

    override fun parse() {
        val lines = inputPath.readLines()
        // val lines = testInputD25.reader().readLines()

        for (line in lines) {
            val (first, rest) = line.split(": ")
            val connections = rest.split(" ")

            input.getOrPut(first) { hashSetOf() } += connections
            for (c in connections) {
                input.getOrPut(c) { hashSetOf() } += first
            }
        }
    }

    private fun walkGraph(
            graph: Map<String, Set<String>>,
            start: String,
            counter: MutableMap<Set<String>, Int>?
    ): Int {
        val seen = hashSetOf<String>()
        val q = ArrayDeque<String>()

        q += start
        seen += start

        while (q.isNotEmpty()) {
            val curr = q.removeFirst()

            for (connection in graph[curr]!!) {
                if (connection in seen) continue

                seen += connection
                q += connection

                if (counter != null) {
                    val s = setOf(curr, connection)
                    if (s !in counter) {
                        counter[s] = 1
                    } else {
                        counter[s] = counter[s]!! + 1
                    }
                }
            }
        }

        return seen.size
    }

    private fun removeConnections(vararg edges: Set<String>): Map<String, Set<String>> {
        return input.mapValues { (k, v) ->
            val edge = edges.filter { k in it }.flatMap { it.filter { it != k } }

            v - edge
        }
    }

    override fun part1(): Int {
        val counter = mutableMapOf<Set<String>, Int>()
        for (start in input.keys) {
            walkGraph(input, start, counter)
        }

        val edges = counter.entries.sortedByDescending { it.value }.map { it.key }

        for (i in 0 ..< edges.size) {
            for (j in i + 1 ..< edges.size) {
                for (k in j + 1 ..< edges.size) {
                    val cutGraph = removeConnections(edges[i], edges[j], edges[k])

                    val seen = walkGraph(cutGraph, input.keys.first(), null)

                    if (seen < input.size) {
                        return seen * (input.size - seen)
                    }
                }
            }
        }

        return -1
    }

    override fun part2(): Any {

        return "TODO"
    }
}

val testInputD25 =
        """jqt: rhn xhk nvd
rsh: frs pzl lsr
xhk: hfx
cmg: qnr nvd lhk bvb
rhn: xhk bvb hfx
bvb: xhk hfx
pzl: lsr hfx nvd
qnr: nvd
ntq: jqt hfx bvb xhk
nvd: lhk
lsr: lhk
rzs: qnr cmg lsr rsh
frs: qnr lhk lsr
"""
