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

func findOrder1(numCourses int, prerequisites [][]int) []int {
	// 入度
	indegrees := make([]int, numCourses)

	// 从某个课程出发，可以学的课程
	graph := make([][]int, numCourses)

	// 统计入度,统计学了某个课后 可以学的课程
	for _, p := range prerequisites {
		indegrees[p[0]]++
		graph[p[1]] = append(graph[p[1]], p[0])
	}

	// 定义 队列
	queue := []int{}

	// 入度为 0 的入队
	for index, val := range indegrees {
		if val == 0 {
			queue = append(queue, index)
		}
	}

	// 存储结果
	ans := []int{}
	for len(queue) > 0 {
		// 出队 放入结果
		course := queue[0]
		queue = queue[1:]
		ans = append(ans, course)

		// 减少相关联的入度
		for _, val := range graph[course] {
			indegrees[val]--

			// 入度减为0后，把入度为0的课程号 入队
			if indegrees[val] == 0 {
				queue = append(queue, val)
			}
		}
	}

	// 有环
	if len(ans) == 0 {
		return []int{}
	}

	return ans
}
