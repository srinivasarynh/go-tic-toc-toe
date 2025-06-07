package main

import (
	"fmt"
	"strconv"
	"strings"
)


type Board [3][3]string

func main() {
	fmt.Println("welcome to tic tac toe...")
	fmt.Println("players alternate turns. enter position 1-9 to place your mark")
	fmt.Println("board postions:")
	showPositions()

	playAgain := true
	for playAgain {
		playGame()
		playAgain = askPlayAgain()
	}

	fmt.Println("thanks for playing")
}


func showPositions() {
	fmt.Println()
	fmt.Println(" 1 | 2 | 3 ")
	fmt.Println("-----------")
	fmt.Println(" 4 | 5 | 6 ")
	fmt.Println("-----------")
	fmt.Println(" 7 | 8 | 9 ")
	fmt.Println()
}

func playGame() {
	board := Board{}
	initializeBoard(&board)
	currentPlayer := "X"
	gameWon := false
	moves := 0

	fmt.Printf("\ngame starts... player %s goes first.\n", currentPlayer)

	for !gameWon && moves < 9 {
		displayBoard(board)

		postion := getPlayerMove(currentPlayer, board)
		makeMove(&board, postion, currentPlayer)
		moves++

		if checkWin(board, currentPlayer) {
			displayBoard(board)
			fmt.Printf("player %s wins...\n", currentPlayer)
			gameWon = true
		} else if moves == 9 {
			displayBoard(board)
			fmt.Println("it's a tie... board is full")
		} else {
			if currentPlayer == "X" {
				currentPlayer = "O"
			} else {
				currentPlayer = "X"
			}
		}
	}
}

func initializeBoard(board *Board) {
	for i := 0; i < 3; i++ {
		for j:= 0; j < 3; j++ {
			board[i][j] = " "
		}
	}
}


func displayBoard(board Board) {
	fmt.Println()
	for i := 0; i < 3; i++ {
		fmt.Printf(" %s | %s | %s \n", board[i][0], board[i][1], board[i][2])
		if i < 2 {
			fmt.Println("---|---|---")
		}
	}
	fmt.Println()
}

func getPlayerMove(player string, board Board) int {
	for {
		fmt.Printf("player %s, enter you move (1-9): ", player)
		var input string
		fmt.Scanf("%s", &input)

		position, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("please enter a valid number")
			continue
		}

		if position < 1 || position > 9 {
			fmt.Println("please enter a number between 1 and 9")
			continue
		}

		if !isValidMove(board, position) {
			fmt.Println("that position is already taken... choose another")
			continue
		}
		return position
	}
}


func isValidMove(board Board, position int) bool {
	row, col := positionToCoords(position)
	return board[row][col] == " "
}

func positionToCoords(position int) (int, int) {
	position--
	row := position/3
	col := position%3
	return row, col
} 

func makeMove(board *Board, position int, player string) {
	row, col := positionToCoords(position)
	board[row][col] = player
}

func checkWin(board Board, player string) bool {
	for i := 0; i < 3; i++ {
		if board[i][0] == player && board[i][1] == player && board[i][2] == player {
			return true
		}
	}

	for j := 0; j < 3; j++ {
		if board[0][j] == player && board[1][j] == player && board[2][j] == player {
			return true
		}
	}

	if board[0][0] == player && board[1][1] == player && board[2][2] == player {
		return true
	}

	if board[0][2] == player && board[1][1] == player && board[2][0] == player {
		return true
	}

	return false
}


func askPlayAgain() bool {
	fmt.Print("\nwould you like to play again? (y/n): ")
	var response string
	fmt.Scanf("%s", &response)

	response = strings.ToLower(response)
	return response == "y" || response == "yes"
}
