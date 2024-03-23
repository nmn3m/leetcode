package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var prev *ListNode
	current := slow
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	secondHalf := prev

	firstHalf := head
	for secondHalf != nil {
		if firstHalf.Val != secondHalf.Val {
			return false
		}
		firstHalf = firstHalf.Next
		secondHalf = secondHalf.Next
	}

	return true
}
