def data = []
new File('../input.txt').eachLine { line ->
    def row = []
    line.each {
        row.add(it as int)
    }
    data.add(row)
}

def getDominant = { arr, col ->
    def currCol = arr.transpose()[col]
    def zeros = currCol.count(0)
    def ones = currCol.count(1)
    return zeros > ones ? 0 : 1
}

def determineRating = { arr, boolean isDominant ->
    def helper = arr
    for (def i = 0; i < arr[0].size(); i++) {
        def tmp = []
        def dominant = getDominant(helper, i)
        dominant = isDominant ? dominant : 1 - dominant
        helper.each {
            if (it[i] == dominant) {
                tmp.add(it)
            }
        }
        helper = tmp
        if (helper.size() <= 1) {
            return helper[0]
        }
    }
}

def oxygen = determineRating(data, true)
def scrubber = determineRating(data, false)

int o = Integer.parseInt(oxygen.join(), 2)
int s = Integer.parseInt(scrubber.join(), 2)
println o*s
