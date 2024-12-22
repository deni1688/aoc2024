import java.io.File

val lines = parse(File("input_sample.txt").readLines())

lines.forEach(::println)

fun parse(input: List<String>) = input.map(String::trim).map {
    it.split(": ")
        .mapIndexed() { index, s -> if (index == 1) s.split(" ") else listOf(s) }
        .flatten()
        .map(String::toInt)
}
