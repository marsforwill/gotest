package main

//* Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil || (root.Right == nil && root.Left == nil) {
		return true
	}
	var ans []int
	midfs(root, &ans)
	for i := 0; i < len(ans)-1; i++ {
		if ans[i] >= ans[i+1] {
			return false
		}
	}
	return true
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
func main() {
	rt := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  5,
			Left: nil,
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
	}
	//path := getPath(rt, rt)
	//fmt.Println(len(path))
	//for i := 0; i < len(path); i++ {
	//	println(path[i].Val)
	//}
	ancestor := lowestCommonAncestor(rt, rt.Left.Right, rt.Right)
	println(ancestor.Val)

}
