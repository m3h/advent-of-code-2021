package main

import "fmt"

func update_step(state []int) []int {
	zc := state[0]

	for i := 0; i < len(state)-1; i++ {
		state[i] = state[i+1]
	}
	state[8] = zc
	state[6] += zc

	return state
}

func main() {
	initial_state := []int{3, 4, 3, 1, 2}
	initial_state = []int{1, 1, 1, 1, 1, 5, 1, 1, 1, 5, 1, 1, 3, 1, 5, 1, 4, 1, 5, 1, 2, 5, 1, 1, 1, 1, 3, 1, 4, 5, 1, 1, 2, 1, 1, 1, 2, 4, 3, 2, 1, 1, 2, 1, 5, 4, 4, 1, 4, 1, 1, 1, 4, 1, 3, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 5, 4, 4, 2, 4, 5, 2, 1, 5, 3, 1, 3, 3, 1, 1, 5, 4, 1, 1, 3, 5, 1, 1, 1, 4, 4, 2, 4, 1, 1, 4, 1, 1, 2, 1, 1, 1, 2, 1, 5, 2, 5, 1, 1, 1, 4, 1, 2, 1, 1, 1, 2, 2, 1, 3, 1, 4, 4, 1, 1, 3, 1, 4, 1, 1, 1, 2, 5, 5, 1, 4, 1, 4, 4, 1, 4, 1, 2, 4, 1, 1, 4, 1, 3, 4, 4, 1, 1, 5, 3, 1, 1, 5, 1, 3, 4, 2, 1, 3, 1, 3, 1, 1, 1, 1, 1, 1, 1, 1, 1, 4, 5, 1, 1, 1, 1, 3, 1, 1, 5, 1, 1, 4, 1, 1, 3, 1, 1, 5, 2, 1, 4, 4, 1, 4, 1, 2, 1, 1, 1, 1, 2, 1, 4, 1, 1, 2, 5, 1, 4, 4, 1, 1, 1, 4, 1, 1, 1, 5, 3, 1, 4, 1, 4, 1, 1, 3, 5, 3, 5, 5, 5, 1, 5, 1, 1, 1, 1, 1, 1, 1, 1, 2, 3, 3, 3, 3, 4, 2, 1, 1, 4, 5, 3, 1, 1, 5, 5, 1, 1, 2, 1, 4, 1, 3, 5, 1, 1, 1, 5, 2, 2, 1, 4, 2, 1, 1, 4, 1, 3, 1, 1, 1, 3, 1, 5, 1, 5, 1, 1, 4, 1, 2, 1}
	sim_days := 256

	state := make([]int, 9)
	for i := 0; i < len(state); i++ {
		state[i] = 0
	}
	for i := 0; i < len(initial_state); i++ {
		state[initial_state[i]]++
	}

	fmt.Println("initial state:", state)

	for i := 0; i < sim_days; i++ {
		state = update_step(state)

		fmt.Printf("After %3d days:", i+1)
		fmt.Println(state)
	}

	s := 0
	for i := 0; i < len(state); i++ {
		s += state[i]
	}
	fmt.Println("Num fishies:", s)

}
