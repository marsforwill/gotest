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

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
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
