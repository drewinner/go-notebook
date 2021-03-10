package 常用题目

/**
*	https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof/
*	输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
*	前序遍历 preorder = [3,9,20,15,7]
*	中序遍历 inorder = [9,3,15,20,7]
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//06_01.jpg
// 前序遍历：preorder 10,6,4,8,14,12,16 第一个数字是根节点
// 中序遍历：inorder 4,6,8,10,12,14,16 根节点在中间、左子树的节点在根节点左面、右子树在根节点的右边
// 后续遍历：4,8,6,12,16,14,10
func buildTree(preorder []int, inorder []int) *TreeNode {
	//退出条件
	if len(preorder) == 0 {
		return nil
	}
	//定义根节点
	root := &TreeNode{preorder[0], nil, nil}
	var i = 0
	//根据中序遍历结果、找到根节点位置
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])

	return root

}
