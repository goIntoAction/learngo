package main

import (
	"fmt"
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
	//caesar()
	data := [10]int{9, 3, 7, 2, 4, 1, 5, 8, 6, 10}
	// quickSortStack(0, len(data)-1, &data)
	bubbleSort(&data)
	fmt.Println(data)
}

func quickSort(startIndex, endIndex int, data *[10]int) {
	if startIndex >= endIndex || data == nil {
		return
	}
	left := startIndex
	right := endIndex
	pivot := data[startIndex]

	for left != right {
		for left < right && data[right] >= pivot {
			right--
		}
		for left < right && data[left] < pivot {
			left++
		}
		if left < right {
			data[left], data[right] = data[right], data[left]
		}
	}
	data[startIndex], data[left] = data[left], data[startIndex]
	quickSort(startIndex, left-1, data)
	quickSort(left+1, endIndex, data)
}

func quickSortStack(startIndex, endIndex int, data *[10]int) {
	if startIndex >= endIndex || data == nil {
		return
	}
	stack := stack.New()
	stack.Push(startIndex)
	stack.Push(endIndex)
	for stack.Len() > 0 {
		right := stack.Pop().(int)
		left := stack.Pop().(int)
		l := left
		r := right
		pivot := data[l]
		for left != right {
			for left < right && data[right] >= pivot {
				right--
			}
			for left < right && data[left] < pivot {
				left++
			}
			if left < right {
				data[left], data[right] = data[right], data[left]
			}
		}
		data[l], data[left] = data[left], data[l]
		if l < left-1 {
			stack.Push(l)
			stack.Push(left - 1)
		}
		if left+1 < r {
			stack.Push(left + 1)
			stack.Push(r)
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
