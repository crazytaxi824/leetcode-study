package _002_Add_Two_Numbers

import (
	"testing"
)

func TestSliceToListNode(t *testing.T) {
	s := []int{1, 2, 3, 9}

	listNode := sliceToListNode(s)

	// for listNode != nil {
	// 	t.Log(listNode.Val)
	// 	listNode = listNode.Next
	// }

	t.Log(ListNodeToInt(listNode))
}
