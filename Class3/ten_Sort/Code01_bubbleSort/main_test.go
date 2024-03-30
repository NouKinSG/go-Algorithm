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

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode = nil
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre

		pre = cur
		cur = next
	}
	return pre
}

func TestReverseList(t *testing.T) {
	// Test case 1: Single node list
	head1 := &ListNode{Val: 1, Next: nil}
	expected1 := &ListNode{Val: 1, Next: nil}
	reversed1 := reverseList(head1)
	if !reflect.DeepEqual(reversed1, expected1) {
		t.Errorf("ReverseList failed for test case 1. Expected %v, got %v", expected1, reversed1)
	}

	// Test case 2: Multiple node list
	head2 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}
	expected2 := &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: nil}}}
	reversed2 := reverseList(head2)
	if !reflect.DeepEqual(reversed2, expected2) {
		t.Errorf("ReverseList failed for test case 2. Expected %v, got %v", expected2, reversed2)
	}

	// Test case 3: Empty list
	var head3 *ListNode = nil
	var expected3 *ListNode = nil
	reversed3 := reverseList(head3)
	if !reflect.DeepEqual(reversed3, expected3) {
		t.Errorf("ReverseList failed for test case 3. Expected %v, got %v", expected3, reversed3)
	}
}

func quickSort(arr []int, low, high int) {
	if low < high {
		pivot := partition(arr, low, high)
		quickSort(arr, low, pivot-1)
		quickSort(arr, pivot+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func TestQuickSort(t *testing.T) {
	// Test case 1: Unsorted array
	arr1 := []int{9, 5, 1, 8, 3, 7, 4, 6, 2}
	expected1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	quickSort(arr1, 0, len(arr1)-1)
	if !reflect.DeepEqual(arr1, expected1) {
		t.Errorf("QuickSort failed for test case 1. Expected %v, got %v", expected1, arr1)
	}

	// Test case 2: Sorted array
	arr2 := []int{1, 2, 3, 4, 5}
	expected2 := []int{1, 2, 3, 4, 5}
	quickSort(arr2, 0, len(arr2)-1)
	if !reflect.DeepEqual(arr2, expected2) {
		t.Errorf("QuickSort failed for test case 2. Expected %v, got %v", expected2, arr2)
	}

	// Test case 3: Empty array
	arr3 := []int{}
	expected3 := []int{}
	quickSort(arr3, 0, len(arr3)-1)
	if !reflect.DeepEqual(arr3, expected3) {
		t.Errorf("QuickSort failed for test case 3. Expected %v, got %v", expected3, arr3)
	}

	// Test case 4: Array with duplicate elements
	arr4 := []int{5, 3, 8, 2, 1, 3, 5}
	expected4 := []int{1, 2, 3, 3, 5, 5, 8}
	quickSort(arr4, 0, len(arr4)-1)
	if !reflect.DeepEqual(arr4, expected4) {
		t.Errorf("QuickSort failed for test case 4. Expected %v, got %v", expected4, arr4)
	}
}
