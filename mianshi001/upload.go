// status.go
package main

import (
	"fmt"
	"net/http"
)

func updateTaskStatus(taskID int, status string) error {
	// 模拟数据库更新操作 (假设操作)
	fmt.Printf("Task %d status updated to %s\n", taskID, status)
	return nil
}

func handleTaskCompletion(w http.ResponseWriter, r *http.Request) {
	taskID := 1
	// 确保在更新数据库后再返回结果
	err := updateTaskStatus(taskID, "completed")
	if err != nil {
		http.Error(w, "任务更新失败", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("任务完成"))
}
