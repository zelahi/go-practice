func rotate(matrix [][]int)  {
    startRow := 0
    startCol := 0
    endRow := len(matrix) - 1
    endCol := len(matrix[0]) - 1

    for startRow <= endRow && startCol <= endCol {
        for i := 0; i < endCol - startCol; i++ {
            tmp := matrix[startRow][startCol + i]
            matrix[startRow][startCol + i] = matrix[endRow - i][startCol]
            matrix[endRow - i][startCol] = matrix[endRow][endCol - i]
            matrix[endRow][endCol - i] = matrix[startRow + i][endCol]
            matrix[startRow + i][endCol] = tmp
        }

        startRow ++
        startCol ++
        endRow --
        endCol --
    }
}
