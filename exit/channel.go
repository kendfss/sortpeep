package main

import (
    "fmt"
    "math/rand"
    "time"
)

var (
    napLength time.Duration = time.Second / 5
    killer    chan bool     = make(chan bool)
    rack      []int         = list(10)
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
    for {
        select {
        case v := <-killer:
            return
        default:
            put(rack)
        }
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
    for !sorted(rack) {
        for i := 1; i < len(rack); i++ {
            if lt(i, i-1) {
                swap(rack, i-1, i)
            }
        }
    }
    killer <- true
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

    fmt.Println(rack, "\n")

    go viewer(rack)
    bubble(rack)

    fmt.Println(rack)
}
