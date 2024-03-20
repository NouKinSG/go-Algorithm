package main

import "fmt"

func findAllPrerequisites(prerequisites [][]int, course int) []int {
	graph := make(map[int][]int)
	for _, p := range prerequisites {
		graph[p[0]] = append(graph[p[0]], p[1])
	}
	visited := make(map[int]bool)
	var result []int
	var dfs func(int)
	dfs = func(c int) {
		for _, pre := range graph[c] {
			if !visited[pre] {
				visited[pre] = true
				result = append(result, pre)
				dfs(pre)
			}
		}
	}

	visited[course] = true
	dfs(course)
	return result
}

func main() {
	prerequisites := [][]int{{4, 2}, {2, 1}, {1, 3}, {5, 8}}
	course := 3
	fmt.Println("前置课程:", findAllPrerequisites(prerequisites, course))
}
