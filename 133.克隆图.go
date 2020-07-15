/*
 * @lc app=leetcode.cn id=133 lang=golang
 *
 * [133] 克隆图
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	m := make(map[*Node]*Node, 0)
	return clone(node, m)
}

func clone(node *Node,m map[*Node]*Node)*Node {
	if node == nil {
		return nil
	}
	if v, ok := m[node]; ok {
		return v
	}
	newNode := &Node{
		Val: node.Val,
		Neighbors:make([]*Node,len(node.Neighbors)),
	}
	m[node]=newNode
    for i:=0;i<len(node.Neighbors);i++{
        newNode.Neighbors[i]=clone(node.Neighbors[i],m)
    }
    return newNode
}
// @lc code=end

