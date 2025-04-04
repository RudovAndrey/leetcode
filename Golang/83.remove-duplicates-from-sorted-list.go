package Golang

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	/*
	   Вводим переменную, которой присваиваем обрабатываемую структуру
	   В цикле исследуем:
	   Если значение Val = значению Vla элемента Next, атрибуту Next
	   присваиваем значение Next.Next, пропуская следующий элемент
	*/
	if head == nil {
		return nil
	}
	instance := head
	for instance.Next != nil {
		if instance.Val == instance.Next.Val {
			instance.Next = instance.Next.Next
		} else {
			instance = instance.Next
		}
	}
	return head
}
