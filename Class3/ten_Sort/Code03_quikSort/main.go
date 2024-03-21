package main

import "fmt"

func findAllPrerequisites(prerequisites [][]int, numCourses int) []int {
	graph := make(map[int][]int)
	// 反转图的构建方向，正确表示依赖关系
	for _, pre := range prerequisites {
		graph[pre[1]] = append(graph[pre[1]], pre[0])
	}

	visited := make(map[int]bool)
	var result []int

	// DFS函数，用于遍历图
	var dfs func(int)
	dfs = func(course int) {
		if visited[course] {
			return
		}
		visited[course] = true
		for _, next := range graph[course] {
			dfs(next)
			// 添加到结果中，如果它还未存在于结果中
			if !contains(result, next) {
				result = append(result, next)
			}
		}
	}

	// 检查结果数组中是否已包含该元素
	contains := func(slice []int, item int) bool {
		for _, s := range slice {
			if s == item {
				return true
			}
		}
		return false
	}

	// 对每个可能的起点执行DFS
	for i := 0; i < numCourses; i++ {
		dfs(i)
	}

	return result
}

// 检查切片中是否包含指定的元素
func contains(slice []int, item int) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func main() {
	prerequisites := [][]int{{4, 2}, {2, 1}, {1, 3}, {5, 8}}
	course := 3
	fmt.Println("前置课程:", findAllPrerequisites(prerequisites, course))
}
