package interviews

import (
	"fmt"

	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
)

// PrintBiTreeInZ prints tree nodes in z
func PrintBiTreeInZ(root *datastructure.TreeNode) {
	if root == nil {
		return
	}
	stack0 := new(datastructure.DynamicStack)
	stack1 := new(datastructure.DynamicStack)
	stack0.Push(root)
	flag := 0
	for !stack0.IsEmpty() || !stack1.IsEmpty() {
		if flag == 0 {
			for !stack0.IsEmpty() {
				node := (*stack0.Pop()).(*datastructure.TreeNode)
				fmt.Printf("%v\t", node.Data)
				if node.LChild != nil {
					stack1.Push(node.LChild)
				}
				if node.RChild != nil {
					stack1.Push(node.RChild)
				}
			}
		} else {
			for !stack1.IsEmpty() {
				node := (*stack1.Pop()).(*datastructure.TreeNode)
				fmt.Printf("%v\t", node.Data)
				if node.RChild != nil {
					stack0.Push(node.RChild)
				}
				if node.LChild != nil {
					stack0.Push(node.LChild)
				}
			}

		}
		fmt.Println()
		flag = 1 - flag
	}
}
