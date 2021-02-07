package queue

type (
	//Queue A collection designed for holding elements prior to processing
	Queue struct {
		head   *element
		tail   *element
		length int
	}
	element struct {
		value interface{}
		next  *element
	}
)

//New a new queue
func New() *Queue {
	return &Queue{nil, nil, 0}
}

//Offer Inserts the specified element into this queue.
func (queue *Queue) Offer(value interface{}) {
	newElement := &element{
		value: value,
	}
	if queue.head == nil {
		queue.head = newElement
	}
	if queue.tail != nil {
		queue.tail.next = newElement
	}
	queue.tail = newElement
	queue.length++
}

//Poll Retrieves and removes the head of this queue, or returns nil if this queue is empty.
func (queue *Queue) Poll() interface{} {
	if queue.length == 0 {
		return nil
	}
	e := queue.head
	queue.head = e.next
	queue.length--
	return e.value
}

//Peek Retrieves, but does not remove, the head of this queue, or returns nil if this queue is empty.
func (queue *Queue) Peek() interface{} {
	if queue.length == 0 {
		return nil
	}
	return queue.head.value
}

//Len Return the number of items in the stack
func (queue *Queue) Len() int {
	return queue.length
}
