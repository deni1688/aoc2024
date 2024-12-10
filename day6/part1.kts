import java.io.File

val args: List<String> = emptyList()
val useSample = args.contains("sample")
val input = File("day6/input${if (useSample)"_sample" else ""}.txt").readLines()

println(input)
