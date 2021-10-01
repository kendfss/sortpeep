package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

var (
	napLength time.Duration = time.Second / 5
	rack      []int         = list(10)
)

// put will print to standard out in a way that enables overwriting.
// Thus giving the effect of a screen update
func put(arg []int) {
	fmt.Printf("%v\r", rack)

}

// lt will be passed to the sort.Slice function to enable comparisons
func lt(i, j int) bool {
	// Calling "put" here enables us to view the slice each time its elements are compared
	put(rack)
	time.Sleep(napLength)
	return rack[i] < rack[j]
}

// list creates a slice of random integers
func list(n int) []int {
	rack := []int{}
	for i := 0; i < n; i++ {
		rack = append(rack, rand.Intn(n))
	}
	return rack
}

func main() {
	// rand.Seed(time.Now().UnixNano())

	fmt.Println(rack, "\n")

	sort.Slice(rack, lt)

	fmt.Println(rack)
}
