import java.io.File

val lines = parse(File("input.txt").readLines())

var total: Long = 0

for (i in 0..lines.size - 1) {
    val row = lines[i]
    val target = row[0]

    var result = check(row, "+");

    if (result == target) {
        println("Row $i is solved with +")
        total += target
        continue
    }

    result = check(row, "*");

    if (result == target) {
        println("Row $i is solved with *")
        total += target
        continue
    }

    result = check(row, "+*");

    if (result == target) {
        println("Row $i is solved with +*")
        total += target
        continue
    }

    result = check(row, "*+");

    if (result == target) {
        println("Row $i is solved with *+")
        total += target
        continue
    }

    result = checkRecursive(target, row, "+", 0);

    if (result == target) {
        println("Row $i is solved with + recursive")
        total += target
        continue
    }

    result = checkRecursive(target, row, "*", 0);

    if (result == target) {
        println("Row $i is solved with * recursive")
        total += target
    }
}

fun checkRecursive(target: Long, row: List<Long>, operator: String, index: Int): Long {
    var result: Long = 0

    for (j in 1..row.size - 1) {
        if (j == 1) {
            result = row[j]
        } else {
            if (operator == "+") {
                if (j < index) {
                    result += row[j]
                } else {
                    result *= row[j]
                }
            } else {
                if (j < index) {
                    result *= row[j]
                } else {
                    result += row[j]
                }
            }
        }
    }

    if(result == target || index == row.size - 1) {
        return result
    }

    return checkRecursive(target, row, operator, index + 1)
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
