import java.io.File

import kotlin.math.absoluteValue

val sorted_columns = File("day1/input.txt").readLines()
    .fold(
        mutableListOf<MutableList<Int>>(
            mutableListOf(),
            mutableListOf(),
        )
    ) { acc, line ->
        line.split("   ").forEachIndexed { index, value ->
            acc[index].add(value.toInt())
        }
        acc
    }
    .map { it.sorted() }

val result = sorted_columns.first().zip(sorted_columns.last()) { a, b -> (a - b).absoluteValue }.sum()

println("$result")
