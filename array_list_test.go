package data_structures

import "testing"

func TestArrayList(t *testing.T) {
	tests := []struct {
		name     string
		actions  func(*ArrayList[int])
		expected []int
	}{
		{
			name: "AddToFront",
			actions: func(list *ArrayList[int]) {
				list.AddToFront(1)
				list.AddToFront(2)
				list.AddToFront(3)
				list.AddToFront(4)
				list.AddToFront(5)
				list.AddToFront(6)
				list.AddToFront(7)
				list.AddToFront(8)
				list.AddToFront(9)
				list.AddToFront(10)
				list.AddToFront(11)
				list.AddToFront(12)
			},
			expected: []int{12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
		{
			name: "AddToBack",
			actions: func(list *ArrayList[int]) {
				list.AddToBack(1)
				list.AddToBack(2)
				list.AddToBack(3)
				list.AddToBack(4)
				list.AddToBack(5)
				list.AddToBack(6)
				list.AddToBack(7)
				list.AddToBack(8)
				list.AddToBack(9)
				list.AddToBack(10)
				list.AddToBack(11)
				list.AddToBack(12)
			},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		},
		{
			name: "RemoveFromFront",
			actions: func(list *ArrayList[int]) {
				list.AddToFront(1)
				list.AddToFront(2)
				list.AddToFront(3)
				list.RemoveFromFront()
			},
			expected: []int{2, 1},
		},
		{
			name: "RemoveFromBack",
			actions: func(list *ArrayList[int]) {
				list.AddToBack(4)
				list.AddToBack(5)
				list.RemoveFromBack()
			},
			expected: []int{4},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			list := NewArrayList[int]()

			test.actions(&list)

			if list.size != len(test.expected) {
				t.Errorf("Expected size to be %d, but got %d", len(test.expected), list.size)
			}

			for i := range list.backingArray {
				if list.backingArray[i] != nil && *list.backingArray[i] != test.expected[i] {
					t.Errorf("Expected element at index %d to be %d, but got %d", i, test.expected[i], *list.backingArray[i])
				}
			}
		})
	}
}
