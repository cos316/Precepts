// Fib computes the n'th number in the Fibonacci series.
package main

func Fib(n int) int {
        switch n {
        case 0:
                return 0
        case 1:
                return 1
        default:
                return Fib(n-1) + Fib(n-2)
        }
}
