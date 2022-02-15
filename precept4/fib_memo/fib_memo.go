// Fib computes the n'th number in the Fibonacci series.
// https://medium.com/@hackintoshrao/analyzing-benchmarks-with-ease-using-benchcmp-and-benchviz-golang-add607fc46d6
package main

func Fibmemo(n int, fibMap map[int]int) int {

        if val, ok := fibMap[n]; ok {
                return val
        }
        fibMap[n] = Fibmemo(n-1, fibMap) + Fibmemo(n-2, fibMap)

        return fibMap[n]
}
