package queue

import "testing"

func testOffer(t *testing.T) {
	queue := New()
	queue.Offer(1)
	queue.Offer(2)
	lenght := queue.Len()
	if lenght != 2 {
		t.Error("lenght is wrong")
	}
	if queue.Poll().(int) != 2 {
		t.Error("value is wrong")
	}
	if queue.Poll().(int) != 1 {
		t.Error("value is wrong")
	}
}
