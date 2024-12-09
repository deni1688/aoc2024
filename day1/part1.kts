import java.io.File

import kotlin.math.absoluteValue

val input = File("day1/input.txt").readLines()

fun getColumns(lines: List<String>): List<List<Int>> {
    return lines.fold(
        mutableListOf<MutableList<Int>>(
            mutableListOf(),
            mutableListOf(),
        )
    ) {
            acc, line ->

        line.split("   ").forEachIndexed { index, value ->
            acc[index].add(value.toInt())
        }

        acc
    }
}

val columns = getColumns(input)
val sorted = columns.map{it.sorted()}
val result = sorted.first().zip(sorted.last()) { a, b -> (a - b).absoluteValue }.sum()

println("$result")
