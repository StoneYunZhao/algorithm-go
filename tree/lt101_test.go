// 判断是否为对称二叉树
package tree

func help(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}

	if t1 == nil && t2 != nil {
		return false
	}

	if t1 != nil && t2 == nil {
		return false
	}

	if t1.Val != t2.Val {
		return false
	}

	return help(t1.Left, t2.Right) && help(t1.Right, t2.Left)
}
