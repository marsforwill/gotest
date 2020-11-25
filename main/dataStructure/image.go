package main

import "fmt"

/**
 * Definition for a Node.
 */
type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	fg := make(map[int]*Node)
	var dfsImage func(node *Node) *Node
	dfsImage = func(n *Node) *Node {
		if n == nil {
			return nil
		}
		nodedeal, ok := fg[n.Val]
		if ok {
			return nodedeal
		}
		ans := Node{
			Val:       n.Val,
			Neighbors: make([]*Node, len(n.Neighbors)),
		}
		fg[n.Val] = &ans
		for i, temp := range n.Neighbors {
			ans.Neighbors[i] = dfsImage(temp)
		}
		return &ans
	}
	return dfsImage(node)
}

//func dfsImage(node *Node, fg *map[int]Node) *Node {
//	if node == nil {
//		return nil
//	}
//	nodedeal,ok := (*fg)[node.Val]
//	if ok {
//		return &nodedeal
//	}
//	ans := Node{
//		Val:       node.Val,
//		Neighbors: make([]*Node,len(node.Neighbors)),
//	}
//	(*fg)[ans.Val] = ans
//	for i,n := range node.Neighbors{
//		ans.Neighbors[i] = dfsImage(n, fg)
//	}
//	return &ans
//}
func main() {
	n1 := &Node{
		Val: 1,
	}
	n2 := &Node{
		Val: 2,
	}
	n3 := &Node{
		Val: 3,
	}
	n4 := &Node{
		Val: 4,
	}
	n1.Neighbors = append(append(n1.Neighbors, n2), n4)
	n3.Neighbors = append(append(n3.Neighbors, n2), n4)
	n2.Neighbors = append(append(n2.Neighbors, n1), n3)
	n4.Neighbors = append(append(n4.Neighbors, n1), n3)
	graph := cloneGraph(n1)
	fmt.Println(*graph)
}
