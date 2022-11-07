package main

type TreeNode struct {
	Val   int //节点
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func sortedArrayToBST(nums []int) *TreeNode {
	return help(nums, 0, len(nums)-1)
}
func help(nums []int, left int, right int) *TreeNode {

	if left > right {
		return nil
	}
	mid := (left + right) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = help(nums, left, mid-1)
	root.Right = help(nums, mid+1, right)
	return root
}
