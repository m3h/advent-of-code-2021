package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func read_input(path string) (signal_pattern [][][]string, output [][][]string) {

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	// side, row, word, char
	data := make([][][][]string, 2)
	for i := 0; i < len(data); i++ {
		data[i] = make([][][]string, 0)
	}

	// loop through rows
	for {
		input, _ := reader.ReadString('\n')
		if input == "" {
			break
		}

		str_arr := strings.Split(strings.TrimSpace(input), "|")

		fmt.Println("row", str_arr)
		// loop through sides
		for s := 0; s < len(str_arr); s++ {
			fmt.Println("side", str_arr[s])

			data[s] = append(data[s], make([][]string, 0))
			r := len(data[s]) - 1
			data[s][r] = make([][]string, 0)

			words := strings.Fields(str_arr[s])
			for j := 0; j < len(words); j++ {
				word := []rune(words[j])
				fmt.Println("word", word)

				data[s][r] = append(data[s][r], make([]string, 0))
				w := len(data[s][r]) - 1

				for k := 0; k < len(word); k++ {
					seg := string(word[k])

					fmt.Println("seg", seg)

					data[s][r][w] = append(data[s][r][w], seg)

				}
			}
		}

	}

	return data[0], data[1]
}

func intersection(a []string, b []string) []string {

	r := make([]string, 0)

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				r = append(r, a[i])
			}
		}
	}

	return r
}

func get_segs(signal_pattern [][]string, n int) []string {
	for w := 0; w < len(signal_pattern); w++ {
		if len(signal_pattern[w]) == n {
			return signal_pattern[w]
		}
	}

	return make([]string, 0)
}

func get_count(words [][]string, n int) []string {

	matches := make([]string, 0)
	segs := []string{"a", "b", "c", "d", "e", "f", "g"}
	for s := 0; s < len(segs); s++ {

		count := 0
		for w := 0; w < len(words); w++ {
			for c := 0; c < len(words[w]); c++ {
				if words[w][c] == segs[s] {
					count += 1
				}
			}
		}

		if count == n {
			matches = append(matches, segs[s])
		}
	}

	return matches
}

func not_in(a []string, b []string) []string {
	ret := make([]string, 0)

	for i := 0; i < len(a); i++ {
		found := false
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				found = true
				break
			}
		}
		if !found {
			ret = append(ret, a[i])
		}
	}

	return ret
}

func arr_eq(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		found := false
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func wtoi(w []string) int {
	switch {
	case arr_eq(w, []string{"a", "b", "c", "e", "f", "g"}):
		return 0
	case arr_eq(w, []string{"c", "f"}):
		return 1
	case arr_eq(w, []string{"a", "c", "d", "e", "g"}):
		return 2
	case arr_eq(w, []string{"a", "c", "d", "f", "g"}):
		return 3
	case arr_eq(w, []string{"b", "c", "d", "f"}):
		return 4
	case arr_eq(w, []string{"a", "b", "d", "f", "g"}):
		return 5
	case arr_eq(w, []string{"a", "b", "d", "e", "f", "g"}):
		return 6
	case arr_eq(w, []string{"a", "c", "f"}):
		return 7
	case arr_eq(w, []string{"a", "b", "c", "d", "e", "f", "g"}):
		return 8
	case arr_eq(w, []string{"a", "b", "c", "d", "f", "g"}):
		return 9
	}
	return -1
}

func substitute(words [][]string, a, b, c, d, e, f, g string) [][]string {
	ret := make([][]string, len(words))

	for w := 0; w < len(words); w++ {

		ret[w] = make([]string, len(words[w]))
		for ci := 0; ci < len(words[w]); ci++ {
			v := ""

			switch words[w][ci] {
			case a:
				v = "a"
			case b:
				v = "b"
			case c:
				v = "c"
			case d:
				v = "d"
			case e:
				v = "e"
			case f:
				v = "f"
			case g:
				v = "g"
			}

			ret[w][ci] = v
		}
	}

	return ret
}

func main() {

	// sp, op := read_input("day8p1_input_test_small.txt")
	// sp, op := read_input("day8p1_input_test.txt")
	sp, op := read_input("day8p1_input.txt")

	fmt.Println()

	sum := 0
	for r := 0; r < len(sp); r++ {
		words := sp[r]

		// seven := get_segs(words, 3)

		// a := intersection(one, seven)[0]

		// a, b, c, d, e, f,
		e := get_count(words, 4)[0]
		b := get_count(words, 6)[0]
		f := get_count(words, 9)[0]

		one := get_segs(words, 2)
		c := not_in(one, []string{f})[0]

		seven := get_segs(words, 3)
		a := not_in(seven, []string{c, f})[0]

		four := get_segs(words, 4)
		d := not_in(four, []string{b, c, f})[0]

		g := not_in([]string{"a", "b", "c", "d", "e", "f", "g"}, []string{a, b, c, d, e, f})[0]

		// fmt.Println(a, b, c, d, e, f, g)
		fmt.Println(a, b, c, d, e, f, g)

		// substitude
		sps := substitute(sp[r], a, b, c, d, e, f, g)
		ops := substitute(op[r], a, b, c, d, e, f, g)

		for w := 0; w < len(words); w++ {
			word := sps[w]
			wi := wtoi(word)

			fmt.Print(wi, ", ")
		}
		fmt.Print(" | ")

		answer := 0
		for w := 0; w < len(op[r]); w++ {
			word := ops[w]
			wi := wtoi(word)

			answer = answer*10 + wi

			fmt.Print(wi, ", ")
		}
		fmt.Print(" | ", answer, "\n")

		sum += answer
	}

	fmt.Println("ANSWER:", sum)

	// fmt.Println(count)
}
