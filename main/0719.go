package main

import "fmt"

func numWaterBottles(numBottles int, numExchange int) int {
	numEmpty := 0
	ans := 0
	for numBottles > 0 || numEmpty > numExchange {
		ans += numBottles
		numEmpty += numBottles
		extrange := numEmpty / numExchange
		numBottles = extrange
		numEmpty -= numExchange * extrange
	}
	return ans
}

// 输入：n = 7, edges = [[0,1],[0,2],[1,4],[1,5],[2,3],[2,6]], labels = "abaedcd"
//输出：[2,1,1,1,1,1,1]
func countSubTrees(n int, edges [][]int, labels string) []int {
	relation := genRelation(edges)
	ans := make([]int, n)
	dfsSubTree(0, relation, labels, -1, &ans)
	return ans
}

// 1. 这棵树的信息需要自底向上的处理
// 2. 信息能够传递/处理/合并的形式 [a b c... x y z] count[0 0 0...0 0 0]
func dfsSubTree(i int, relation [][]int, labels string, visited int, ans *[]int) []int {
	count := make([]int, 26)
	// 当前节点计数+1 count 存下标对应的字母计数
	count[labels[i]-'a']++
	//是按照节点顺序 0 1 2 3。。。 来循环，最后先dfs的是叶子节点后续遍历
	for _, c := range relation[i] {
		// visited是当前节点i的父亲节点 不需要处理
		if c == visited {
			continue
		}
		// i是c的父亲节点
		re := dfsSubTree(c, relation, labels, i, ans)
		// 合并子树的统计结果
		for i := 0; i < 26; i++ {
			count[i] += re[i]
		}
	}
	(*ans)[i] = count[labels[i]-'a']
	return count
}

func genRelation(edges [][]int) [][]int {
	relation := make([][]int, len(edges)+1)
	for _, e := range edges {
		relation[e[0]] = append(relation[e[0]], e[1])
		relation[e[1]] = append(relation[e[1]], e[0])
	}
	return relation
}

//* Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//好叶子节点对 输入：root = [1,2,3,null,4], distance = 3
//输出：1
func countPairs(root *TreeNode, distance int) int {
	ans := 0
	dfsPair(root, distance, &ans)
	return ans
}

func dfsPair(node *TreeNode, distance int, ans *int) []int {
	if node == nil {
		return []int{}
	}
	if node.Left == nil && node.Right == nil {
		return []int{0}
	}
	//当前node下所有叶子节点的距离
	var num []int
	l := dfsPair(node.Left, distance, ans)
	for i := 0; i < len(l); i++ {
		if l[i]+1 <= distance {
			num = append(num, l[i]+1)
		}
	}
	r := dfsPair(node.Right, distance, ans)
	for i := 0; i < len(r); i++ {
		if r[i]+1 <= distance {
			num = append(num, r[i]+1)
		}
	}
	// 计算过当前节点的左右子树符合条件的节点对
	for i := 0; i < len(l); i++ {
		for j := 0; j < len(r); j++ {
			if l[i]+1+r[j]+1 <= distance {
				*ans++
			}
		}
	}
	return num
}

func main() {
	//ans := countSubTrees(7, [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}, "abaedcd")
	//fmt.Println(ans)
	ans := countPairs(&TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   4,
			Left:  nil,
			Right: nil,
		},
	}, 3)
	fmt.Println(ans)
}
