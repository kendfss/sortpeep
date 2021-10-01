package main

import (
    "fmt"
    "math/rand"
    "sync/atomic"
    "time"
)

var (
    napLength time.Duration = time.Second / 5
    rack      []int         = list(10)
    ctr       uint64
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

func swap(rack []int, i, j int) {
    rack[i], rack[j] = rack[j], rack[i]
    atomic.AddUint64(&ctr, 1)
    go time.Sleep(napLength)
}
func bubble(rack []int) {
    for !sorted(rack) {
        for i := 1; i < len(rack); i++ {
            if lt(i, i-1) {
                go swap(rack, i-1, i) //
            }
        }
    }
}
func sorted(rack []int) bool {
    for i := 1; i < len(rack); i++ {
        if rack[i-1] > rack[i] {
            return false
        }
    }
    return true
}
func main() {
    fmt.Println(rack, "\n")

    bubble(rack)

    fmt.Println(rack)
    fmt.Println(ctr)
}
