import java.io.File

val lines = parse(File("input.txt").readLines())

var total: Long = 0

for (i in 0..lines.size-1) {
    val row = lines[i]
    val target = row[0]

    var result = check(row, "+");

    if (result == target) {
        println("Row $i is solved with +")
        total+=target
        continue
    }

    result = check(row, "*");

    if (result == target) {
        println("Row $i is solved with *")
        total+=target
        continue
    }

    result = check(row, "+*");

    if (result == target) {
        println("Row $i is solved with +*")
        total+=target
        continue
    }

    result = check(row, "*+");

    if (result == target) {
        println("Row $i is solved with *+")
        total+=target
        continue
    }
}

println("Total: $total")

fun parse(input: List<String>) = input.map(String::trim).map {
    it.split(": ")
        .mapIndexed() { index, s -> if (index == 1) s.split(" ") else listOf(s) }
        .flatten()
        .map(String::toLong)
}


fun check(row: List<Long>, operator: String): Long {
    var result: Long = 0
    var nextOp = operator.first().toString()
    for (j in 1..row.size - 1) {
        if (j == 1) {
            result = row[j]
        } else {
            when (operator) {
                "+" -> result += row[j]
                "*" -> result *= row[j]
                "+*" -> {
                    if (nextOp == "+") {
                        result += row[j]
                        nextOp = "*"
                    } else {
                        result *= row[j]
                        nextOp = "+"
                    }
                }
                "*+" -> {
                    if (nextOp == "*") {
                        result *= row[j]
                        nextOp = "+"
                    } else {
                        result += row[j]
                        nextOp = "*"
                    }
                }
            }
        }
    }
    return result
}
