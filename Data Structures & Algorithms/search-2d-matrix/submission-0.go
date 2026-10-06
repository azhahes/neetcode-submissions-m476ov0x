func searchMatrix(matrix [][]int, target int) bool {
    r, c := len(matrix), len(matrix[0])
    l, r := 0 , (r*c)-1
    for l<r{
        mid := l + (r-l)/2
        i, j := mid/c, mid%c
        if matrix[i][j] >= target{
            r = mid
        } else {
            l = mid+1
        }
    }

    i, j := l/c, l%c
    if matrix[i][j] == target{
        return true
    }
    return false
}
