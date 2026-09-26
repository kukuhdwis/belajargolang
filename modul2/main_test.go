package main

import (
	"reflect"
	"testing"
)

func TestSwap(t *testing.T) {
	slice := []int{10, 20, 30}
	Swap(slice, 0)
	expected := []int{20, 10, 30}
	if !reflect.DeepEqual(slice, expected) {
		t.Errorf("Swap failed: expected %v, got %v", expected, slice)
	}

	Swap(slice, 1)
	expected = []int{20, 30, 10}
	if !reflect.DeepEqual(slice, expected) {
		t.Errorf("Swap failed: expected %v, got %v", expected, slice)
	}
}

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Already sorted",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Reverse order",
			input:    []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:     "Random order with negatives and duplicates",
			input:    []int{5, -2, 8, 1, 3, -2, 0, 9},
			expected: []int{-2, -2, 0, 1, 3, 5, 8, 9},
		},
		{
			name:     "Single element",
			input:    []int{42},
			expected: []int{42},
		},
		{
			name:     "Empty slice",
			input:    []int{},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := make([]int, len(tt.input))
			copy(data, tt.input)
			BubbleSort(data)
			if !reflect.DeepEqual(data, tt.expected) {
				t.Errorf("BubbleSort() = %v, want %v", data, tt.expected)
			}
		})
	}
}
