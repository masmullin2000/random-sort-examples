package main

import (
    //"fmt"
    "math/rand"
    //"time"
    "sort"
)

func make_random_vec(sz uint64, mod int) []int {
    v := make([]int, 1)
    //rand.Seed(time.Now().Unix())

    for i := uint64(0); i < sz; i++ {
        val := rand.Intn(mod) 
        //v = append(v, val)
        v[i] = val
    }

    return v
}

type sort_int_interface []int

func (s sort_int_interface) Len() int {
    return len(s)
}

func (s sort_int_interface) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s sort_int_interface) Less(i, j int) bool {
    return s[i] < s[j]
}

func main() {
    //Generate a random array of length n
    vec := make_random_vec(1000000, 100)

    for i := 0; i < 250; i++ {
        var v []int
        v = append(v, vec...)
        //sort.Sort(sort_int_interface(v))
        sort.Stable(sort_int_interface(v))
    }
}
