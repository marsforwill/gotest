package main

import "fmt"

/**
 * Definition for a GNode.
 */
type GNode struct {
	Val       int
	Neighbors []*GNode
}

func cloneGraph(node *GNode) *GNode {
	fg := make(map[int]*GNode)
	var dfsImage func(node *GNode) *GNode
	dfsImage = func(n *GNode) *GNode {
		if n == nil {
			return nil
		}
		nodedeal, ok := fg[n.Val]
		if ok {
			return nodedeal
		}
		ans := GNode{
			Val:       n.Val,
			Neighbors: make([]*GNode, len(n.Neighbors)),
		}
		fg[n.Val] = &ans
		for i, temp := range n.Neighbors {
			ans.Neighbors[i] = dfsImage(temp)
		}
		return &ans
	}
	return dfsImage(node)
}

//func dfsImage(node *GNode, fg *map[int]GNode) *GNode {
//	if node == nil {
//		return nil
//	}
//	nodedeal,ok := (*fg)[node.Val]
//	if ok {
//		return &nodedeal
//	}
//	ans := GNode{
//		Val:       node.Val,
//		Neighbors: make([]*GNode,len(node.Neighbors)),
//	}
//	(*fg)[ans.Val] = ans
//	for i,n := range node.Neighbors{
//		ans.Neighbors[i] = dfsImage(n, fg)
//	}
//	return &ans
//}
func main() {
	n1 := &GNode{
		Val: 1,
	}
	n2 := &GNode{
		Val: 2,
	}
	n3 := &GNode{
		Val: 3,
	}
	n4 := &GNode{
		Val: 4,
	}
	n1.Neighbors = append(append(n1.Neighbors, n2), n4)
	n3.Neighbors = append(append(n3.Neighbors, n2), n4)
	n2.Neighbors = append(append(n2.Neighbors, n1), n3)
	n4.Neighbors = append(append(n4.Neighbors, n1), n3)
	graph := cloneGraph(n1)
	fmt.Println(*graph)
}
