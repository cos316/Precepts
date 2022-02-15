// Test cases for mysort.go
package mysort

import (
    "testing"  // import to use Go's testing framework
    "sort"
)

// Generates a slice of size, size filled with some numbers 
func generateSlice(size int) []int {
    slice := make([]int, size, size)
    for i := 0; i < size; i++ {
        slice[i] = i;
	if i % 2 == 0 {
	   slice[i] = -i;
	}
    }
    return slice
}

// test MergeSort - use Golang sort package to test results
func TestMergeSort(t *testing.T) {
    slice := generateSlice(1000)
    MergeSort(slice)
    if !sort.IntsAreSorted(slice) {
        t.Error("Did not sort array properly")
    }
}

// test InsertSort - use Golang sort package to test results
func TestInsertionSort(t *testing.T) {
    slice := generateSlice(1000)
    InsertionSort(slice)
    if !sort.IntsAreSorted(slice) {
        t.Error("Did not sort array properly")
    }
}

