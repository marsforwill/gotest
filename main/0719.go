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

func main() {
	ans := countSubTrees(7, [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}, "abaedcd")
	fmt.Println(ans)
}
