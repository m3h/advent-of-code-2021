package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func read_input(path string) (xs []int, ys []int, fold_types []int, fold_points []int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	xs = make([]int, 0)
	ys = make([]int, 0)

	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "" {
			break
		}
		p := strings.Split(scanner.Text(), ",")

		if len(p) != 2 {
			log.Fatal("Too many points. Found:", p, "in line:", scanner.Text())
		}

		xi, err := strconv.Atoi(p[0])
		if err != nil {
			log.Fatal("err in line", scanner.Text())
		}
		yi, err := strconv.Atoi(p[1])
		if err != nil {
			log.Fatal("err in line", scanner.Text())
		}

		xs = append(xs, xi)
		ys = append(ys, yi)
	}

	fold_types = make([]int, 0)
	fold_points = make([]int, 0)

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) == 0 {
			// blank line, skip
			continue
		} else if len(fields) != 3 {
			log.Fatal("Too many fields in line", scanner.Text())
		}

		fold_info := strings.Split(fields[2], "=")

		fold_type_str, fold_point_str := fold_info[0], fold_info[1]

		fold_point, err := strconv.Atoi(fold_point_str)
		if err != nil {
			log.Fatal(err)
		}

		var fold_type int
		if fold_type_str == "x" {
			fold_type = 0
		} else if fold_type_str == "y" {
			fold_type = 1
		} else {
			log.Fatal("Unrecognized fold type", fold_type_str)
		}

		fold_types = append(fold_types, fold_type)
		fold_points = append(fold_points, fold_point)
	}

	return xs, ys, fold_types, fold_points
}

func max(arr []int) (m int, mi int) {

	m = arr[0]
	mi = 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > m {
			m = arr[i]
			mi = i
		}
	}
	return m, mi
}

func print_page(xs []int, ys []int) {
	xmax, _ := max(xs)
	ymax, _ := max(ys)

	render := make([][]string, ymax+1)
	for i := 0; i < len(render); i++ {
		render[i] = make([]string, xmax+1)

		for j := 0; j < len(render[i]); j++ {
			render[i][j] = "."
		}
	}
	for i := 0; i < len(xs); i++ {
		y := ys[i]
		x := xs[i]
		render[y][x] = "#"
	}

	fmt.Println()
	for i := 0; i < len(render); i++ {
		for j := 0; j < len(render[i]); j++ {
			fmt.Print(render[i][j])
		}
		fmt.Println()
	}
	fmt.Println()

}

func index(arr []int, v int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == v {
			return i
		}
	}
	return -1
}

func contains(xs []int, ys []int, x int, y int) bool {
	for i := 0; i < len(xs); i++ {
		if xs[i] == x && ys[i] == y {
			return true
		}
	}
	return false
}

// func contains(arr []int, v int) bool {
// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] == v {
// 			return true
// 		}
// 	}
// 	return false
// }

func fold(xs []int, ys []int, ftype int, fpoint int) ([]int, []int) {

	xs_new := make([]int, 0)
	ys_new := make([]int, 0)

	for i := 0; i < len(xs); i++ {
		var x_new, y_new int
		if ftype == 0 {
			if xs[i] > fpoint {
				x_new = fpoint - (xs[i] - fpoint)
			} else {
				x_new = xs[i]
			}
			y_new = ys[i]

			if contains(xs_new, ys_new, x_new, y_new) {
				continue
			} else {
				xs_new = append(xs_new, x_new)
				ys_new = append(ys_new, y_new)
			}

		} else if ftype == 1 {
			if ys[i] > fpoint {
				y_new = fpoint - (ys[i] - fpoint)
			} else {
				y_new = ys[i]
			}
			x_new = xs[i]

			if contains(xs_new, ys_new, x_new, y_new) {
				continue
			} else {
				xs_new = append(xs_new, x_new)
				ys_new = append(ys_new, y_new)
			}

		}
	}

	return xs_new, ys_new
}

func main() {

	input_path := "day13_input_test.txt"
	input_path = "day13_input.txt"
	if len(os.Args) > 1 {
		input_path = os.Args[1]
	}

	xs, ys, fold_types, fold_points := read_input(input_path)
	if len(xs) != len(ys) || len(fold_types) != len(fold_points) {
		log.Fatal("mismatch in input arrays")
	}

	fmt.Println("input points")
	for i := 0; i < len(xs); i++ {
		fmt.Println(xs[i], ys[i])
	}

	fmt.Println("input folds")
	for i := 0; i < len(fold_types); i++ {
		fmt.Println(fold_types[i], fold_points[i])
	}

	// print_page(xs, ys)
	fmt.Println("perform folds")
	for i := 0; i < len(fold_types); i++ {

		xs, ys = fold(xs, ys, fold_types[i], fold_points[i])
		fmt.Println("fold", i, "type:", fold_types[i], "point:", fold_points[i], "remaining points:", len(xs))
		// print_page(xs, ys)
	}

	print_page(xs, ys)
}
