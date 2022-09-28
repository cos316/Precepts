// Creates a stack data structure for ints
package stack

import (
//       "fmt"
)

// ItemStack the stack of Items
type ItemStack struct {
        items []int // holds the items
        n     int   // number of items in stack
}

// New creates a new ItemStack of given capacity
func (s *ItemStack) New(capacity int) *ItemStack {
        s.items = make([]int, capacity)
        s.n = 0
        return s
}

// Push adds an int to the top of the stack
func (s *ItemStack) Push(t int) {
        // allocate space for a new array if needed
        if s.IsFull() {
                newItems := make([]int, s.n+1)
                for i := 0; i < s.n; i++ {
                        newItems[i] = s.items[i]
                }
                s.items = newItems
        }

        s.items[s.n] = t
        s.n = s.n + 1
}

// Pop removes an int from the top of the stack
func (s *ItemStack) Pop() int {
        item := s.items[s.n-1]
        return item
}

// Top returns the top of the stack
func (s *ItemStack) Top() int {
        item := s.items[s.n-1]
        return item
}

// IsFull return true if stack is at capacity
func (s *ItemStack) IsFull() bool {
        if s.n == len(s.items) {
                return true
        } else {
                return false
        }
}

// IsEmpty return true if stack is empty
func (s *ItemStack) IsEmpty() bool {
        if s.n == 0 {
                return true
        } else {
                return false
        }

}

// current number of elements in stack
func (s *ItemStack) Size() int {
        return s.n
}
