package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Open("input_test.txt")
	// file, err := os.Open("input.txt");
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byte_array := make([]byte, 1)
	bits_array := make([][]int, 0)
	bits := make([]int, 0)
	for {
		file.Read(byte_array)

		if byte_array[0] == '\n' {
			if len(bits) == 0 {
				break
			} else {
				bits_array = append(bits_array, bits)
				bits = make([]int, 0)
			}
		} else {
			v := -1
			if byte_array[0] == '1' {
				v = 1
			} else if byte_array[0] == '0' {
				v = 0
			}

			bits = append(bits, v)
		}
	}

	for i := 0; i < len(bits_array); i++ {
		fmt.Println(bits_array[i])
	}

}
