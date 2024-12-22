import java.io.File

val args: List<String> = emptyList()
val useSample = args.contains("sample")
val positions = mutableSetOf<String>()

val grid = File("day6/input${if (useSample) "_sample" else ""}.txt")
    .readLines()
    .map { it.trim().split("").filter { row -> row.isNotEmpty() }.toMutableList() }
    .toMutableList()
var guard = "^"

while (true) {
    val row = getRow(guard, grid)
    val col = getCol(guard, grid)
    val direction = getDirection(guard)
    val right = getRight(guard)

    positions.add("$row,$col")

    if (nextLeavingGrid(row, direction, col)) break

    val next = grid[row + direction.first][col + direction.second]

    if (next == "#") {
        grid[row][col] = right
        guard = right
    } else {
        grid[row][col] = "o"
        grid[row + direction.first][col + direction.second] = guard
    }
}

println(positions.size)

fun getRow(guard: String, input: MutableList<MutableList<String>>): Int = input.indexOfFirst { guard in it }

fun getCol(guard: String, input: MutableList<MutableList<String>>): Int = input[getRow(guard, input)].indexOf(guard)

fun getDirection(s: String) =
    when (s) {
        "^" -> Pair(-1, 0)
        "v" -> Pair(1, 0)
        "<" -> Pair(0, -1)
        ">" -> Pair(0, 1)
        else -> throw IllegalArgumentException("Invalid direction")
    }

fun getRight(c: String) =
    when (c) {
        "^" -> ">"
        ">" -> "v"
        "v" -> "<"
        else -> "^"
    }

fun nextLeavingGrid(row: Int, direction: Pair<Int, Int>, col: Int) =
    row + direction.first !in grid.indices || col + direction.second !in grid[row].indices
