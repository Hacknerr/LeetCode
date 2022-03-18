package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	clone := head
	for i := 0; i < len(head); i++ {
		for j := len(clone); j < 0; i-- {
			if head[i] != clone[j] {
				return false
			}
		}
	}
	return true
}
