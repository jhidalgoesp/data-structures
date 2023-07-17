package data_structures

import (
	"fmt"
	"testing"
)

func TestDoubleLinkedList(t *testing.T) {
	tests := []struct {
		name     string
		actions  func(list *DoubleLinkedList[int])
		expected []int
	}{
		{
			name: "AddToTheFront",
			actions: func(list *DoubleLinkedList[int]) {
				list.AddToTheFront(1)
				list.AddToTheFront(2)
				list.AddToTheFront(3)
			},
			expected: []int{3, 2, 1},
		},
		{
			name: "AddToTheBack",
			actions: func(list *DoubleLinkedList[int]) {
				list.AddToTheBack(1)
				list.AddToTheBack(2)
				list.AddToTheBack(3)
			},
			expected: []int{1, 2, 3},
		},
		{
			name: "RemoveFromTheFront",
			actions: func(list *DoubleLinkedList[int]) {
				list.AddToTheFront(1)
				list.AddToTheFront(2)
				list.AddToTheFront(3)
				list.RemoveFromTheFront()
			},
			expected: []int{2, 1},
		},
		{
			name: "RemoveFromTheBack",
			actions: func(list *DoubleLinkedList[int]) {
				list.AddToTheBack(1)
				list.AddToTheBack(2)
				list.AddToTheBack(3)
				list.RemoveFromTheBack()
			},
			expected: []int{1, 2},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			list := NewDoubleLinkedList[int]()
			test.actions(&list)

			if list.size != len(test.expected) {
				t.Errorf("Expected size to be %d, got %d", len(test.expected), list.size)
			}

			iterator := list.CreateIterator()
			for _, value := range test.expected {
				if !iterator.HasNext() {
					t.Errorf("Expected iterator to have next value")
				}

				data := iterator.GetNext()
				if *data != value {
					t.Errorf("Expected %d, got %d", value, *data)
				}
			}

			if iterator.HasNext() {
				t.Errorf("Expected iterator to reach the end")
			}
		})
	}
}

func DoubleLinkedListExample() {
	list := NewDoubleLinkedList[int]()

	list.AddToTheFront(3)
	list.AddToTheFront(2)
	list.AddToTheFront(1)
	list.AddToTheBack(4)
	list.AddToTheBack(5)
	list.RemoveFromTheFront()
	list.RemoveFromTheBack()

	iterator := list.CreateIterator()

	for iterator.HasNext() {
		element := iterator.GetNext()
		fmt.Printf("Element is %+v\n", *element)
		/* Output:
		Element is 2
		Element is 3
		Element is 4
		*/
	}
}
