package main

import (
	"fmt"
)

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// head -> node1 -> node2
// node1
// head <- node1 <- node2
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	l := head
	r := head.Next
	for r != nil {
		// 类似交换赋值的写法 不要把自己绕进去
		tmp := r.Next
		r.Next = l
		l = r
		r = tmp
	}
	head.Next = nil
	return l
}

// 输入: 1->2->3->4->5->NULL, m = 2, n = 4
//输出: 1->4->3->2->5->NULL
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || m == n {
		return head
	}
	// 跳到反转需要改变的节点前
	l := head
	for i := 0; i < m-2; i++ {
		l = l.Next
	}
	// 如果 m==1从头开始反转 可以先造个假头
	dummyhead := &ListNode{
		Val:  -1,
		Next: head,
	}
	if m == 1 {
		l = dummyhead
	}
	// 用来接反转完的链表
	sta := l
	end := sta.Next
	//正常反转链表操作
	r := l.Next
	for i := 0; i <= n-m; i++ {
		if r == nil {
			break
		}
		tmp := r.Next
		r.Next = l
		l = r
		r = tmp
	}
	// 反转完接回去
	sta.Next = l
	end.Next = r
	if m == 1 {
		return l
	}
	return head
}
func main() {
	fmt.Println(reverseBetween(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: &ListNode{Val: 5},
				},
			},
		},
	}, 1, 5))

}
