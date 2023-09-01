package main

/**
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

示例 1：


输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]
示例 2：

输入：head = [1], n = 1
输出：[]
示例 3：

输入：head = [1,2], n = 1
输出：[1]


提示：

链表中结点的数目为 sz
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz


进阶：你能尝试使用一趟扫描实现吗？

**/

//Definition for singly-linked list.
type ListNodeq struct {
	Val  int
	Next *ListNode
}

//第一种比较糟糕的想法,可以通过先遍历链表一遍,记录每一个Node坐标,然后再次遍历时,修改倒数第N及N-1个元素的坐标
//func removeNthFromEnd(head *ListNode, n int) *ListNode {
//
//}

//如果想要一趟就实现就需要通过借助额外的空间记录每个元素的下标与指针
