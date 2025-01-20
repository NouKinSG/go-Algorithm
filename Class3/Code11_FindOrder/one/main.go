package one

func findOrder(numCourses int, prerequisites [][]int) []int {

	// 入度数组
	coming := make([]int, numCourses)
	// 构建图
	grid := make([][]int, numCourses)

	// 统计入度
	for _, val := range prerequisites {
		coming[val[0]]++

		grid[val[1]] = append(grid[val[1]], val[0])
	}
	// 队列
	queue := []int{}
	//入度为 0的入队
	for i, v := range coming {
		if v == 0 {
			queue = append(queue, i)
		}
	}

	ans := []int{}
	for len(queue) > 0 {
		// 出队
		courseId := queue[0]
		queue = queue[1:]
		ans = append(ans, courseId)

		// 减度操作
		for _, v := range grid[courseId] {
			coming[v]--
			if coming[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	// 判断是否有环
	if len(ans) != numCourses {
		return []int{}
	}
	return ans
}
