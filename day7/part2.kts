import java.io.File

val lines = File("input.txt").readLines()

fun evaluateExpression(numbers: List<Long>, operators: List<String>): Long {
    var result = numbers[0]
    for (i in operators.indices) {
        result = when (operators[i]) {
            "+" -> result + numbers[i + 1]
            "*" -> result * numbers[i + 1]
            "||" -> (result.toString() + numbers[i + 1].toString()).toLong()
            else -> result
        }
    }
    return result
}

fun generateOperatorCombinations(length: Int): List<List<String>> {
    if (length == 0) return listOf(emptyList())
    val smallerCombinations = generateOperatorCombinations(length - 1)
    val result = mutableListOf<List<String>>()
    for (comb in smallerCombinations) {
        result.add(comb + "+")
        result.add(comb + "*")
        result.add(comb + "||")
    }
    return result
}

fun calculateCalibrationResult(input: List<String>): Long {
    var totalCalibrationResult = 0L

    for (line in input) {
        val parts = line.split(":")
        val target = parts[0].trim().toLong()
        val numbers = parts[1].trim().split(" ").map { it.toLong() }

        val operatorCombinations = generateOperatorCombinations(numbers.size - 1)
        var isSolvable = false

        for (operators in operatorCombinations) {
            if (evaluateExpression(numbers, operators) == target) {
                isSolvable = true
                break
            }
        }

        if (isSolvable) {
            totalCalibrationResult += target
        }
    }

    return totalCalibrationResult
}


val result = calculateCalibrationResult(lines)
println("Total Calibration Result: $result")
