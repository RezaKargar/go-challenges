package main

import "fmt"

type Element struct {
  Value string
  Next  *Element
}

type SinglyLinkedList struct {
  Count int
  Head  *Element
  Tail  *Element
}

func (l *SinglyLinkedList) Append(value string) {
	newNode := Element {
		Value: value,
		Next: nil,
	}

	if l.Head == nil {
		l.Head = &newNode
		l.Tail = &newNode
	} else {
		current := l.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = &newNode
		l.Tail = &newNode
	}

	l.Count++
}

// You will have to ensure when you add new elements
// that this method still returns the correct value
func (l *SinglyLinkedList) Size() int {
  return l.Count
}

func (l *SinglyLinkedList) Print() {
  current := l.Head
  fmt.Printf("%+v\n", current.Value)
  for current.Next != nil {
    fmt.Printf("%+v\n", current.Value)
    current = current.Next
  }
}

func main() {
  fmt.Println("Singly Linked List Challenge")

  var llist SinglyLinkedList

  values := []string{"First", "Second", "Third"}
  for _, value := range values {
    llist.Append(value)
  }
  llist.Print()
}
