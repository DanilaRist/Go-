package sprint

func BetweenLimits(from, to rune) string {
    var start, end rune
    if from < to {
        start = from
        end = to
    } else {
        start = to
        end = from
    }

    var result []rune
    for r := start + 1; r < end; r++ {
        result = append(result, r)
    }
    return string(result)
}