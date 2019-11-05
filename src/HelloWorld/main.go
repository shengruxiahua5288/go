package main

// HelloWorld project main.go
// package main

import (
	"fmt"
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// func main() {
// 	nums1 := []int{-3, -1, 1, 1, 4}
// 	nums2 := []int{-3, 1, 4}
// 	nums3 := []int{-2, -1, 0, 2}
// 	pos := -1

// 	// var numbers [][]int = threeSum(nums)
// 	// selfPrintSecond(numbers)
// 	//var a int = firstMissingPositive(nums)
// 	// 创建单链表
// 	var q1 *ListNode = createListNode(nums1, pos)
// 	var q2 *ListNode = createListNode(nums2, pos)
// 	var q3 *ListNode = createListNode(nums3, pos)
// 	var lists []*ListNode = []*ListNode{q1, q3}
// 	printListNode(q1)
// 	println()
// 	printListNode(q2)
// 	println()
// 	printListNode(q3)
// 	println()

// 	//printListNode(mergeLink(q1, q2))
// 	// 判断链表是否有环
// 	//print(hasCycle(head))
// 	printListNode(mergeKLists(lists))

// }

// 创建链表
func createListNode(nums []int, pos int) *ListNode {

	var head *ListNode = new(ListNode)
	var q *ListNode
	var w *ListNode
	head.Val = nums[0]
	head.Next = nil
	q = head
	if pos == 0 && len(nums) == 1 {
		head.Next = head
		return head
	}
	for i := 1; i < len(nums); i++ {
		var p *ListNode = new(ListNode)
		p.Val = nums[i]
		p.Next = nil
		if pos == i {
			w = p
		}
		q.Next = p
		q = p

	}
	q.Next = w
	return head
}

// 输出链表
func printListNode(head *ListNode) {
	var temp *ListNode = head
	print(head.Val)
	for temp.Next != nil {
		temp = temp.Next
		print(temp.Val)
	}
}

func selfPrint(nums []int) {
	for a := 0; a < len(nums); a++ {
		fmt.Println(nums[a])
	}
}

func selfPrintSecond(numbers [][]int) {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers[i]); j++ {
			fmt.Print(numbers[i][j])
		}
		fmt.Println()
	}
}
func jieCheng(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * jieCheng(n-1)
	}
}
func initNumbers(size int) [][]int {
	n := jieCheng(size) * jieCheng(3) / jieCheng(size-3)
	var numbers [][]int
	for i := 0; i < n; i++ {
		temp := make([]int, 3)
		numbers = append(numbers, temp)
	}
	return numbers
}

// leetcode 三叔之和
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	size := len(nums)
	var numbers [][]int
	j := 0
	for a := 0; a < size; a++ {
		if nums[a] > 0 {
			break
		}
		// 去重
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		target := nums[a]
		l := a + 1
		r := size - 1
		for l < r {
			for target+nums[l]+nums[r] < 0 && l < r {

				l++
			}
			for target+nums[l]+nums[r] > 0 && l < r {
				r--
			}
			for target+nums[l]+nums[r] == 0 && l < r {

				if !(nums[l] == nums[l-1] && r < size-1 && nums[r] == nums[r+1]) {

					temp := make([]int, 3)
					numbers = append(numbers, temp)
					numbers[j][0] = target
					numbers[j][1] = nums[l]
					numbers[j][2] = nums[r]
					j++
				}

				l++
				r--
			}
		}

	}
	return numbers
}

// 求纵数
// func majorityElement(nums []int) int {
// 	sort.Ints(nums)
// 	selfPrint(nums)
// 	max := 1
// 	temp := 1
// 	var result int
// 	if len(nums) == 1 {
// 		return nums[0]
// 	}

// 	for i := 1; i < len(nums); i++ {

// 		if nums[i] == nums[i-1] {
// 			temp++
// 			if i == len(nums)-1 {
// 				max = temp
// 				result = nums[i-1]
// 			}

// 		} else {

// 			if temp > max {

// 				max = temp
// 			}
// 			temp = 1
// 			if max > len(nums)/2 {

// 				result = nums[i-1]
// 				break
// 			}
// 		}

// 	}

// 	return result
// }

// 别人的方法

func majorityElement(nums []int) int {
	sort.Ints(nums)

	return nums[len(nums)/2]
}

// 缺失的第一个正数
func firstMissingPositive(nums []int) int {
	sort.Ints(nums)
	var first int = 1
	var result int = 1

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {

			if nums[i] > first {
				result = first
				return result
			} else {
				if i+1 < len(nums) && nums[i+1] != nums[i] {
					first = first + 1
				}
				if i == len(nums)-1 {
					first = first + 1
					result = first
				} else {
					first = first
				}
			}
		}
	}
	return result
}

//环形链表
func hasCycle(head *ListNode) bool {
	var temp *ListNode = head
	print(temp)
	var countryCapitalMap map[*ListNode]int /*创建集合*/
	countryCapitalMap = make(map[*ListNode]int)

	for temp != nil {

		if _, ok := countryCapitalMap[temp]; ok {
			// key 存在
			return ok
		} else {
			//key 不存在
			countryCapitalMap[temp] = temp.Val
		}
		temp = temp.Next
	}
	return false

}

//合并两个有序链表
func mergeLink(l1 *ListNode, l2 *ListNode) *ListNode {
	var p1 *ListNode = l1
	var p2 *ListNode = l2
	var q1 *ListNode
	var q2 *ListNode
	var head *ListNode

	var flag bool = false
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 != nil && l2 == nil {
		return l1
	}
	if l2 != nil && l1 == nil {
		return l2
	}
	if l1.Val >= l2.Val {
		head = l2
	} else {
		head = l1
	}
	for p1 != nil && p2 != nil {

		for p1 != nil && p2 != nil && p1.Val >= p2.Val {

			q2 = p2
			p2 = p2.Next
			flag = true
		}
		if flag {
			q2.Next = p1
			p1 = p1.Next
			if (p2 != nil && p1 != nil && p2.Val <= p1.Val) || (p1 == nil && p2 != nil) {
				q2.Next.Next = p2
			}

			flag = false
		}

		for p1 != nil && p2 != nil && p1.Val < p2.Val {
			q1 = p1
			p1 = p1.Next
			flag = true
		}
		if flag {
			q1.Next = p2
			p2 = p2.Next
			if (p1 != nil && p2 != nil && p1.Val < p2.Val) || (p2 == nil && p1 != nil) {
				q1.Next.Next = p1
			}

			flag = false
		}

	}
	printListNode(head)
	println()
	return head

}

func mergeKLists(lists []*ListNode) *ListNode {

	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	var l1 *ListNode = lists[0]
	for i := 1; i < len(lists); i++ {
		printListNode(l1)
		println()
		l1 = mergeLink(l1, lists[i])

		printListNode(l1)
		println()

	}
	return l1
}
