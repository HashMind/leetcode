package main

import "fmt"

type ListNode25 struct {
	Val  int
	Next *ListNode25
}

var headNode *ListNode25 //等于第一个反转的reverseTail
var reverseHead *ListNode25
var reverseTail *ListNode25

/*
*
给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
k 是一个正整数，它的值小于或等于链表的长度。
如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
*
*/

func reverseKGroup(head *ListNode25, k int) {
	var newListHead *ListNode25
	reverseHead = head
	reverseTail = head.Next
	var seamTail *ListNode25
	var seamNextTail *ListNode25
	reverseCount := 0
	var reverseOriginHeader *ListNode25 = head
	for reverseTail != nil {
		var reverseNextTail *ListNode25
		for ; reverseTail != nil && reverseCount <= k-2; reverseCount++ {
			reverseNextTail = reverseTail.Next
			reverseTail.Next = reverseHead
			reverseHead = reverseTail
			reverseTail = reverseNextTail
		}
		seamNextTail = reverseNextTail
		if newListHead == nil {
			newListHead = reverseHead
		}

		if seamTail == nil {
			seamTail = head
		} else {
			seamTail.Next = reverseHead
			seamTail = seamNextTail
		}

		if reverseNextTail != nil {
			reverseCount = 0
			reverseHead = reverseNextTail
			reverseTail = reverseNextTail.Next
			reverseOriginHeader = reverseHead
		} else {
			reverseOriginHeader.Next = nil
		}
	}
	fmt.Println(newListHead)
}

func main() {
	listNodeHead := ListNode25{Val: 1,
		Next: &ListNode25{Val: 2,
			Next: &ListNode25{Val: 3,
				Next: &ListNode25{Val: 4,
					Next: &ListNode25{Val: 5,
						Next: &ListNode25{Val: 6,
							Next: &ListNode25{Val: 7,
								Next: &ListNode25{Val: 8,
									Next: &ListNode25{Val: 9,
										Next: &ListNode25{Val: 10,
											Next: nil}}}}}}}}}}
	reverseKGroup(&listNodeHead, 5)
}
