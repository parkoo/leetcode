// 初等数学，两数相加
// 时间复杂度：O(max(m,n))  空间复杂度：O(max(m,n))

/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */

public class Solution_1 {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {

        ListNode dummyHead = new ListNode(-1);  // 虚拟头结点
        ListNode cur = dummyHead;

        int carry = 0;  // 进位
        while(l1!=null && l2!=null) {
            int temp = l1.val + l2.val + carry;
            carry = temp/10;
            cur.next = new ListNode(temp%10);
            cur = cur.next;
            l1 = l1.next;
            l2 = l2.next;
        }

        while(l1!=null) {
            int temp = l1.val + carry;
            carry = temp/10;
            cur.next = new ListNode(temp%10);
            cur = cur.next;
            l1 = l1.next;
        }

        while(l2!=null) {
            int temp = l2.val + carry;
            carry = temp/10;
            cur.next = new ListNode(temp%10);
            cur = cur.next;
            l2 = l2.next;
        }

        if(carry!=0) {
            cur.next = new ListNode(carry);
            cur = cur.next;
        }

        return dummyHead.next;
    }
}