package main

import (
	"testing"
)

func TestConstruction(t *testing.T) {
	//Arrange
	stack := Stack{}

	//Act

	//Assert
	if len(stack.items) != 0 {
		t.Errorf("stack.items is not empty")
	}
}

func TestStack_Push_WhenStackIsEmpty(t *testing.T) {
	//Arrange
	stack := Stack{}
	item := struct{}{}

	//Act
	stack.Push(item)

	//Assert
	if len(stack.items) != 1 {
		t.Errorf("Expected stack length of 1, but got %d", len(stack.items))
	}

	if stack.items[0] != item {
		t.Errorf("Expected %v, but got %v", item, stack.items[0])
	}
}

func TestStack_Push_WhenStackIsNotEmpty(t *testing.T) {
	//Arrange
	stack := Stack{
		items: []interface{}{nil},
	}

	//Act
	item := struct{}{}
	stack.Push(item)

	//Assert
	if len(stack.items) != 2 {
		t.Errorf("Expected stack length of 1, but got %d", len(stack.items))
	}

	if stack.items[1] != item {
		t.Errorf("Expected %v, but got %v", item, stack.items[1])
	}
}

func TestStack_Peek_WhenStackIsEmpty(t *testing.T) {
	//Arrange
	stack := Stack{}

	//Act
	peekedItem, err := stack.Peek()

	//Assert
	if err == nil {
		t.Errorf("Expected error, but got nil [Peek on empty stack]")
	}

	if peekedItem != nil {
		t.Errorf("Expected nil, but got %v [Peek on empty stack]", peekedItem)
	}
}

func TestStack_Peek_WhenStackIsNotEmpty(t *testing.T) {
	//Arrange
	item := struct{}{}
	stack := Stack{
		items: []interface{}{nil, item},
	}

	//Act
	peekedItem, err := stack.Peek()

	//Assert
	if err != nil {
		t.Errorf("Peek returned an error [Peek on non-empty stack]")
	}

	if peekedItem != item {
		t.Errorf("Expected %v, but got %v", item, peekedItem)
	}
}

func TestStack_Pop_WhenStackIsEmpty(t *testing.T) {
	//Arrange
	stack := Stack{}

	//Act
	poppedItem, err := stack.Pop()

	//Assert
	if err == nil {
		t.Errorf("Expected error, but got nil [Pop on empty stack]")
	}

	if poppedItem != nil {
		t.Errorf("Expected nil, but got %v [Pop on empty stack]", poppedItem)
	}
}

func TestStack_Pop_WhenStackIsNotEmpty(t *testing.T) {
	//Arrange
	item := struct{}{}
	stack := Stack{
		items: []interface{}{nil, item},
	}

	//Act
	poppedItem, err := stack.Pop()

	//Assert
	if err != nil {
		t.Errorf("Pop returned an error [Pop on non-empty stack]")
	}

	if poppedItem != item {
		t.Errorf("Expected %v, but got %v", item, poppedItem)
	}

	if len(stack.items) != 1 {
		t.Errorf("Expected stack length of 1, but got %d", len(stack.items))
	}
}

func TestStack_IsEmpty_WhenStackIsEmpty(t *testing.T) {
	//Arrange
	stack := Stack{}

	//Act

	//Assert
	if !stack.IsEmpty() {
		t.Errorf("Expected to detect stack is empty")
	}
}

func TestStack_IsEmpty_WhenStackIsNotEmpty(t *testing.T) {
	//Arrange
	stack := Stack{
		items: []interface{}{struct{}{}},
	}

	//Act

	//Assert
	if stack.IsEmpty() {
		t.Errorf("Expected to detect stack is not empty")
	}
}
