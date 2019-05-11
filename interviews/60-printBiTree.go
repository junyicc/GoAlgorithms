package interviews

import (
	"fmt"

	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
)

// PrintBinaryTree prints nodes multiline
func PrintBinaryTree(root *datastructure.TreeNode) {
	queue := datastructure.NewDynamicQueue()
	queue.Enqueue(root)
	cnt := 0
	toBePrinted := 1
	for !queue.IsEmpty() {
		node := (*queue.Dequeue()).(*datastructure.TreeNode)
		fmt.Printf("%v\t", node.Data)
		toBePrinted--

		if node.LChild != nil {
			queue.Enqueue(node.LChild)
			cnt++
		}
		if node.RChild != nil {
			queue.Enqueue(node.RChild)
			cnt++
		}

		if toBePrinted == 0 {
			fmt.Println()
			toBePrinted = cnt
			cnt = 0
		}
	}
}
