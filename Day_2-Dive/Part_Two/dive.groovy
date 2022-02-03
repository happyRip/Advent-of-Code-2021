def data = []
new File('../input.txt').eachLine { line -> 
    def (dir, val) = line.tokenize(' ')
    data.add([dir, val as int])
}

def (aim, depth, position) = [0, 0, 0]
data.each { dir, val ->
    switch (dir) {
        case 'down':
            aim += val
            break
        case 'up':
            aim -= val
            break
        case 'forward':
            position += val
            depth += aim * val
            break
    }
}
println "${depth * position}"