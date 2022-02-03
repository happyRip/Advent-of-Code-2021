def data = []
new File('../input.txt').eachLine {
    line -> data.add(line as int)
}

def triples = []
data[0..-3].eachWithIndex { _, i ->
    triples.add(data[i..i+2].sum())
}

def result = 0
triples.eachWithIndex { v, i ->
    if (i == 0) {
        return
    } else if (v > triples[i-1]) {
        result++
    }
}
println result