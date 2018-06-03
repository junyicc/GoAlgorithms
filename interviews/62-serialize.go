package interviews

import (
	"bytes"
	"fmt"

	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
)

// Serialize binary tree
func Serialize(root *datastructure.TreeNode, buffer *bytes.Buffer) {
	if root == nil {
		buffer.WriteString(fmt.Sprint("$,"))
		return
	}
	buffer.WriteString(fmt.Sprintf("%v,", root.Data))
	Serialize(root.LChild, buffer)
	Serialize(root.RChild, buffer)
}

// Deserialize binary tree
func Deserialize(root **datastructure.TreeNode, buffer *bytes.Buffer) {
	var data datastructure.Elem
	if readStream(buffer, &data) {
		*root = new(datastructure.TreeNode)
		(*root).Data = data
		(*root).LChild = nil
		(*root).RChild = nil

		Deserialize(&((*root).LChild), buffer)
		Deserialize(&((*root).RChild), buffer)
	}
}

func readStream(buffer *bytes.Buffer, s *datastructure.Elem) bool {
	str, err := buffer.ReadString(',')
	if err != nil {
		return false
	}
	if str == "$," {
		return false
	}
	// remove ','
	*s = str[:len(str)-1]
	return true
}
