package main

import (
    "fmt"
    "os"
    "log"
    )

func main() {
    // file, err := os.Open("input_test.txt")
    file, err := os.Open("input.txt");
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    byte_array := make([]byte, 1)
    sum_array := make([]int, 0)
    for {
        file.Read(byte_array)

        if byte_array[0] == '\n' {
            break
        } else {
            sum_array = append(sum_array, 0)
        }
    }

    count := 0
    i := 0
    for {
        file.Read(byte_array)
        if byte_array[0] == '\n' {
            if i == 0 {
                break
            }
            i = 0
            count += 1
        } else {
            if byte_array[0] == '1' {
                sum_array[i] += 1
            }
            i += 1
        }
    }

    gamma := 0
    epsilon := 0

    for j := 0; j < len(sum_array); j++ {
        if sum_array[j] > count - sum_array[j] {
            sum_array[j] = 1

            gamma += 1<<(len(sum_array)-1-j)
        } else {
            sum_array[j] = 0

            epsilon += 1<<(len(sum_array)-1-j)
        }
    }

    fmt.Println(sum_array)
    fmt.Println("gamma", gamma)
    fmt.Println("epsilon", epsilon)
    fmt.Println("product", gamma * epsilon)

}


