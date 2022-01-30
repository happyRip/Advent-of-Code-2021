class Diagnostic {
    static void main(String[] args) {
        def data = []
        new File('../test.txt').eachLine { line ->
            data.add(line)
        }

        int[][] count = [[0] * 2] * data[0].size()
        data.each {
            it.eachWithIndex { v, i ->
                count[i][v as int]++
            }
        }

        def (gamma, epsilon) = ['', '']
        count.each {
            def (zero, one) = [it[0], it[1]]
            if (zero > one) {
                gamma += 0
                epsilon += 1
            } else {
                gamma += 1
                epsilon += 0
            }
        }

        int g = Integer.parseInt(gamma, 2)
        int e = Integer.parseInt(epsilon, 2)
        println g*e
    }
}
