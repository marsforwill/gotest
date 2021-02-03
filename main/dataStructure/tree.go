package main

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

//* Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func midfs(root *TreeNode, ans *[]int) {
	if root.Left != nil {
		midfs(root.Left, ans)
	}
	*ans = append(*ans, root.Val)
	if root.Right != nil {
		midfs(root.Right, ans)
	}
	return
}

/**
情况 1，如果p和q都在以root为根的树中，函数返回的即使p和q的最近公共祖先节点。

情况 2，那如果p和q都不在以root为根的树中怎么办呢？函数理所当然地返回null呗。

情况 3，那如果p和q只有一个存在于root为根的树中呢？函数就会返回那个节点。

题目说了输入的p和q一定存在于以root为根的树中，但是递归过程中，以上三种情况都有可能发生，所以说这里要定义清楚，后续这些定义都会在代码中体现。

！！！！OK，第一个问题就解决了，把这个定义记在脑子里，无论发生什么，都不要怀疑这个定义的正确性，这是我们写递归函数的基本素养
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left == nil && right == nil {
		return nil
	}
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	} else {
		return left
	}

}

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	pl := getPath(p, root)
	i := len(pl) - 1
	ql := getPath(q, root)
	j := len(ql) - 1
	if i <= 0 || j <= 0 {
		return root
	}
	for i > 0 && j > 0 {
		if pl[i-1] != ql[j-1] {
			return &pl[i]
		}
		i--
		j--
	}
	return &pl[i]
}

func getPath(p *TreeNode, root *TreeNode) []TreeNode {
	var ans []TreeNode
	dfsNode(p, root, &ans)
	return ans
}

func dfsNode(p *TreeNode, root *TreeNode, ans *[]TreeNode) bool {
	if p == root {
		*ans = append(*ans, *p)
		return true
	}
	if p == nil {
		return false
	}
	if root.Left != nil && dfsNode(p, root.Left, ans) {
		*ans = append(*ans, *root)
		return true
	}
	if root.Right != nil && dfsNode(p, root.Right, ans) {
		*ans = append(*ans, *root)
		return true
	}
	return false
}

func BSTSequences2(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{{}}
	}
	var ret [][]int
	// 初始调用
	helper([]*TreeNode{root}, nil, &ret)
	return ret
}

// ret保存答案
// nodes 和 last 维护对应的状态，未处理的node 和 已存在的value last
// 每一次递归调用都会处理掉一个node 继续递归状态转移 直到 node处理完 存答案
func helper(node []*TreeNode, last []int, ret *[][]int) {
	// 递归出口 没有要处理的node
	if len(node) == 0 {
		*ret = append(*ret, last)
	}

	// 对当前未处理的每一个node
	for i, treeNode := range node {

		var newSlice []int = make([]int, len(last))
		copy(newSlice, last)
		//把当前节点val加入slice
		newSlice = append(newSlice, treeNode.Val)

		newNodes := make([]*TreeNode, len(node))
		copy(newNodes, node)
		// 剔除掉当前node
		newNodes = append(newNodes[:i], newNodes[i+1:]...)
		// 添加左右子节点 这就保证了 子节点的处理都在父节点后
		if treeNode.Left != nil {
			newNodes = append(newNodes, treeNode.Left)
		}
		if treeNode.Right != nil {
			newNodes = append(newNodes, treeNode.Right)
		}
		// 对新的nodes 和 slice继续递归遍历 直到 nodes == 0，将slice加入ans ret
		helper(newNodes, newSlice, ret)
	}

}

func BSTSequences(root *TreeNode) [][]int {
	var res [][]int
	var value []int
	if root == nil {
		return [][]int{{}}
	}
	dfsBST([]*TreeNode{root}, value, &res)
	return res
}

func dfsBST(nodes []*TreeNode, value []int, res *[][]int) {

	if len(nodes) == 0 {
		*res = append(*res, value)
		return
	}
	for i := 0; i < len(nodes); i++ {
		newValue := make([]int, len(value))
		copy(newValue, value)
		newValue = append(newValue, nodes[i].Val)

		newNodes := make([]*TreeNode, len(nodes))
		copy(newNodes, nodes)
		newNodes = append(newNodes[:i], newNodes[i+1:]...)
		if nodes[i].Left != nil {
			newNodes = append(newNodes, nodes[i].Left)
		}
		if nodes[i].Right != nil {
			newNodes = append(newNodes, nodes[i].Right)
		}
		dfsBST(newNodes, newValue, res)
	}
	return
}

//设计一个算法，找出二叉搜索树中指定节点的“下一个”节点（也即中序后继）。
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {

	var ans []*TreeNode
	midfs2(root, &ans)
	for i := 0; i < len(ans); i++ {
		if ans[i] == p && i < len(ans)-1 {
			return ans[i+1]
		}
	}
	return nil

}

func midfs2(root *TreeNode, ans *[]*TreeNode) {
	if root.Left != nil {
		midfs2(root.Left, ans)
	}
	*ans = append(*ans, root)
	if root.Right != nil {
		midfs2(root.Right, ans)
	}
	return
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var num []int
	dfsMid(root, &num)
	for i := 1; i < len(num); i++ {
		if num[i] <= num[i-1] {
			return false
		}
	}
	return true
}

func dfsMid(root *TreeNode, num *[]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		dfsMid(root.Left, num)
	}
	*num = append(*num, root.Val)
	if root.Right != nil {
		dfsMid(root.Right, num)
	}
}

//输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
// 前序遍历 preorder = [3 | 9 | 20,15,7]
//中序遍历 inorder = [9 | 3 | 15,20,7]
//  3
// / \
//9  20
//  /  \
// 15   7
func buildTree1(preorder []int, inorder []int) *TreeNode {
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			return &TreeNode{
				Val:   preorder[0],
				Left:  buildTree1(preorder[1:i+1], inorder[0:i]),
				Right: buildTree1(preorder[i+1:], inorder[i+1:]),
			}
		}
	}
	return nil
}

//106. 从中序与后序遍历序列构造二叉树
//中序遍历 inorder = [9,|3,|15,20,7]
//后序遍历 postorder = [9,|15,7,20,|3]
func buildTree(inorder []int, postorder []int) *TreeNode {
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == postorder[len(postorder)-1] {
			return &TreeNode{
				Val:   inorder[i],
				Left:  buildTree(inorder[0:i], postorder[0:i]),
				Right: buildTree(inorder[i+1:], postorder[i:len(postorder)-1]),
			}
		}
	}
	return nil
}

//输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A != nil && B != nil {
		return dfsIsSub(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
	}
	return false
}

func dfsIsSub(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil || A.Val != B.Val {
		return false
	}
	return dfsIsSub(A.Left, B.Left) && dfsIsSub(A.Right, B.Right)
}

// 二叉树非递归后序遍历 (可以通过前序遍历逻辑逆转)
func postorderTraversal(root *TreeNode) []int {
	stack := list.New()
	var ans []int
	tmp := root
	for tmp != nil || stack.Len() > 0 {
		if tmp != nil {
			stack.PushBack(tmp)
			ans = append([]int{tmp.Val}, ans...)
			tmp = tmp.Right
		} else {
			f := stack.Back()
			stack.Remove(f)
			tmp = f.Value.(*TreeNode).Left
		}
	}
	return ans
}

// 二叉树非递归前序遍历
func preorderTraversal(root *TreeNode) []int {
	stack := list.New() //用list模拟栈 push pop
	var ans []int
	tmp := root
	for stack.Len() > 0 || tmp != nil {
		if tmp != nil {
			stack.PushBack(tmp)
			ans = append(ans, tmp.Val)
			tmp = tmp.Left
		} else {
			f := stack.Back()
			stack.Remove(f)
			tmp = f.Value.(*TreeNode).Right
		}
	}
	return ans
}

// 二叉树非递归中序遍历
func inorderTraversal(root *TreeNode) []int {
	stack := list.New()
	tmp := root
	var ans []int
	for tmp != nil || stack.Len() > 0 {
		if tmp != nil {
			stack.PushBack(tmp)
			tmp = tmp.Left
		} else {
			b := stack.Back()
			ans = append(ans, b.Value.(*TreeNode).Val)
			stack.Remove(b)
			tmp = b.Value.(*TreeNode).Right
		}
	}
	return ans
}

//将BST转化为双向循环链表，不允许新建节点
//为防止歧义，左指针表示双链表向前指，右指针表示双链表向后指
var pre *TreeNode //必须在全局变量上才可以实现
func treeToDoublyList(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	dfsTransfer(root)
	head, tail := root, root
	for head.Left != nil {
		head = head.Left
	}
	for tail.Right != nil {
		tail = tail.Right
	}
	head.Left = tail
	tail.Right = head
	return head
}

func dfsTransfer(cur *TreeNode) {
	if cur == nil {
		return
	}
	dfsTransfer(cur.Left)
	if pre != nil {
		cur.Left = pre
		pre.Right = cur
	}
	pre = cur
	dfsTransfer(cur.Right)
}

// 二叉树前序遍历 序列化与反序列化
// Serializes a tree to a single string.
func serialize(root *TreeNode) string {
	if root == nil {
		return "X"
	}
	return strconv.Itoa(root.Val) + "." + serialize(root.Left) + "." + serialize(root.Right)
}

// Deserializes your encoded data to tree.
func deserialize(data string) *TreeNode {
	list := strings.Split(data, ".")
	return buildT(&list)
}

// 这个有点精髓
func buildT(list *[]string) *TreeNode {
	rootVal := (*list)[0]
	*list = (*list)[1:]
	if rootVal == "X" {
		return nil
	}
	val, _ := strconv.Atoi(rootVal)
	return &TreeNode{
		Val:   val,
		Left:  buildT(list),
		Right: buildT(list),
	}
}

//109. 有序链表转换二叉搜索树 分治加中序遍历
var iter *ListNode

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	tail := head
	iter = head
	leng := 0
	for tail.Next != nil {
		leng++
		tail = tail.Next
	}
	return buildBT(0, leng)
}

func buildBT(left int, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right + 1) / 2
	head := &TreeNode{}
	head.Left = buildBT(left, mid-1)
	head.Val = iter.Val
	iter = iter.Next
	head.Right = buildBT(mid+1, right)
	return head
}
func main() {
	root := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 0},
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: -4},
			Right: &TreeNode{Val: 3},
		},
	}
	fmt.Println(serialize(root))
	a := deserialize("1.0.X.X.1.-4.X.X.3.X.X")
	fmt.Println(a)
	//fmt.Println(treeToDoublyList(root))
	//fmt.Println(isSubStructure(root, &TreeNode{
	//	Val:   1,
	//	Left:  &TreeNode{Val: -4},
	//	Right: nil,
	//}))

	//rt := &TreeNode{
	//	Val: 3,
	//	Left: &TreeNode{
	//		Val:  5,
	//		Left: nil,
	//		Right: &TreeNode{
	//			Val:   4,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//	Right: &TreeNode{
	//		Val:   1,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//}
	////path := getPath(rt, rt)
	////fmt.Println(len(path))
	////for i := 0; i < len(path); i++ {
	////	println(path[i].Val)
	////}
	//ancestor := lowestCommonAncestor(rt, rt.Left.Right, rt.Right)
	//println(ancestor.Val)

}
