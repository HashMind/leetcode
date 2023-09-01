package main

import "fmt"

func main23() {
	var linkListA = ListNode23{Val: 1, Next: &ListNode23{Val: 4, Next: &ListNode23{Val: 5, Next: nil}}}
	var linkListB = ListNode23{Val: 1, Next: &ListNode23{Val: 3, Next: &ListNode23{Val: 4, Next: nil}}}
	var linkListC = ListNode23{Val: 2, Next: &ListNode23{Val: 6, Next: nil}}
	linkListArray := make([]*ListNode23, 3)
	linkListArray[0] = &linkListA
	linkListArray[1] = &linkListB
	linkListArray[2] = &linkListC
	headerNode := mergeSortedLinkList(linkListArray)
	fmt.Println(headerNode.Val)

}

type ListNode23 struct {
	Val  int
	Next *ListNode23
}

//N个数组,最差的情况挨个排序,把排序N1,N2,然后再和N3,时间复杂度:n^k

//所有的N个数组同时排序,然后得到是那个数组里面的最小的那个

type MatchedNode struct {
	node  *ListNode23
	index int
}

func mergeSortedLinkList(lists []*ListNode23) *ListNode23 {
	var headerNode *ListNode23
	var patchNode *ListNode23
	for {
		hasPendingNode := false
		var minimumNode *MatchedNode
		for i := range lists {
			if lists[i] == nil {
				continue
			}
			hasPendingNode = true
			if minimumNode == nil {
				minimumNode = &MatchedNode{node: lists[i], index: i}
				if patchNode == nil {
					patchNode = lists[i]
				}
			}
			if lists[i].Val < minimumNode.node.Val {
				minimumNode = &MatchedNode{node: lists[i], index: i}
			}
		}
		if !hasPendingNode {
			break
		}
		if headerNode == nil {
			headerNode = minimumNode.node

		} else {
			patchNode.Next = minimumNode.node
			patchNode = patchNode.Next
		}
		lists[minimumNode.index] = minimumNode.node.Next

	}
	return headerNode
}
