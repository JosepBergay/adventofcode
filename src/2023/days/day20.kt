package aoc2023.days

import kotlin.io.path.readLines

data class Comm(
        val type: String,
        val name: String,
        val destinations: List<String>,
        // Conjunction state -> true for 'high', false for 'low'
        var receivedPulses: HashMap<String, Boolean>,
        // Flip-flop state -> true for 'on', false for 'off'
        var state: Boolean
)

class Day20 : BaseDay(20) {
    var input = mapOf<String, Comm>()

    override fun parse() {
        val lines = inputPath.readLines()
        // val lines = testInputD20.reader().readLines()

        input =
                lines
                        .map {
                            val (name, destinations) = it.split(" -> ")

                            val typeName =
                                    if (name == "broadcaster") name to name
                                    else name.first().toString() to name.drop(1)

                            Comm(
                                    typeName.first,
                                    typeName.second,
                                    destinations.split(", "),
                                    hashMapOf(),
                                    false
                            )
                        }
                        .associateBy { it.name }

        // Initialize Conjunctions receiver map
        for (comm in input.values) {
            for (dest in comm.destinations) {
                if (input[dest]?.type != "&") continue

                input[dest]!!.receivedPulses[comm.name] = false
            }
        }
    }

    private fun processPulse(
            from: String,
            dest: String,
            pulse: Boolean
    ): Triple<String, Boolean, List<String>> {
        if (dest !in input) return Triple("", pulse, emptyList())

        val comm = input[dest]!!

        when (comm.type) {
            "%" -> {
                if (!pulse) {
                    comm.state = !comm.state
                    return Triple(dest, comm.state, comm.destinations)
                }
            }
            "&" -> {
                comm.receivedPulses[from] = pulse
                val allHighPulses = comm.receivedPulses.values.all { it }

                return Triple(dest, !allHighPulses, comm.destinations)
            }
            "broadcaster" -> {
                return Triple(dest, pulse, comm.destinations)
            }
        }

        // Send nothing
        return Triple("", pulse, emptyList())
    }

    var cycleMap = mutableMapOf<String, Long>()

    private fun pushButtonOnce(count: Long = -1L): Pair<Long, Long> {
        // Send low pulse to broadcaster.
        // Triple(source, pulse, destinations)
        var curr = listOf(Triple("button", false, listOf("broadcaster")))

        var lowCount = 0L
        var highCount = 0L

        while (curr.isNotEmpty()) {
            curr =
                    curr.flatMap { (source, pulse, destinations) ->
                        destinations.map {
                            if (pulse) highCount++ else lowCount++

                            processPulse(source, it, pulse)
                        }
                    }

            if (count == -1L) continue

            // For part2
            for ((source, pulse) in curr) {
                if (!pulse) continue

                if (source in cycleMap && cycleMap[source] == 0L) {
                    // High pulse delivered to $source @ $count
                    cycleMap[source] = count
                }
            }
        }

        return lowCount to highCount
    }

    override fun part1(): Long {
        var low = 0L
        var high = 0L

        repeat(1000) {
            val (l, h) = pushButtonOnce()
            low += l
            high += h
        }

        return low * high
    }

    override fun part2(): Long {
        // Reset all modules to their default states.
        parse()

        val parent = input.values.find { "rx" in it.destinations }

        if (parent == null) return -1

        cycleMap = parent.receivedPulses.keys.associate { it to 0L }.toMutableMap()

        var count = 0L
        while (cycleMap.any { it.value == 0L }) {
            count++
            pushButtonOnce(count)
        }

        return lcm(cycleMap.values.distinct())
    }
}

// val testInputD20 = """broadcaster -> a, b, c
// %a -> b
// %b -> c
// %c -> inv
// &inv -> a
// """

val testInputD20 = """broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output
"""
