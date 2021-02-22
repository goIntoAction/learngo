package leetcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func LeetCode1() []int {
	arr := [...]int{7, 4, 3, 2, 1, 0}
	sum := 9
	m := make(map[int]int)
	for i, v := range arr {
		another := sum - v
		if value, ok := m[another]; ok {
			return []int{i, value}
		}
		m[v] = i
	}
	return nil
}
func LeetCode2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	carry := 0
	var cur *ListNode
	for l1 != nil || l2 != nil {
		val1 := 0
		val2 := 0
		if l1 != nil {
			val1 = l1.Val
		}
		if l2 != nil {
			val2 = l2.Val
		}

		sum := val1 + val2 + carry
		if sum >= 10 {
			carry = 1
			sum = sum % 10
		} else {
			carry = 0
		}
		next := &ListNode{
			Val: sum,
		}
		if cur != nil {
			cur.Next = next
		}
		cur = next
		l1 = l1.Next
		l2 = l2.Next
	}
	return cur
}

func LeetCode3(s string) int {
	if len(s) == 0 {
		return 0
	} 


	return  0
}

func LeetCode20(str string) bool {
	if len(str) == 0 {
		return true
	}
	runes := []rune(str)
	symbol := make(map[rune]rune)
	symbol['{'] = '}'
	symbol['['] = ']'
	symbol['('] = ')'
	data := *new([]rune)
	for _, v := range runes {
		if _, ok := symbol[v]; ok {
			data = append(data, v)
		} else {
			leftSymbol := data[len(data)-1]
			data = data[:len(data)-1]
			if v != symbol[leftSymbol] {
				return false
			}
		}
	}
	return true
}

func LeetCode42() {
	arr := [...]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}

	water, leftIndex, leftMax, rightIndex, rightMax := 0, 0, 0, len(arr)-1, 0
	for leftIndex <= rightIndex {
		if arr[leftIndex] <= arr[rightIndex] {
			if arr[leftIndex] > leftMax {
				leftMax = arr[leftIndex]
			} else {
				water += leftMax - arr[leftIndex]
			}
			leftIndex++
		} else {
			if arr[rightIndex] > rightMax {
				rightMax = arr[rightIndex]
			} else {
				water += rightMax - arr[rightIndex]
			}
			rightIndex--
		}
	}
	fmt.Println(water)
}
