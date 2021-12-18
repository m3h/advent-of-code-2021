package main

import (
	"fmt"
	"log"
	"os"
)

func bit_array_to_int(bit_array []int) int {
    v := 0
    for i := 0; i < len(bit_array); i++ {
        v += bit_array[i] * 1<<(len(bit_array)-1-i)
    }
    return v
}

func find_most_common(byte_array [][]int, idx int) int {
    count_1 := 0
    count_0 := 0
    for i := 0; i < len(byte_array); i++ {
        if byte_array[i][idx] == 1 {
            count_1++
        } else if byte_array[i][idx] == 0 {
            count_0++
        } else {
            fmt.Println("error in find_most_common")
        }
    }

    if count_1 == count_0 {
        return -1
    } else if count_1 > count_0 {
        return 1
    } else {
        return 0
    }
}


func rating(byte_array [][]int, idx int, most_common bool) []int {
    if len(byte_array) == 1 {
        return byte_array[0]
    } else if len(byte_array) == 0 {
        fmt.Println("error in rating")
    }

    byte_array_filt := make([][]int, 0)
    most_common_bit := find_most_common(byte_array, idx)

    if most_common_bit == -1 {
        // tie
        most_common_bit = 1
        // if most_common {
        //     most_common_bit = 1
        // } else {
        //     most_common_bit = 0
        // }

        fmt.Println("tie! choose", most_common_bit)
    }

    for i := 0; i < len(byte_array); i++ {
        v := byte_array[i][idx]
        if (most_common && (v == most_common_bit)) || (!most_common && (v != most_common_bit)) {
            byte_array_filt = append(byte_array_filt, byte_array[i])
        }
    }

    fmt.Println("filtered: idx", idx, "most_common", most_common, "arr", byte_array_filt)

    return rating(byte_array_filt, idx+1, most_common)
}

func main() {

	// file, err := os.Open("input_test.txt")
	file, err := os.Open("input.txt");
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

    fmt.Println("input")
	for i := 0; i < len(bits_array); i++ {
		fmt.Println(i, bits_array[i])
	}

    fmt.Println("determine for oxy")
    oxygen_rating_arr := rating(bits_array, 0, true)
    fmt.Println("determine for scrubby")
    scrubber_rating_arr := rating(bits_array, 0, false)

    oxygen_rating := bit_array_to_int(oxygen_rating_arr)
    scrubber_rating := bit_array_to_int(scrubber_rating_arr)

    life_support_rating := oxygen_rating * scrubber_rating

    fmt.Println("oxygen_rating", oxygen_rating_arr, oxygen_rating)
    fmt.Println("scrubber_rating", scrubber_rating_arr, scrubber_rating)
    fmt.Println("life support rating", life_support_rating)

}
