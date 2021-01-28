package main

import (
	"fmt"
)

type ListNode struct {
	data int
	next *ListNode
}

func reverseLinkedList(root *ListNode) *ListNode {
	if root == nil {
		return nil
	}
	tail := root
	curr := root.next
	tail.next = nil
	for next != nil {
		tmp := curr.next
		curr.next = head
		curr = tmp
	}
	return curr
}
