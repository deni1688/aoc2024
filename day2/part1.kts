import java.io.File

import kotlin.math.absoluteValue

val input = File("day2/input.txt").readLines()

fun isSafe(report: List<Int>): Boolean {
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

val reports = input.map { it.split(" ").map(String::toInt) }
val result = reports.count(::isSafe)

println("$result")
