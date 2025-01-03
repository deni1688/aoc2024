import java.io.File

import kotlin.math.absoluteValue

val input = File("day2/input.txt").readLines()

fun <T> Iterable<T>.withoutItemAt(index: Int): List<T> =
    filterIndexed { i, _ -> i != index }

fun isSafeWithExcluded(report: List<Int>, index: Int, badLevels: Int): Boolean =
    isSafe(report.withoutItemAt(index), badLevels + 1)
            || isSafe(report.withoutItemAt(index - 1), badLevels + 1)
            || isSafe(report.withoutItemAt(report.size - 1), badLevels + 1)

fun isSafe(report: List<Int>, badLevels: Int = 0): Boolean {
    if (badLevels > 1) {
        return false
    }

    for (index in 1..<report.size) {
        val prev = report[index - 1]
        val current = report[index]
        val rateOfChange = (prev - current).absoluteValue
        val isLast = index == report.size - 1

        if (rateOfChange !in 1..3) {
            return isSafeWithExcluded(report, index, badLevels)
        }

        if (!isLast && prev < current && report.last() < current) {
            return isSafeWithExcluded(report, index, badLevels)
        }

        if (!isLast && prev > current && report.last() > current) {
            return isSafeWithExcluded(report, index, badLevels)
        }
    }

    return true
}

val reports = input.map { it.split(" ").map(String::toInt) }
val result = reports.count(::isSafe)

println("$result")

