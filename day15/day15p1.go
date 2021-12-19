package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func str_to_arr(str string) []int {
	arr := make([]int, len(str))

	for i := 0; i < len(str); i++ {
		n, ok := strconv.Atoi(str[i : i+1])
		if !ok {
			log.Fatal(str)
		}
		arr[i] = n
	}

	return arr
}

func read_input(path string) [][]int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	mp := make([][]int, 0)
	for scanner.Scan() {

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		line_arr := str_to_arr(line)
		mp = append(mp, line_arr)

	}
	return mp
}

func extend(a []string, b []string) []string {
	r := make([]string, len(a)+len(b))

	for i := 0; i < len(a); i++ {
		r[i] = a[i]
	}
	j := len(a)
	for i := 0; i < len(b); i++ {
		r[j] = b[i]
		j++
	}
	return r
}

func main() {

	input_path := "day15_input_test.txt"
	input_path = "day15_input.txt"
	if len(os.Args) > 1 {
		input_path = os.Args[1]
	}

}
