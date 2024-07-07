package main

import (
	"errors"
)

type Queue struct {
	items []interface{}
}

func (q *Queue) Pop() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is Empty")
	} else {
		item := q.items[0]
		q.items = q.items[1:]
		return item, nil
	}
}

func (q *Queue) Push(item interface{}) {
	q.items = append(q.items, item)
}

func (q *Queue) Peek() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is Empty")
	} else {
		return q.items[0], nil
	}
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}
