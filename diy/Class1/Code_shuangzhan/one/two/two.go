package main

//
///*
//牛牛有一个长度为n 的链表 a1,a2,..,an 和一个长度为m 的链表 b,b2,..,bm。这两个链表含有公共的前缀节点和公共的后缀节点，其可以表示成如下的形式:
//
//*/
//
//// ListNode 定义
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
///**
// * 合并链表的公共前缀和公共后缀部分
// * @param a ListNode类 链表a
// * @param b ListNode类 链表b
// * @return ListNode类 合并后的链表
// */
//
//func mergeList(a *ListNode, b *ListNode) *ListNode {
//	// 提取公共前缀，并同时更新 a 和 b 的指针位置
//	prefix, aNew, bNew := extractPrefix(a, b)
//
//	// 获取链表长度
//	aLen := getLength(aNew)
//	bLen := getLength(bNew)
//
//	// 对齐两个链表的起始点
//	if aLen > bLen {
//		aNew = advanceBy(aNew, aLen-bLen)
//	} else if bLen > aLen {
//		bNew = advanceBy(bNew, bLen-aLen)
//	}
//
//	// 找到公共后缀起点
//	var common *ListNode
//	for aNew != nil && bNew != nil {
//		if aNew == bNew { // 找到公共后缀
//			common = aNew
//			break
//		}
//		aNew = aNew.Next
//		bNew = bNew.Next
//	}
//
//	// 拼接前缀和后缀
//	if prefix != nil {
//		last := prefix
//		for last.Next != nil {
//			last = last.Next
//		}
//		last.Next = common
//		return prefix
//	}
//	return common
//}
