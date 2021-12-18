package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// add a->b
func add_edge(edges map[string][]string, a string, b string) map[string][]string {

	if _, ok := edges[a]; !ok {
		edges[a] = make([]string, 0)
	}
	edges[a] = append(edges[a], b)
	return edges
}

func read_input(path string) (edges map[string][]string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	edges = make(map[string][]string, 0)
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "" {
			continue
		}
		e := strings.Split(scanner.Text(), "-")

		if len(e) != 2 {
			log.Fatal("Too many edges. Found:", e, "in line:", scanner.Text())
		}

		edges = add_edge(edges, e[0], e[1])
		edges = add_edge(edges, e[1], e[0])
	}

	return edges
}

func isbig(a string) bool {
	return (strings.ToUpper(a) == a)
}

func copy(s []string) []string {
	r := make([]string, len(s))
	for i := 0; i < len(s); i++ {
		r[i] = s[i]
	}
	return r
}

func contains(arr []string, x string) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == x {
			return true
		}
	}
	return false
}

func bfs(edges map[string][]string, start string, end string, current_path []string, twice bool) [][]string {

	path := append(current_path, start)

	paths := make([][]string, 0)
	if start == end {
		return append(paths, path)
	}

	ns := edges[start]
	for i := 0; i < len(ns); i++ {
		n := ns[i]

		twice_now := twice
		// don't go through small caves twice
		if !isbig(n) && contains(path, n) {
			if n == "start" || n == "end" || twice {
				continue
			} else {
				twice_now = true
			}

		}
		new_paths := bfs(edges, n, end, path, twice_now)

		for j := 0; j < len(new_paths); j++ {
			paths = append(paths, new_paths[j])
		}
	}

	return paths
}

func main() {

	input_path := "maps/small1.txt"
	input_path = "maps/input.txt"
	if len(os.Args) > 1 {
		input_path = os.Args[1]
	}

	start_node := "start"
	end_node := "end"

	e := read_input(input_path)
	fmt.Println("edges", e)

	paths := bfs(e, start_node, end_node, make([]string, 0), false)

	for i := 0; i < len(paths); i++ {
		fmt.Println(paths[i])
	}

	fmt.Println("number of paths:", len(paths))
}
