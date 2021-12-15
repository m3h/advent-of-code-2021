package main

import (
    "fmt"
    "os"
    "bufio"
    "log"
    "strconv"
    )

func main() {
    // file, err := os.Open("input_test.txt")
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    scanner.Scan()
    p, err := strconv.Atoi(scanner.Text())
    if err != nil {
        log.Fatal(err)
    }

    increases := 0
    for scanner.Scan() {
        i, err := strconv.Atoi(scanner.Text())
        if err != nil {
            log.Fatal(err)
        }

        if i > p {
            fmt.Println(p, i, "increased")
            increases += 1
        } else {
            fmt.Println(p, i, "decreased")
        }
        p = i
    }

    fmt.Println("Total increases:", increases)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}


