import java.io.File

val args: List<String> = emptyList()
val useSample = args.contains("sample")
val input = File("day6/input${if (useSample) "_sample" else ""}.txt").readLines().map { it.split("").toMutableList() }.toMutableList()
var guard = "^"

fun getRow(guard: String, input: MutableList<MutableList<String>>): Int = input.indexOfFirst { guard in it }
fun getCol(guard: String, input: MutableList<MutableList<String>>): Int = input[getRow(guard, input)].indexOf(guard)


val direction = { s: String ->
    when (s) {
        "^" -> Pair(-1, 0)
        "v" -> Pair(1, 0)
        "<" -> Pair(0, -1)
        ">" -> Pair(0, 1)
        else -> throw IllegalArgumentException("Invalid direction")
    }
}

while (true) {
    val row = getRow(guard, input)
    val col = getCol(guard, input)
    val current = input[row][col]
    val dir = direction(current)

    if (row + dir.first !in input.indices || col + dir.second !in input[row].indices) {
        break
    }


    val right = when (current) {
        "^" -> ">"
        ">" -> "v"
        "v" -> "<"
        else -> "^"
    }

    val next = input[row + dir.first][col + dir.second]

    if (next == "#") {
        input[row][col] = right
        guard = right
    } else {
        input[row][col] = "o"
        input[row + dir.first][col + dir.second] = current
    }

    print("\r")
    print(input.joinToString("\n") { it.joinToString("") })

    Thread.sleep(100)
}





