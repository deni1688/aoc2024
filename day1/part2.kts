import java.io.File

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
val result = columns.first().fold(0) {
    acc, value -> acc + value * columns.last().count {it == value}
}

println("$result")
