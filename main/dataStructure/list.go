package main

import "fmt"

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
//输出: 1->4->3->2->5->NULL 反转链表中间一部分
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

// k个一组反转链表
func reverseKGroup(head *ListNode, k int) *ListNode {
	node := head
	for i := 0; i < k; i++ {
		if node == nil {
			return head
		}
		node = node.Next
	}
	newHead := reverse(head, node)
	head.Next = reverseKGroup(node, k)
	return newHead
}

// reverse and return new head
func reverse(first *ListNode, last *ListNode) *ListNode {
	var pre *ListNode
	for first != last {
		tmp := first.Next
		first.Next = pre
		pre = first
		first = tmp
	}
	return pre
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

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 1 {
		return lists[0]
	}
	if len(lists) == 0 {
		return nil
	}
	ans := lists[0]
	for i := 1; i < len(lists); i++ {
		ans = merge(ans, lists[i])
	}
	return ans
}

func merge(a *ListNode, b *ListNode) *ListNode {
	ans := &ListNode{}
	temp := ans
	for a != nil && b != nil {
		if a.Val < b.Val {
			ans.Next = a
			a = a.Next
		} else {
			ans.Next = b
			b = b.Next
		}
		ans = ans.Next
	}
	if a != nil {
		ans.Next = a
	} else {
		ans.Next = b
	}
	return temp.Next
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ans := head.Next
	cur := head
	if cur.Next != nil {
		tmp := cur.Next
		cur.Next = tmp.Next
		tmp.Next = cur

	}
	cur.Next = swapPairs(cur.Next)
	return ans
}

//138. 复制带随机指针的链表 链表深拷贝 为什么用递归要先赋值再递归呢
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

var m map[*Node]*Node

func copyRandomList(head *Node) *Node {
	m = make(map[*Node]*Node)
	ans := deepCopy(head)
	return ans
}

func deepCopy(node *Node) *Node {
	if node == nil {
		return nil
	}
	if nod, ok := m[node]; ok {
		return nod
	}
	copyNode := &Node{
		Val:    node.Val,
		Next:   nil,
		Random: nil,
	}
	//要先赋值再递归呢 ？？
	m[node] = copyNode
	copyNode.Next = deepCopy(node.Next)
	copyNode.Random = deepCopy(node.Random)
	return copyNode
}

//19. 删除链表的倒数第 N 个结点  双指针的精妙优化解法 todo
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{
		Val:  0,
		Next: head,
	}
	fast, slow := head, dummyHead
	//双指针让fast领先slow n个节点
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast != nil && fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummyHead.Next
}

func main() {
	fmt.Println(copyRandomList(&Node{
		Val: 3,
		Next: &Node{
			Val: 5,
			Next: &Node{
				Val:    7,
				Next:   nil,
				Random: nil,
			},
			Random: nil,
		},
		Random: nil,
	}))

	//fmt.Println(merge(&ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val:  2,
	//		Next: nil,
	//	},
	//}, &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 2,
	//		Next: &ListNode{
	//			Val: 3,
	//			Next: &ListNode{
	//				Val:  4,
	//				Next: &ListNode{Val: 5},
	//			},
	//		},
	//	},
	//}))

}
