package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

var (
    wg        sync.WaitGroup = sync.WaitGroup{}
    napLength time.Duration  = time.Second / 5
    rack      []int          = list(10)
    done      bool           = false
)

func lt(a, b int) bool {
    return rack[a] < rack[b]
}

func list(n int) []int {
    rack := []int{}
    for i := 0; i < n; i++ {
        rack = append(rack, rand.Intn(n))
    }
    return rack
}

func viewer(rack []int) {
    // wg.Add(1) // unfortunately this will not catch the wait group before the program exits
    defer wg.Done()
    for !done {
        put(rack)
    }
}
func put(arg []int) {
    fmt.Printf("%v\r", arg)
}
func swap(rack []int, i, j int) {
    rack[i], rack[j] = rack[j], rack[i]
    time.Sleep(napLength)

}
func bubble(rack []int) {
    // wg.Add(1) // unfortunately this will not catch the wait group before the program exits
    defer wg.Done()
    for !sorted(rack) {
        for i := 1; i < len(rack); i++ {
            if lt(i, i-1) {
                swap(rack, i-1, i)
            }
        }
    }
    done = true
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
    // rand.Seed(time.Now().UnixNano())

    wg.Add(2)

    fmt.Println(rack, "\n")
    go viewer(rack)
    go bubble(rack)
    wg.Wait()
    fmt.Println(rack)
}
