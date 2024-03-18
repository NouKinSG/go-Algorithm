package main

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	// Test case 1: Unsorted array
	arr1 := []int{5, 3, 8, 2, 1}
	expected1 := []int{1, 2, 3, 5, 8}
	bubbleSort(arr1)
	if !reflect.DeepEqual(arr1, expected1) {
		t.Errorf("BubbleSort failed for test case 1. Expected %v, got %v", expected1, arr1)
	}

	// Test case 2: Sorted array
	arr2 := []int{1, 2, 3, 4, 5}
	expected2 := []int{1, 2, 3, 4, 5}
	bubbleSort(arr2)
	if !reflect.DeepEqual(arr2, expected2) {
		t.Errorf("BubbleSort failed for test case 2. Expected %v, got %v", expected2, arr2)
	}

	// Test case 3: Empty array
	arr3 := []int{}
	expected3 := []int{}
	bubbleSort(arr3)
	if !reflect.DeepEqual(arr3, expected3) {
		t.Errorf("BubbleSort failed for test case 3. Expected %v, got %v", expected3, arr3)
	}

	// Test case 4: Array with duplicate elements
	arr4 := []int{5, 3, 8, 2, 1, 3, 5}
	expected4 := []int{1, 2, 3, 3, 5, 5, 8}
	bubbleSort(arr4)
	if !reflect.DeepEqual(arr4, expected4) {
		t.Errorf("BubbleSort failed for test case 4. Expected %v, got %v", expected4, arr4)
	}
}
