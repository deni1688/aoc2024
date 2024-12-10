import java.io.File

val columns = File("day1/input.txt").readLines().fold(
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

val result = columns.first().fold(0) { acc, value ->
    acc + value * columns.last().count { it == value }
}

println("$result")
