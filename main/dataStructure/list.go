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

//编写程序以 x 为基准分割链表，使得所有小于 x 的节点排在大于或等于 x 的节点之前。如果链表中包含 x，x 只需出现在小于 x 的元素之后(如下所示)。
//分割元素 x 只需处于“右半部分”即可，其不需要被置于左右两部分之间。
// 输入: head = 3->5->8->5->10->2->1, x = 5
//输出: 3->1->2->10->5->5->8
func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var temp = head.Next
	var before = head
	for temp != nil {
		if temp.Val < x {
			// todo 链表操作的关键在于赋值执行的顺序！ 不要覆盖/
			before.Next = temp.Next
			cur := temp
			cur.Next = head
			head = cur
			temp = before.Next
			continue
		}
		before = temp
		temp = temp.Next
	}
	return head
}
func main() {
	fmt.Println(partition(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  2,
					Next: &ListNode{Val: 5},
				},
			},
		},
	}, 3))

}
