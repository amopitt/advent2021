package day04

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	// BingoCardCells is the rows/cells of the bingo card
	BingoCardCells = 5
)

type Board struct {
	rowCells [BingoCardCells][BingoCardCells]int
	HasBingo bool
}

// Bingo is the main game
type Bingo struct {
	Boards []*Board
}

func Day4() {
	bingo := new(Bingo)
	bingo.Boards = make([]*Board, 0)
	// read input of bit strings
	moves, err := getBingoInput("day04/Input.txt", bingo)
	if err != nil {
		fmt.Println(err)
		return
	}

	// part 1 - first to get bingo
	winningBoard, moveNumber := bingo.PlayGame(moves, false)
	result := getScore(winningBoard, moveNumber)
	fmt.Printf("result: %d\n", result)

	// part 2 - find last board to get bingo
	winningBoard, moveNumber = bingo.PlayGame(moves, true)
	result = getScore(winningBoard, moveNumber)
	fmt.Printf("result: %d\n", result)
}

func getScore(board *Board, moveNumber int) int {
	sum := 0
	for _, cell := range board.rowCells {
		for _, value := range cell {
			if value != -1 {
				sum += value
			}
		}
	}
	fmt.Printf("move Number: %d, sum: %d\n", moveNumber, sum)
	return sum * moveNumber
}

func (b *Bingo) PlayGame(moves []int, checkAllBoards bool) (*Board, int) {
	isBingo := false
	moveNumber := 0

	// create int array
	playedMoves := make([]int, 0)

	// Draw numbers until a board gets bingo
	for !isBingo {
		move := moves[moveNumber]
		playedMoves = append(playedMoves, move)
		// fmt.Println("Move:", move)

		for _, board := range b.Boards {
			if !board.HasBingo {
				board.CheckBingoNumber(move)

				// if the board doesn't have bingo, check if it's bingo
				if board.IsBingo(playedMoves) {
					// fmt.Println("board got a bingo")
					if checkAllBoards {
						// if all boards have bingo, return the winning board
						// loop over b.boards
						allBingos := true
						for _, ba := range b.Boards {
							if !ba.HasBingo {
								allBingos = false
								break
							}
						}
						if allBingos {
							isBingo = true
							return board, move
						}
					} else {
						isBingo = true
						return board, move
					}

				}
			}
		}
		moveNumber++
	}
	return nil, 0
}

// MarkNumberChecked marks the number as checked with value passed
func (b *Board) CheckBingoNumber(number int) {
	for row, cellsArray := range b.rowCells {
		for cellKey, cellValue := range cellsArray {
			if cellValue == number {
				b.rowCells[row][cellKey] = -1
			}
		}
	}
}

func (b *Board) IsBingo(numbers []int) bool {
	for i := 0; i < len(b.rowCells); i++ {
		var rowMatch, colMatch int
		for j := 0; j < len(b.rowCells); j++ {
			if b.rowCells[i][j] == -1 {
				rowMatch++
			}
			if b.rowCells[j][i] == -1 {
				colMatch++
			}
		}
		if colMatch == len(b.rowCells) || rowMatch == len(b.rowCells) {
			b.HasBingo = true
			return true
		}
	}
	return false
}

func getBingoInput(fileName string, bingo *Bingo) ([]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var moves []int

	scanner := bufio.NewScanner(file)
	line := 0
	board := -1
	row := 0
	for scanner.Scan() {
		if line == 0 {
			movesString := strings.Split(scanner.Text(), ",")
			for i := range movesString {
				intMove, _ := strconv.Atoi(movesString[i])
				moves = append(moves, intMove)
			}
			line++
			continue
		}
		if scanner.Text() == "" {
			board++
			row = 0
		} else {
			numStrings := strings.Fields(scanner.Text())
			if len(bingo.Boards) <= board {
				bingo.Boards = append(bingo.Boards, new(Board))
			}
			for i := range numStrings {
				bingo.Boards[board].rowCells[row][i], _ = strconv.Atoi(numStrings[i])
			}
			row++
		}
	}
	return moves, scanner.Err()
}
