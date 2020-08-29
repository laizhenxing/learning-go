package caculate

//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
//
//如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//
//您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
//示例：
//
//输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers1(l1, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil  {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	sum := l1.Val + l2.Val
	next := addTwoNumbers1(l1.Next, l2.Next)

	if sum < 10 {
		return &ListNode{
			Val: sum,
			Next: next,
		}
	} else {
		tmp := &ListNode{
			Val: 1,
			Next: nil,
		}
		return &ListNode{
			Val: sum - 10,
			Next: addTwoNumbers1(next, tmp),
		}
	}
}

func addTowNumbers2(l1, l2 *ListNode) *ListNode {
	l := new(ListNode)
	cur := l
	sum := 0

	for sum > 0 || l1 != nil || l2 != nil {
		cur.Next = new(ListNode)
		cur = cur

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2.Next = l2
		}
		cur.Val = sum % 10
		sum /= 10
	}

	return cur.Next
}
