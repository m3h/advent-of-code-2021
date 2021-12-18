package main

import (
	"fmt"
	"log"
	"os"
	"sort"
)

func read_input(path string) [][]rune {
	inp := make([][]rune, 1)

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b := make([]byte, 1)
	for {
		n, err := file.Read(b)
		if n == 0 {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if b[0] == '\n' {
			inp = append(inp, make([]rune, 0))
		} else {

			inp[len(inp)-1] = append(inp[len(inp)-1], rune(b[0]))
		}

	}
	return inp
	// return inp[:len(inp)-1]
}

func map_key(mp map[rune]rune, v rune) bool {
	for k := range mp {
		if k == v {
			return true
		}
	}
	return false
}

func map_val(mp map[rune]rune, v rune) bool {
	for k := range mp {
		if mp[k] == v {
			return true
		}
	}
	return false
}

func map_vtok(mp map[rune]rune, v rune) rune {
	for k := range mp {
		if mp[k] == v {
			return k
		}
	}
	return -1
}

func remove(slice []rune, i int) []rune {
	return append(slice[:i], slice[i+1:]...)
}

func pop(slice []rune) ([]rune, rune) {
	return slice[:len(slice)-1], slice[len(slice)-1]
}

func find_corrupted(inp [][]rune, pairs map[rune]rune) ([]int, []int, []rune, [][]rune) {
	corrupted_rows := make([]int, 0)
	corrupted_col := make([]int, 0)
	corrupted_type := make([]rune, 0)

	completions := make([][]rune, 0)

	// c := make(map[rune]int)
	// for k := range pairs {
	// 	c[k] = 0
	// }

	for i := 0; i < len(inp); i++ {
		fmt.Println()
		fmt.Println(inp[i])

		stack := make([]rune, 0)
		corrupt := false
		for j := 0; j < len(inp[i]); j++ {

			var k rune
			if map_key(pairs, inp[i][j]) {
				// open
				k = inp[i][j]

				stack = append(stack, k)
			} else if map_val(pairs, inp[i][j]) {
				k = map_vtok(pairs, inp[i][j])

				var stack_k rune
				stack, stack_k = pop(stack)

				if stack_k != k {
					corrupted_rows = append(corrupted_rows, i)
					corrupted_col = append(corrupted_col, j)
					corrupted_type = append(corrupted_type, k)

					corrupt = true
					break
				}
			}
		}

		if corrupt {
			fmt.Println("corrupt!")
		} else if len(stack) > 0 {
			fmt.Println("incomplete")

			completions = append(completions, make([]rune, 0))
			for len(stack) > 0 {
				var stack_k rune
				stack, stack_k = pop(stack)

				idx := len(completions) - 1
				v := pairs[stack_k]
				completions[idx] = append(completions[idx], v)
			}
		}

	}

	return corrupted_rows, corrupted_col, corrupted_type, completions
}

func main() {
	// inp := read_input("day10_input_test.txt")
	inp := read_input("day10_input.txt")

	pairs := make(map[rune]rune)
	pairs['('] = ')'
	pairs['['] = ']'
	pairs['{'] = '}'
	pairs['<'] = '>'

	costs := make(map[rune]int)
	costs['('] = 3
	costs['['] = 57
	costs['{'] = 1197
	costs['<'] = 25137

	fmt.Println("inp", inp)

	_, _, ct, completions := find_corrupted(inp, pairs)

	score := 0
	for i := 0; i < len(ct); i++ {
		k := ct[i]
		if map_val(pairs, k) {
			k = map_vtok(pairs, k)
		}

		sc := costs[k]
		score += sc
	}
	fmt.Println("corrupt score", score)

	icosts := make(map[rune]int)
	icosts[')'] = 1
	icosts[']'] = 2
	icosts['}'] = 3
	icosts['>'] = 4
	i_score := make([]int, len(completions))

	for i := 0; i < len(completions); i++ {
		i_score_t := 0
		for j := 0; j < len(completions[i]); j++ {
			is := icosts[completions[i][j]]
			i_score_t = i_score_t*5 + is
		}
		i_score[i] = i_score_t
		fmt.Println("completion:", completions[i], "score", i_score_t)
	}

	sort.Ints(i_score)

	fmt.Println("scores:", i_score)

	middle := i_score[len(i_score)/2]
	fmt.Println("SCORE:", middle)
}
