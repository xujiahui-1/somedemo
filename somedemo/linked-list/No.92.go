package main

import "fmt"

func reverseBetween(head *ListNode, left int, right int) *ListNode {

	dummy := &ListNode{
		Next: head,
	}
	po := dummy

	for i := 0; i < left-1; i++ {
		po = po.Next
	}
	var pre, cur *ListNode = nil, po.Next
	for i := 0; i < right-left+1; i++ {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	fmt.Println(cur.Val)
	po.Next.Next = cur
	po.Next = pre
	return dummy.Next
}
