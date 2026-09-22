/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummyHead := &ListNode{-1, nil}
    head := dummyHead

    currList1 := list1
    currList2 := list2

    for currList1 != nil && currList2 != nil {
        if currList1.Val <= currList2.Val {
            head.Next = currList1
            currList1 = currList1.Next
        } else {
            head.Next = currList2
            currList2 = currList2.Next
        }
        head = head.Next
    }

    if currList1 != nil {
        head.Next = currList1
    } else {
        head.Next = currList2
    }

    return dummyHead.Next
}
