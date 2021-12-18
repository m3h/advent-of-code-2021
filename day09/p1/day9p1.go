package main

import (
	"fmt"
	"log"
	"os"
	"sort"
)

func read_input(path string) [][]int {
	mp := make([][]int, 1)

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
			mp = append(mp, make([]int, 0))
		} else {
			row := mp[len(mp)-1]
			v := int(b[0] - '0')
			row = append(row, v)
			mp[len(mp)-1] = row
		}

	}
	return mp[:len(mp)-1]
}

func lowpoints(mp [][]int) ([]int, []int, []int) {
	pt := make([]int, 0)
	pt_i := make([]int, 0)
	pt_j := make([]int, 0)

	for i := 0; i < len(mp); i++ {
		for j := 0; j < len(mp[i]); j++ {

			lowest := true
			for x := -1; x <= 1; x++ {
				for y := -1; y <= 1; y++ {
					if x == 0 && y == 0 {
						continue
					} else if x != 0 && y != 0 {
						continue
					}

					ni := i + x
					nj := j + y

					if ni < 0 || nj < 0 || ni >= len(mp) || nj >= len(mp[ni]) {
						continue
					}

					if mp[i][j] >= mp[ni][nj] {
						lowest = false
						break
					}
				}
				if !lowest {
					break
				}
			}
			if lowest {
				pt = append(pt, mp[i][j])
				pt_i = append(pt_i, i)
				pt_j = append(pt_j, j)

			}
		}
	}
	return pt, pt_i, pt_j
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func basins_one(mp [][]int, i, j int) int {

	if mp[i][j] == 9 || mp[i][j] < 0 {
		return 0
	}

	s := 1
	// mp[i][j] = -abs(mp[i][j])
	mp[i][j] = -1

	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			} else if x != 0 && y != 0 {
				continue
			}

			ni := i + x
			nj := j + y

			if ni < 0 || nj < 0 || ni >= len(mp) || nj >= len(mp[ni]) {
				continue
			}

			s += basins_one(mp, ni, nj)
		}
	}

	return s
}

func basins(mp [][]int, lps_i []int, lps_j []int) []int {

	sizes := make([]int, len(lps_i))

	for i := 0; i < len(lps_i); i++ {
		sizes[i] = basins_one(mp, lps_i[i], lps_j[i])
	}

	return sizes
}

func main() {
	// mp := read_input("day9_input_test.txt")
	mp := read_input("day9_input.txt")

	fmt.Println("map", mp)

	low_points, lp_i, lp_j := lowpoints(mp)
	fmt.Println("low_points", low_points)

	sum := 0
	for i := 0; i < len(low_points); i++ {
		rl := low_points[i] + 1
		sum += rl
	}
	// fmt.Println("ANSWER", sum)

	sizes := basins(mp, lp_i, lp_j)

	sort.Ints(sizes)

	fmt.Println(sizes)

	prod := 1
	for i := len(sizes) - 3; i < len(sizes); i++ {
		prod *= sizes[i]
	}

	fmt.Println("PRODUCT", prod)

}
