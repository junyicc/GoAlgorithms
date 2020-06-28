package interviews

import (
	"container/list"
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
		node := (queue.Dequeue()).(*datastructure.TreeNode)
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

func PrintBinaryTree2(root *datastructure.TreeNode) {
	if root == nil {
		return
	}

	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		l := queue.Len()
		for i := 0; i < l; i++ {
			node := queue.Remove(queue.Front()).(*datastructure.TreeNode)
			fmt.Printf("%v\t", node.Data)

			if node.LChild != nil {
				queue.PushBack(node.LChild)
			}
			if node.RChild != nil {
				queue.PushBack(node.RChild)
			}
		}
		fmt.Println()
	}
}
