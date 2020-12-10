package _002_Add_Two_Numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func sliceToListNode(s []int) *ListNode {
	l := len(s)
	if l == 0 {
		return nil
	}

	var tmp *ListNode
	for i := l; i > 0; i-- {
		node := &ListNode{
			Val:  s[i-1],
			Next: tmp,
		}
		tmp = node
	}

	return tmp
}

func pow10(y int) int {
	result := 1
	for i := 0; i < y; i++ {
		result = result * 10
	}
	return result
}

func ListNodeToInt(l *ListNode) int {
	var result int
	for tmp, i := l, 0; tmp != nil; i++ {
		result = result + tmp.Val*pow10(i)
		tmp = tmp.Next
	}
	return result
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// ListNode 转 int
	a := ListNodeToInt(l1) + ListNodeToInt(l2)

	return nil
}
