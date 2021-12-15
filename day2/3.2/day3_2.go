package main

import (
    "fmt"
    "os"
    "bufio"
    "log"
    "strconv"
    "strings"
    )

func main() {
    //file, err := os.Open("input_test.txt")
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    
    horizontal_position := 0
    vertical_position := 0
    aim := 0

    fmt.Println("direction, magnitude, aim, horizontal_position, vertical_position")
    for scanner.Scan() {

        words := strings.Fields(scanner.Text())
        direction := words[0]
        magnitude, err := strconv.Atoi(words[1])
        if err != nil {
            log.Fatal(err)
        }

        if direction == "up" {
            direction = "down"
            magnitude *= -1
        }

        if direction == "down" {
            aim += magnitude
        } else if direction == "forward" {
            horizontal_position += magnitude
            vertical_position += magnitude * aim
        }

        fmt.Println(direction, magnitude, aim, horizontal_position, vertical_position)
    }

    fmt.Println("horizontal:", horizontal_position)
    fmt.Println("vertical:", vertical_position)
    fmt.Println("product:", horizontal_position * vertical_position)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}


