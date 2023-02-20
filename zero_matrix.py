func setZeroes(matrix [][]int)  {
    r := len(matrix)
    c := len(matrix[0])

    rows := make(map[int]bool)
    cols := make(map[int]bool)

    // keep track of all rows and cols with a zero
    for i := 0; i < r; i++ {
        for j := 0; j < c; j++ {
            if matrix[i][j] == 0 {
                rows[i] = true
                cols[j] = true
            }
        }
    }

    // iterate and find all the places where zeros are set
    for i := 0; i < r; i++ {
        for j := 0; j < c; j++ {
            if rows[i] || cols[j] {
                matrix[i][j] = 0
            }
        }
    }
}
