package main

import "fmt"

//delete node

type ListNode struct {
	Val  int
	Next *ListNode
}

func Shownode(p *ListNode) { //遍历
	for p != nil {
		fmt.Println(*p)
		p = p.Next //移动指针
	}
}

func removeElements(head *ListNode, val int) *ListNode {
	cur := &ListNode{Next: head}
	for tmp := cur; tmp.Next != nil; {
		if tmp.Next.Val == val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}
	return cur.Next
}

func addnode() {
	var head = new(ListNode)
	head.Val = 0
	var cur *ListNode
	cur = head //tail用于记录头结点的地址，刚开始tail的的指针指向头结点
	for i := 1; i < 10; i++ {
		node := ListNode{Val: i}
		node.Next = cur //将新插入的node的next指向头结点
		cur = &node     //重新赋值头结点
	}
}
