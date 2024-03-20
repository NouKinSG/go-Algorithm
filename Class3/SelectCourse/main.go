package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	// 记录入度, indegress 的索引是课程，元素是入度的数量
	indegrees := make([]int, numCourses)

	// 构建图，即从每个顶点出发可以到达的顶点列表
	// graph的索引也是课程号，是先决条件的课程号，意思是 第 index 个切片，里的元素是可选课程{}
	graph := make([][]int, numCourses)

	for _, p := range prerequisites {
		// p[0] 这门课的入度 加一
		indegrees[p[0]]++
		// p[0] 课程号，p[1] 前修课程号。
		graph[p[1]] = append(graph[p[1]], p[0])
	}

	// 找到所有入度为 0 的顶点
	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if indegrees[i] == 0 {
			queue = append(queue, i)
		}
	}

	// 存储排序结果
	order := []int{}
	for len(queue) > 0 {
		// 从队列取出一个顶点
		// 出队
		course := queue[0]
		queue = queue[1:]

		order = append(order, course)

		// 减少与该顶点相连的顶点入度
		for _, next := range graph[course] {
			indegrees[next]--
			// 如果入度变为0，则加入队列
			if indegrees[next] == 0 {
				queue = append(queue, next)
			}
		}
	}
	// 如果排序结果中的顶点数量，不等于课程总数，说明存在环，返回空数组
	if len(order) != numCourses {
		return []int{}
	}

	return order
}
