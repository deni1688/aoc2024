import java.io.File

val lines = File("input_sample.txt").readLines().map(String::trim).map {
    it.split(": ").mapIndexed() { index, s -> if (index == 1) s.split(" ") else s }
}

lines.forEach(::println)
