package main

import (
	"fmt"
	"log"
	"os"
)

func read_input(path string) [][]int {
	inp := make([][]int, 1)

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
			inp = append(inp, make([]int, 0))
		} else {

			inp[len(inp)-1] = append(inp[len(inp)-1], int(b[0])-int('0'))
		}

	}

	if len(inp[len(inp)-1]) == 0 {
		inp = inp[:len(inp)-1]
	}
	return inp
	// return inp[:len(inp)-1]

}

func increment(oct [][]int, i, j int) ([][]int, int) {
	flashes := 0

	if i < 0 || j < 0 || i >= len(oct) || j >= len(oct[i]) {
		return oct, flashes
	}

	oct[i][j]++

	if oct[i][j] == 10 {
		// flash flash
		flashes++
		// brighten up neighbours
		for x := -1; x <= 1; x++ {
			for y := -1; y <= 1; y++ {
				if x == 0 && y == 0 {
					continue
				}
				var nf int
				oct, nf = increment(oct, i+x, j+y)
				flashes += nf
			}
		}

	}

	return oct, flashes
}

func step(oct [][]int) ([][]int, int) {

	flashes := 0
	for i := 0; i < len(oct); i++ {
		for j := 0; j < len(oct[i]); j++ {
			var nf int
			oct, nf = increment(oct, i, j)
			flashes += nf
		}
	}

	// reset any flashes occies to 0
	// doing it now lets potential re-flashers only reset after we flash
	// e.g. if an octo hits 10, flash. Higher than 10, don't flash again
	// but reset now
	for i := 0; i < len(oct); i++ {
		for j := 0; j < len(oct[i]); j++ {
			if oct[i][j] >= 10 {
				oct[i][j] = 0
			}
		}
	}

	return oct, flashes
}

func print_occies(oct [][]int) {
	for i := 0; i < len(oct); i++ {
		fmt.Println(oct[i])
	}
	fmt.Println()
}

func main() {
	var path string

	// path = "day11p1_input_small.txt"
	// path = "day11p1_input_test.txt"
	path = "day11p1_input.txt"

	oct := read_input(path)

	fmt.Println("input:", len(oct))
	print_occies(oct)

	nf := 0
	flashes := 0

	for s := 1; ; s++ {
		oct, nf = step(oct)
		flashes += nf

		fmt.Println("step:", s)
		fmt.Println("oct:")
		print_occies(oct)
		fmt.Println("total flashes:", flashes)

		if nf == (len(oct)*len(oct[0])) {
			break
		}
	}
}
