package main

import "fmt"

func main02() {
	l1 := &ListNode{Val: 9, Next: &ListNode{6, &ListNode{5, &ListNode{9, nil}}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	resultNode := addTwoNumbers(l1, l2)
	printListNode(resultNode)
}

func printListNode(node *ListNode) {
	if node.Next != nil {
		fmt.Printf("%d,", node.Val)
		printListNode(node.Next)
	} else {
		fmt.Printf("%d\n", node.Val)
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var resultNode = new(ListNode)
	var sum = 0
	var lNode *ListNode
	var rNode *ListNode
	if l1 != nil {
		sum += l1.Val
		lNode = l1.Next
	}
	if l2 != nil {
		sum += l2.Val
		rNode = l2.Next
	}

	if sum > 9 {
		resultNode.Val = sum - 10
		if lNode == nil {
			lNode = new(ListNode)
		}
		lNode.Val += 1
	} else {
		resultNode.Val = sum
	}
	if lNode != nil || rNode != nil {
		resultNode.Next = addTwoNumbers(lNode, rNode)
	}
	return resultNode
}
