def data = []
new File('../input.txt').eachLine {
    line -> data.add(line)
}

def result = 0
data.eachWithIndex { v, i ->
    if (i == 0) {
        return
    }
    if (v > data[i-1]) {
        result++
    }
}
println result