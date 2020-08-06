package interviews

import "github.com/CAIJUNYI/GoAlgorithms/datastructure"

// 给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。
//
// 返回删除后的链表的头节点。
//
// 注意：此题对比原题有改动
//
// 示例 1:
// 输入: head = [4,5,1,9], val = 5
// 输出: [4,1,9]
// 解释: 给定你链表中值为5的第二个节点，那么在调用了你的函数之后，该链表应变为 4 -> 1 -> 9.
//
// 示例 2:
// 输入: head = [4,5,1,9], val = 1
// 输出: [4,5,9]
// 解释: 给定你链表中值为1的第三个节点，那么在调用了你的函数之后，该链表应变为 4 -> 5 -> 9.

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	sentinel := new(ListNode)
	sentinel.Next = head

	node := sentinel
	for node.Next != nil && node.Next.Val != val {
		node = node.Next
	}
	// not found val
	if node.Next == nil {
		return sentinel.Next
	}

	if node.Next.Next == nil {
		// tail node
		node.Next = nil
	} else {
		node.Next.Val = node.Next.Next.Val
		// delete node.Next.Next
		node.Next.Next = node.Next.Next.Next
	}

	return sentinel.Next
}

// DeleteNode deletes the toBeDel node in the linkedlist in O(1)
func DeleteNode(head **datastructure.ListNode, toBeDel *datastructure.ListNode) *datastructure.Elem {
	if head == nil || toBeDel == nil {
		return nil
	}
	e := toBeDel.Data
	if toBeDel.Next != nil {
		// O(1)
		next := toBeDel.Next
		toBeDel.Data = next.Data
		toBeDel.Next = next.Next
		next.Next = nil
		next = nil
	} else if *head == toBeDel {
		// only one node in linkedlist, delete head(tail) node
		*head = nil
	} else {
		// delete tail node
		node := *head
		// O(n)
		for node.Next != toBeDel {
			node = node.Next
		}
		node.Next = nil
	}
	return &e
}
