package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
    count := 0
	if  f, ok := cb[file]; ok {
        for _, boolean := range f {
            if boolean {
                count++
            }
        }
        return count
    } else {
        return 0
    }
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
    count := 0
    if len(cb) < rank || rank < 1 {
        return 0
    }
	for _, v := range cb {
        if v[rank-1] {
            count++
        }
    }
    return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
    count := 0
    for _, k := range cb {
        count += len(k)
    }
    return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0
	for k, _ := range cb {
        count += CountInFile(cb, k)
    }
    return count
}
