package main

import (
	"testing"
)

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return myquick(x, n)
	}
	return 1.0 / myquick(x, n)
}

func myquick(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := myquick(x, n/2)
	if n%2 == 0 {
		return y * y
	}

	return y * y * x
}

func TestMyPow(t *testing.T) {
	// Test case 1: x = 2, n = 3
	result := myPow(2, 3)
	expected := 8.0
	if result != expected {
		t.Errorf("Test case 1 failed: expected %f, got %f", expected, result)
	}

	// Test case 2: x = 3, n = 0
	result = myPow(3, 0)
	expected = 1.0
	if result != expected {
		t.Errorf("Test case 2 failed: expected %f, got %f", expected, result)
	}

	// Test case 3: x = 5, n = -2
	result = myPow(5, -2)
	expected = 0.04
	if result != expected {
		t.Errorf("Test case 3 failed: expected %f, got %f", expected, result)
	}

	// Add more test cases here...
}
