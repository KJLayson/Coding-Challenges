package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	slice, exists := cb[file]
    tally := 0
    if exists == true {
        for _,val := range slice {
            if val == true {
                tally += 1
            }
        }
    }

    return tally
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
    tally := 0
    index := rank -1
    if rank >= 1 && rank <= 8 {
        for _, file := range cb {
            if file[index] == true {
                tally += 1
            }
        }
    } 

    return tally
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	tally := 0
    for _,file := range cb {
        for range file {
            tally += 1
        }
    }

    return tally
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	tally := 0
    for _,file := range cb {
        for _, square := range file {
            if square == true {
                tally += 1
            }
        }
    }

    return tally
}
