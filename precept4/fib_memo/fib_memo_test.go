// https://medium.com/@hackintoshrao/analyzing-benchmarks-with-ease-using-benchcmp-and-benchviz-golang-add607fc46d6
package main

import (
        "testing"
)

func BenchmarkFibmemo(b *testing.B) {
        // run the Fibmemo function b.N times
        k := make(map[int]int)
        k[0] = 0
        k[1] = 1
        for n := 0; n < b.N; n++ {

                Fibmemo(20, k)
        }
}
