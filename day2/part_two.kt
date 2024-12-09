package days.two

import kotlin.math.absoluteValue


fun partTwo(lines: List<String>): Int {
    val reports = lines.map { it.split(" ").map(String::toInt) }

    return reports.count(::isSafe)
}

var badLevels = 0;

private fun isSafe(report: List<Int>): Boolean {
    var result = true
    for (index in 1..<report.size) {
        val prev = report[index - 1]
        val current = report[index]
        val rateOfChange = (prev - current).absoluteValue
        val isLast = index == report.size - 1

        if (rateOfChange !in 1..3) {
            result = false
        }

        if (!isLast && prev < current && report.last() < current) {
            result = false
        }

        if (!isLast && prev > current && report.last() > current) {
            result = false
        }

        if(!result && badLevels < 1) {
            badLevels++
            result = isSafe(report.filterIndexed { i, _ -> i != index })
                    .or(isSafe(report.filterIndexed { i, _ -> i != index - 1 }))
        }
    }

    return result
}

