package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func arr_str_to_int(str_arr []string) []int {
	int_arr := make([]int, 0)
	for i := 0; i < len(str_arr); i++ {
		v, _ := strconv.Atoi(str_arr[i])
		int_arr = append(int_arr, v)
	}
	return int_arr
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func score(board [][]int, num int) int {
	sum := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			sum += max(board[i][j], 0)
		}
	}
	return sum * num
}

func check_winner(boards [][][]int, num int) []int {
	winner_idx := make([]int, 0)

	for i := 0; i < len(boards); i++ {
		for j := 0; j < len(boards[i]); j++ {
			for k := 0; k < len(boards[i][j]); k++ {
				if boards[i][j][k] == num {
					appended := false
					boards[i][j][k] = -1

					winner := true
					for m := 0; m < len(boards[i][j]); m++ {
						if boards[i][j][m] != -1 {
							winner = false
							break
						}
					}
					if winner {
						winner_idx = append(winner_idx, i)
						appended = true
					}

					winner = true
					for m := 0; m < len(boards[i]); m++ {
						if boards[i][m][k] != -1 {
							winner = false
							break
						}
					}
					if winner && !appended {
						winner_idx = append(winner_idx, i)
					}
				}
			}
		}
	}
	return winner_idx
}

func remove(slice [][][]int, s int) [][][]int {
	return append(slice[:s], slice[s+1:]...)
}

func print_boards(boards [][][]int) {
	fmt.Println()
	for i := 0; i < len(boards); i++ {
		for j := 0; j < len(boards[i]); j++ {
			for k := 0; k < len(boards[i][j]); k++ {
				fmt.Printf("%2d ", boards[i][j][k])
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func main() {

	// file, err := os.Open("input_test.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	input, _ := reader.ReadString('\n')

	chosen_str := strings.Split(input, ",")
	chosen := arr_str_to_int(chosen_str)

	boards := make([][][]int, 0)
	board := make([][]int, 0)
	for len(input) > 0 {
		input, _ = reader.ReadString('\n')

		if strings.TrimSpace(input) == "" {
			// new board
			if len(board) > 0 {
				boards = append(boards, board)
				board = make([][]int, 0)
			}
		} else {
			nums_str := strings.Fields(input)
			nums := arr_str_to_int(nums_str)
			board = append(board, nums)
		}
	}

	fmt.Println("chosen", chosen)
	print_boards(boards)

	for i := 0; i < len(chosen); i++ {
		fmt.Println("draw", chosen[i])
		winner_idx := check_winner(boards, chosen[i])
		print_boards(boards)

		if len(winner_idx) > 0 {

			if len(boards) > 1 {
				// removing winners
				for w := len(winner_idx) - 1; w >= 0; w-- {
					boards = remove(boards, winner_idx[w])
				}

			} else {

				// for w := 0; w < len(winner_idx); w++ {
				w := 0
				fmt.Println("winnder_idx", winner_idx)
				fmt.Println("winnder score", score(boards[winner_idx[w]], chosen[i]))
				break
				// }
			}
		}
	}

}
