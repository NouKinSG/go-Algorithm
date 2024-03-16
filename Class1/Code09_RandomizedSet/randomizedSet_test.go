package main

import (
	"testing"
)

func TestRandommizedSet(t *testing.T) {
	randomizedSet := Constructor()
	// 测试Insert操作
	if !randomizedSet.Insert(1) {
		t.Errorf("Insert(1) failed, expected true, got false")
	}
	if randomizedSet.Insert(1) {
		t.Errorf("Insert(1) failed, expected false, got true")
	}

	// 测试Remove操作
	if !randomizedSet.Remove(1) {
		t.Errorf("Remove(1) failed, expected true, got false")
	}
	if randomizedSet.Remove(1) {
		t.Errorf("Remove(1) failed, expected false, got true")
	}

	// 测试Insert和GetRandom操作
	randomizedSet.Insert(2)
	if val := randomizedSet.GetRandom(); val != 2 {
		t.Errorf("GetRandom() failed, expected 2, got %d", val)
	}

	// 因为GetRandom的结果依赖于随机性，对其进行有效测试较为复杂，可能需要考虑测试随机性的统计方法或模拟随机源。
}
