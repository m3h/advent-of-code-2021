package main

import (
    "fmt"
    "os"
    "bufio"
    "log"
    "strconv"
    )

func main() {
    //file, err := os.Open("input_test.txt")
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    scanner.Scan()
    n1, err := strconv.Atoi(scanner.Text())
    if err != nil {
        log.Fatal(err)
    }
    scanner.Scan()
    n2, err := strconv.Atoi(scanner.Text())
    if err != nil {
        log.Fatal(err)
    }
    scanner.Scan()
    n3, err := strconv.Atoi(scanner.Text())
    if err != nil {
        log.Fatal(err)
    }

    increases := 0
    for scanner.Scan() {
        n4, err := strconv.Atoi(scanner.Text())
        if err != nil {
            log.Fatal(err)
        }

        s1 := n1 + n2 + n3
        s2 := n2 + n3 + n4

        fmt.Println(n1, n2, n3, n4, s1, s2)
        if s2 > s1 {
            increases += 1
        }

        n1 = n2
        n2 = n3
        n3 = n4

    }

    fmt.Println("Total increases:", increases)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}


