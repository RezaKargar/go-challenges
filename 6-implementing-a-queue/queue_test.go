package main

import (
	"testing"
)

func TestConstruction(t *testing.T) {
	//Arrange
	queue := Queue{}

	//Act

	//Assert
	if len(queue.items) != 0 {
		t.Errorf("queue.items is not empty")
	}
}

func TestQueue_Push_WhenQueueIsEmpty(t *testing.T) {
	//Arrange
	queue := Queue{}
	item := struct{}{}

	//Act
	queue.Push(item)

	//Assert
	if len(queue.items) != 1 {
		t.Errorf("Expected queue length of 1, but got %d", len(queue.items))
	}

	if queue.items[0] != item {
		t.Errorf("Expected %v, but got %v", item, queue.items[0])
	}
}

func TestQueue_Push_WhenQueueIsNotEmpty(t *testing.T) {
	//Arrange
	queue := Queue{
		items: []interface{}{nil},
	}

	//Act
	item := struct{}{}
	queue.Push(item)

	//Assert
	if len(queue.items) != 2 {
		t.Errorf("Expected queue length of 1, but got %d", len(queue.items))
	}

	if queue.items[1] != item {
		t.Errorf("Expected %v, but got %v", item, queue.items[1])
	}
}

func TestQueue_Peek_WhenQueueIsEmpty(t *testing.T) {
	//Arrange
	queue := Queue{}

	//Act
	peekedItem, err := queue.Peek()

	//Assert
	if err == nil {
		t.Errorf("Expected error, but got nil [Peek on empty queue]")
	}

	if peekedItem != nil {
		t.Errorf("Expected nil, but got %v [Peek on empty queue]", peekedItem)
	}
}

func TestQueue_Peek_WhenQueueIsNotEmpty(t *testing.T) {
	//Arrange
	item := struct{}{}
	queue := Queue{
		items: []interface{}{item, nil},
	}

	//Act
	peekedItem, err := queue.Peek()

	//Assert
	if err != nil {
		t.Errorf("Peek returned an error [Peek on non-empty queue]")
	}

	if peekedItem != item {
		t.Errorf("Expected %v, but got %v", item, peekedItem)
	}
}

func TestQueue_Pop_WhenQueueIsEmpty(t *testing.T) {
	//Arrange
	queue := Queue{}

	//Act
	poppedItem, err := queue.Pop()

	//Assert
	if err == nil {
		t.Errorf("Expected error, but got nil [Pop on empty queue]")
	}

	if poppedItem != nil {
		t.Errorf("Expected nil, but got %v [Pop on empty queue]", poppedItem)
	}
}

func TestQueue_Pop_WhenQueueIsNotEmpty(t *testing.T) {
	//Arrange
	item := struct{}{}
	queue := Queue{
		items: []interface{}{item, nil},
	}

	//Act
	poppedItem, err := queue.Pop()

	//Assert
	if err != nil {
		t.Errorf("Pop returned an error [Pop on non-empty queue]")
	}

	if poppedItem != item {
		t.Errorf("Expected %v, but got %v", item, poppedItem)
	}

	if len(queue.items) != 1 {
		t.Errorf("Expected queue length of 1, but got %d", len(queue.items))
	}
}

func TestQueue_IsEmpty_WhenQueueIsEmpty(t *testing.T) {
	//Arrange
	queue := Queue{}

	//Act

	//Assert
	if !queue.IsEmpty() {
		t.Errorf("Expected to detect queue is empty")
	}
}

func TestQueue_IsEmpty_WhenQueueIsNotEmpty(t *testing.T) {
	//Arrange
	queue := Queue{
		items: []interface{}{struct{}{}},
	}

	//Act

	//Assert
	if queue.IsEmpty() {
		t.Errorf("Expected to detect queue is not empty")
	}
}
