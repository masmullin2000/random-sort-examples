package main

import (
    //"fmt"
    "math/rand"
    //"time"
    "sort"
)

func make_random_vec(sz uint64, mod int64) []int64 {
    v := make([]int64, 1)
    //rand.Seed(time.Now().Unix())

    for i := uint64(0); i < sz; i++ {
        val := rand.Int63n(mod) 
        v = append(v, val)
    }

    return v
}

type sort_int_interface []int64

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
    //Provide seed

    //Generate a random array of length n
    vec := make_random_vec(1000000, 100)

    for i := 0; i < 1000; i++ {
        var v []int64
        v = append(v, vec...)
        //sort.Sort(sort_int_interface(v))
        sort.Stable(sort_int_interface(v))
    }
    //fmt.Println(vec)
}
