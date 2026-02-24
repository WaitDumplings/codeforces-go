package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

// https://space.bilibili.com/206214
func sumRootToLeaf1(root *TreeNode) (ans int) {
	// 从根到 node（不含）的路径值为 num
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, num int) {
		if node == nil {
			return
		}
		num = num<<1 | node.Val
		if node.Left == nil && node.Right == nil {
			ans += num
			return
		}
		dfs(node.Left, num)
		dfs(node.Right, num)
	}
	dfs(root, 0)
	return
}

func dfs(node *TreeNode, num int) int {
	if node == nil {
		return 0
	}
	num = num<<1 | node.Val
	if node.Left == nil && node.Right == nil {
		return num
	}
	return dfs(node.Left, num) + dfs(node.Right, num)
}

func sumRootToLeaf(root *TreeNode) int {
	return dfs(root, 0)
}
