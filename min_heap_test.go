package data_structures

import (
	"fmt"
	"testing"
)

func TestMinHeap(t *testing.T) {
	tests := []struct {
		name     string
		actions  func(heap *MinHeap[int])
		expected []int
	}{
		{
			name: "AddAndRemove",
			actions: func(heap *MinHeap[int]) {
				heap.Add(3)
				heap.Add(1)
				heap.Add(4)
				heap.Add(2)
				heap.Add(5)
				heap.Remove()
			},
			expected: []int{2, 3, 4, 5},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			heap, err := NewMinHeap[int](func(a, b int) bool {
				return a < b
			})
			if err != nil {
				t.Errorf("Expected the heap to be valid")
			}

			test.actions(heap)

			if heap.size != len(test.expected) {
				t.Errorf("Expected size to be %d, got %d", len(test.expected), heap.size)
			}

			for _, value := range test.expected {
				data := heap.Remove()
				if data == nil {
					t.Errorf("Expected non-nil value, got nil")
				} else if *data != value {
					t.Errorf("Expected %d, got %d", value, *data)
				}
			}

			if heap.size != 0 {
				t.Errorf("Expected size to be 0, got %d", heap.size)
			}
		})
	}
}

func MinHeapExample() {
	type person struct {
		age int
	}

	h, _ := NewMinHeap[person](func(a, b person) bool { return a.age < b.age })

	h.Add(person{5})
	h.Add(person{4})
	h.Add(person{6})
	h.Add(person{7})
	h.Add(person{3})
	h.Add(person{2})

	fmt.Println(h.Remove())
	fmt.Println(h.Remove())
	/* Output:
	Element is 2
	Element is 3
	*/
}
