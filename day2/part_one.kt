package days.two

import kotlin.math.absoluteValue


fun partOne(lines: List<String>): Int {
    val reports = lines.map { it.split(" ").map(String::toInt) }

    return reports.count(::isSafe)
}

private fun isSafe(report: List<Int>): Boolean {
    for (index in 1..<report.size) {
        val prev = report[index - 1]
        val current = report[index]
        val rateOfChange = (prev - current).absoluteValue
        val isLast = index == report.size - 1

        if (rateOfChange !in 1..3) return false

        if (!isLast && prev < current && report.last() < current) {
            return false
        }

        if (!isLast && prev > current && report.last() > current) {
            return false
        }
    }

    return true;
}

