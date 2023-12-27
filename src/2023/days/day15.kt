package aoc2023.days

data class LensBox(val id: Int, var lenses: MutableMap<String, Int> = mutableMapOf())

class Day15 : BaseDay(15) {
    var input = listOf<String>()
    val boxes = buildList<LensBox>(256) { addAll((0..255).map { LensBox(it) }) }

    override fun parse() {
        // input = testInputD15.reader().readText().split(",")
        input = inputPath.toFile().readText().dropLast(1).split(",")
    }

    private fun applyHash(str: String): Int {
        return str.fold(0) { acc, c -> ((acc + c.code) * 17) % 256 }
    }

    override fun part1(): Int {
        return input.sumOf { applyHash(it) }
    }

    private fun applyProcedure(str: String) {
        val isDash = str.endsWith('-')
        val label = if (isDash) str.dropLast(1) else str.dropLast(2)
        val hash = applyHash(label)
        val box = boxes[hash]

        if (isDash) {
            box.lenses.remove(label)
        } else {
            box.lenses[label] = str.last().digitToInt()
        }
    }

    override fun part2(): Long {
        for (str in input) {
            applyProcedure(str)
        }

        val focusingPower =
                boxes.sumOf { box ->
                    box.lenses.values.foldIndexed(0L) { slot, acc, focalLength ->
                        acc + (1 + box.id) * (slot + 1) * focalLength
                    }
                }

        return focusingPower
    }
}

val testInputD15 = """rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"""
