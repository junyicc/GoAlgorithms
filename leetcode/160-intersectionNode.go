package leetcode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	nodeA, nodeB := headA, headB
	cntA, cntB := 0, 0
	for nodeA != nil {
		cntA++
		nodeA = nodeA.Next
	}
	for nodeB != nil {
		cntB++
		nodeB = nodeB.Next
	}
	nodeA = headA
	nodeB = headB
	if cntA > cntB {
		d := cntA - cntB
		for i := 0; i < d; i++ {
			nodeA = nodeA.Next
		}
	} else {
		d := cntB - cntA
		for i := 0; i < d; i++ {
			nodeB = nodeB.Next
		}
	}

	for nodeA != nil && nodeB != nil && nodeA != nodeB {
		nodeA = nodeA.Next
		nodeB = nodeB.Next
	}
	return nodeA
}

func getIntersectionNodeWithLoop(headA, headB *ListNode) *ListNode {
	// find the entry of loop in list A
	// find the entry of loop in list B

	// if entryA == entryB {
	// 	the intersection node is not in the loop
	// 	use the entry as the last node
	// 	find the intersection node between head and entry
	// } else {
	// 	the intersection node is in the loop
	// }

	return nil
}
