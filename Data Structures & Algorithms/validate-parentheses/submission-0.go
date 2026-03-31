func isValid(s string) bool {
    dict := map[rune]rune{
        '(': ')',
        '[': ']',
        '{': '}',
    }
    stack := make([]rune, 0)
    for _, r := range []rune(s) {
        if _, ok := dict[r]; ok {
            stack = append(stack, r)
        } else {
            n := len(stack)
            if n == 0 || dict[stack[n-1]] != r {
                return false
            }
            stack = stack[:n-1]
        }
    }
    return len(stack) == 0
}