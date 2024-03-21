package main

import (
	"reflect"
	"testing"
)

func TestFindAllPrerequisites(t *testing.T) {
	prerequisites := [][]int{{1, 2}, {2, 3}, {3, 4}}
	numCourses := 5

	expectedResult := []int{2, 3, 4}
	result := findAllPrerequisites(prerequisites, numCourses)

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("TestFindAllPrerequisites failed, expected %v but got %v", expectedResult, result)
	}
}

func TestFindAllPrerequisites_Empty(t *testing.T) {
	prerequisites := [][]int{}
	numCourses := 5

	expectedResult := []int{}
	result := findAllPrerequisites(prerequisites, numCourses)

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("TestFindAllPrerequisites_Empty failed, expected %v but got %v", expectedResult, result)
	}
}

func TestFindAllPrerequisites_Cyclic(t *testing.T) {
	prerequisites := [][]int{{1, 2}, {2, 3}, {3, 1}}
	numCourses := 4

	expectedResult := []int{2, 3, 1}
	result := findAllPrerequisites(prerequisites, numCourses)

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("TestFindAllPrerequisites_Cyclic failed, expected %v but got %v", expectedResult, result)
	}
}
