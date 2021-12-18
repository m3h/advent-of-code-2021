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

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func remove(slice [][][]int, s int) [][][]int {
	return append(slice[:s], slice[s+1:]...)
}

func printField(field [][]int) {
	fmt.Println()
	for y := 0; y < len(field); y++ {
		for x := 0; x < len(field[y]); x++ {
			if field[y][x] == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(field[y][x])
			}
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

	x1 := make([]int, 0)
	y1 := make([]int, 0)
	x2 := make([]int, 0)
	y2 := make([]int, 0)

	x_max := 0
	y_max := 0

	for {
		input, _ := reader.ReadString('\n')
		if strings.TrimSpace(input) == "" {
			break
		}
		spl_p1 := strings.Split(input, "->")
		spl := make([]int, 0)

		for i := 0; i < len(spl_p1); i++ {
			spl_p2 := strings.Split(strings.TrimSpace(spl_p1[i]), ",")

			for j := 0; j < len(spl_p2); j++ {
				n, _ := strconv.Atoi(spl_p2[j])
				spl = append(spl, n)
			}
		}

		x1 = append(x1, spl[0])
		y1 = append(y1, spl[1])
		x2 = append(x2, spl[2])
		y2 = append(y2, spl[3])

	}

	for i := 0; i < len(x1); i++ {
		fmt.Printf("%d;%d --> %d;%d\n", x1[i], y1[i], x2[i], y2[i])

		x_max = max(x_max, x1[i])
		x_max = max(x_max, x2[i])

		y_max = max(y_max, y1[i])
		y_max = max(y_max, y2[i])
	}

	field := make([][]int, y_max+1)
	for i := 0; i < len(field); i++ {
		field[i] = make([]int, x_max+1)

		for j := 0; j < len(field[i]); j++ {
			field[i][j] = 0
		}
	}

	printField(field)

	for i := 0; i < len(x1); i++ {
		if x1[i] == x2[i] {
			ys := min(y1[i], y2[i])
			ye := max(y1[i], y2[i])
			for y := ys; y <= ye; y++ {
				field[y][x1[i]] += 1
			}
		}

		if y1[i] == y2[i] {
			xs := min(x1[i], x2[i])
			xe := max(x1[i], x2[i])
			for x := xs; x <= xe; x++ {
				field[y1[i]][x] += 1
			}
		}
	}

	printField(field)

	c := 0
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			if field[i][j] >= 2 {
				c += 1
			}
		}
	}
	fmt.Println("Answer", c)
}
