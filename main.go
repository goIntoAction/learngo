package main

import (
	"fmt"
	"learngo/leetcode"
	"learngo/queue"
	"learngo/stack"
)

type Element struct {
	Data rune
	Pre  *Element
	Next *Element
}

var startElement *Element
var keepElement *Element

func main() {
	leetcode.LeetCode42()
}

func quickSort(startIndex, endIndex int, data *[10]int) {
	if startIndex >= endIndex || data == nil {
		return
	}
	left := startIndex
	right := endIndex
	pivot := data[startIndex]

	for left != right {
		for left < right && data[right] > pivot {
			right--
		}
		for left < right && data[left] <= pivot {
			left++
		}
		if left < right {
			data[left], data[right] = data[right], data[left]
		}
	}

	data[startIndex], data[left] = data[left], data[startIndex]

	if startIndex < left-1 {
		quickSort(startIndex, left-1, data)
	}
	if left+1 < endIndex {
		quickSort(left+1, endIndex, data)
	}
}

func quickSortByStack(data *[10]int) {
	if data == nil {
		return
	}
	startIndex := 0
	endIndex := len(data) - 1

	stack := stack.New()

	stack.Push(startIndex)
	stack.Push(endIndex)
	for stack.Len() > 0 {
		end := stack.Pop().(int)
		start := stack.Pop().(int)
		left := start
		right := end
		pivot := data[start]
		for left != right {
			for left < right && data[right] < pivot {
				right--
			}
			for left < right && data[left] >= pivot {
				left++
			}
			if left < right {
				data[left], data[right] = data[right], data[left]
			}
		}
		data[left], data[start] = data[start], data[left]

		if start < left-1 {
			stack.Push(start)
			stack.Push(left - 1)
		}

		if left+1 < end {
			stack.Push(left + 1)
			stack.Push(end)
		}
	}
}

func selectSort(data *[10]int) {
	lenght := len(data)
	for i := 0; i < lenght-1; i++ {
		maxIndex := i
		for j := i + 1; j < lenght; j++ {
			if data[j] > data[maxIndex] {
				maxIndex = j
			}
		}
		data[maxIndex], data[i] = data[i], data[maxIndex]
	}
}

func insertSort(data *[10]int) {
	length := len(data)
	for i := 1; i < length; i++ {
		temp := data[i]
		j := i - 1
		for ; j >= 0 && data[j] > temp; j-- {
			data[j+1] = data[j]
		}
		data[j+1] = temp
	}
}

func bubbleSort(data *[10]int) {
	length := len(data)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1-i; j++ {
			if data[j+1] < data[j] {
				data[j+1], data[j] = data[j], data[j+1]
			}
		}
	}
}

func caesar() {
	var startIndex int
	fmt.Scan(&startIndex)
	aElement := getElement()
	startElement = aElement
	if startIndex > 0 {
		for i := 0; i < startIndex; i++ {
			startElement = startElement.Next
		}
	} else {
		for i := 0; i > startIndex; i-- {
			startElement = startElement.Pre
		}
	}
	keepElement = startElement
	for {
		fmt.Printf("%s", string(startElement.Data))
		if startElement.Next != nil && startElement.Next != keepElement {
			startElement = startElement.Next
		} else {
			break
		}
	}
}

func getElement() (start *Element) {
	var pre *Element
	var latest *Element
	for i := 0; i < 26; i++ {
		element := new(Element)
		if i == 0 {
			start = element
		}
		element.Data = rune(i + 'A')
		element.Pre = pre
		if pre != nil {
			pre.Next = element
		}
		pre = element
		latest = element
	}
	if start != nil {
		start.Pre = latest
		latest.Next = start
	}
	return
}

// func isValid(s string) bool {

// }

func testQueue() {
	queue := queue.New()
	queue.Offer(1)
	queue.Offer(2)
	queue.Offer(3)
	queue.Offer(4)
	queue.Offer(5)
	for queue.Len() > 0 {
		value := queue.Peek()
		fmt.Println(value)
	}
}
