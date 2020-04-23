package main

import (
	"container/list"
	"fmt"
)

//leetcode submit region begin(Prohibit modification and deletion)
type MyStack struct {
	  l *list.List
}


/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		l: list.New(),
	}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	this.l.PushBack(x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	back := this.l.Back()
	ans := back.Value.(int)
	this.l.Remove(back)
	return ans
}


/** Get the top element. */
func (this *MyStack) Top() int {
	return this.l.Back().Value.(int)
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.l.Len()==0
}


  type ListNode struct {
      Val int
      Next *ListNode
  }
  // 连上指针 空判断
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	temp := &ListNode{
		Val:  0,
		Next: nil,
	}
	flag := 0
	var num,num1,num2 int
	for i:=0; true ; temp = temp.Next  {

		if l1==nil {
			num1 = 0
		} else {
			num1 = l1.Val
		}
		if l2==nil {
			num2=0
		} else {
			num2 = l2.Val
		}

		num = num1 + num2 + flag
		flag = 0
		if num == 0 && l1==nil && l2==nil {
			break
		}
		if num >= 10 {
			flag = 1
			num -= 10
		}
		node :=&ListNode{
			Val:  num,
			Next: nil,
		}
		temp.Next = node
		if i==0 {
			head = node
		}

		if l1!=nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		i++
	}
	if head == nil {
		head = &ListNode{
			Val:  0,
			Next: nil,
		}
	}
	return head
}

func main() {
	fmt.Println("jh")
	a := &ListNode{
		Val:  2,
		Next: &ListNode{
			Val:  4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	b := &ListNode{
		Val:  5,
		Next: &ListNode{
			Val:  6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	addTwoNumbers(a,b)
	//mystack := Constructor()
	////mystack.Push(1)
	//mystack.Push(2)
	//fmt.Println(mystack.l.Len())
	//fmt.Println(mystack.Top())
	//fmt.Println(mystack.Pop())
	//fmt.Println(mystack.Empty())
	//fmt.Println("hello atty")
}
