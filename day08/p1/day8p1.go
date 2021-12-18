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

func main() {
	// _, output := read_input("day8p1_input_test.txt")
	_, output := read_input("day8p1_input.txt")

	count := 0
	for r := 0; r < len(output); r++ {
		for w := 0; w < len(output[r]); w++ {
			word := output[r][w]
			lw := len(word)
			if lw == 2 || lw == 4 || lw == 3 || lw == 7 {
				count += 1
			}
		}
	}
	fmt.Println(count)
}
