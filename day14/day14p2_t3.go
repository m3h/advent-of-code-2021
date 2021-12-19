package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func str_to_arr(str string) []string {
	arr := make([]string, len(str))

	for i := 0; i < len(str); i++ {
		arr[i] = str[i : i+1]
	}

	return arr
}

func read_input(path string) (template []string, rules map[string]map[string]string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	template = str_to_arr(scanner.Text())

	rules = make(map[string]map[string]string)

	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "" {
			continue
		}

		fields := strings.Split(scanner.Text(), " -> ")
		pattern := str_to_arr(strings.TrimSpace(fields[0]))
		insert := strings.TrimSpace(fields[1])

		l, r := pattern[0], pattern[1]

		if _, ok := rules[l]; !ok {
			rules[l] = make(map[string]string)
		}

		rules[l][r] = insert
	}

	return template, rules
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

func add_counts(count map[string]int, add map[string]int) {
	for k, v := range add {
		if _, ok := count[k]; !ok {
			count[k] = 0
		}
		count[k] += v
	}
}

func update_cache(cache map[string]map[int]map[string]int, lmr string, depth int, count map[string]int) {
	if _, ok := cache[lmr]; !ok {
		cache[lmr] = make(map[int]map[string]int)
	}
	if _, ok := cache[lmr][depth]; ok {
		log.Fatal("cache double create")
	}

	cache[lmr][depth] = count
}

func apply_rules_help(l, m, r string, rules map[string]map[string]string, depth int, limit int, cache map[string]map[int]map[string]int) map[string]int {

	// check if result is cached
	lmr := l + m + r
	if v, ok := cache[lmr][depth]; ok {
		return v
	}

	// count new element
	count := make(map[string]int)
	count[m] = 1

	if depth >= limit {
		update_cache(cache, lmr, depth, count)
		return count
	}

	// l rule
	if v, ok := rules[l][m]; ok {
		count_l := apply_rules_help(l, v, m, rules, depth+1, limit, cache)
		add_counts(count, count_l)
	}

	// r rule
	if v, ok := rules[m][r]; ok {
		count_r := apply_rules_help(m, v, r, rules, depth+1, limit, cache)
		add_counts(count, count_r)
	}

	update_cache(cache, lmr, depth, count)
	return count
}

func apply_rules(template []string, rules map[string]map[string]string, limit int) map[string]int {

	// create cache
	cache := make(map[string]map[int]map[string]int)
	// count new elements
	count := count_elements(template)

	for i := 0; i < len(template)-1; i++ {

		l, r := template[i], template[i+1]

		if _, okl := rules[l]; okl {
			if v, okr := rules[l][r]; okr {

				c := apply_rules_help(l, v, r, rules, 1, limit, cache)
				add_counts(count, c)
			}
		}
	}
	return count
}

func count_elements(template []string) map[string]int {
	count := make(map[string]int)
	// count
	for i := 0; i < len(template); i++ {
		v := template[i]
		if _, ok := count[v]; !ok {
			count[v] = 0
		}
		count[template[i]]++
	}
	return count
}

func main() {

	input_path := "day14_input_test.txt"
	input_path = "day14_input.txt"
	if len(os.Args) > 1 {
		input_path = os.Args[1]
	}

	template, rules := read_input(input_path)

	fmt.Println("template", template)
	fmt.Println("rules:", rules)

	limit := 40
	count := apply_rules(template, rules, limit)

	// add untouched to template
	// count elements
	fmt.Println("count", count)

	max_count, max_val := math.MinInt, ""
	min_count, min_val := math.MaxInt, ""

	for val, val_count := range count {
		fmt.Println("Key:", val, "=>", "Element:", val_count)

		if val_count > max_count {
			max_count = val_count
			max_val = val
		}
		if val_count < min_count {
			min_count = val_count
			min_val = val
		}
	}

	fmt.Println("Most common", max_val, max_count)
	fmt.Println("Least common", min_val, min_count)

	fmt.Println("ANSWER", max_count-min_count)

}
