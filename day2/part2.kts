import java.io.File

import kotlin.math.absoluteValue

val input = File("./input.txt").readLines()

var badLevels = 0;

fun isSafe(report: List<Int>): Boolean {
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

val reports = input.map { it.split(" ").map(String::toInt) }
val result = reports.count(::isSafe)

println("$result")
