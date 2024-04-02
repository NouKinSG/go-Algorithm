package main

import "testing"

func TestIsHuiWen(t *testing.T) {
	// Test case 1: Palindrome string
	s1 := "level"
	expected1 := true
	result1 := isHuiWen(s1)
	if result1 != expected1 {
		t.Errorf("isHuiWen failed for test case 1. Expected %v, got %v", expected1, result1)
	}

	// Test case 2: Non-palindrome string
	s2 := "hello"
	expected2 := false
	result2 := isHuiWen(s2)
	if result2 != expected2 {
		t.Errorf("isHuiWen failed for test case 2. Expected %v, got %v", expected2, result2)
	}

	// Test case 3: Empty string
	s3 := ""
	expected3 := true
	result3 := isHuiWen(s3)
	if result3 != expected3 {
		t.Errorf("isHuiWen failed for test case 3. Expected %v, got %v", expected3, result3)
	}

	// Test case 4: Single character string
	s4 := "a"
	expected4 := true
	result4 := isHuiWen(s4)
	if result4 != expected4 {
		t.Errorf("isHuiWen failed for test case 4. Expected %v, got %v", expected4, result4)
	}
}
