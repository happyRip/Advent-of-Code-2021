def data = []
new File('../input.txt').eachLine { line -> 
    def (dir, val) = line.tokenize(' ')
    data.add([dir, val as int])
}

def (depth, position) = [0, 0]
data.each { dir, val ->
    switch (dir) {
        case 'forward':
            position += val
            break
        case 'up':
            depth -= val
            break
        case 'down':
            depth += val
            break
    }
}
println "${depth*position}"