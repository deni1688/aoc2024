import java.io.File

import kotlin.math.absoluteValue

val input = File("day2/input.txt").readLines()

fun isSafeWithExcluded(report: List<Int>, index: Int, badLevels: Int) = isSafe(
    report.filterIndexed() { i, _ -> i != index },
    badLevels
).or(
    isSafe(
        report.filterIndexed() { i, _ -> i != index - 1 },
        badLevels
    )
).or(
    isSafe(
        report.filterIndexed() { i, _ -> i != report.size - 1 },
        badLevels
    )
)

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
            return isSafeWithExcluded(report, index, badLevels+1)
        }

        if (!isLast && prev < current && report.last() < current) {
            return isSafeWithExcluded(report, index, badLevels+1)
        }

        if (!isLast && prev > current && report.last() > current) {
            return isSafeWithExcluded(report, index, badLevels+1)
        }
    }

    return true;
}

val reports = input.map { it.split(" ").map(String::toInt) }
val result = reports.count(::isSafe)

println("$result")

