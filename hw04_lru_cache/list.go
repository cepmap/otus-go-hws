package hw04lrucache

import "fmt"

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
	PrintList()
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len       int
	firstNode *ListItem
	lastNode  *ListItem
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.firstNode
}

func (l *list) Back() *ListItem {
	return l.lastNode
}

func (l *list) PushFront(v interface{}) *ListItem {
	l.firstNode = &ListItem{Value: v, Next: l.firstNode, Prev: nil}
	if l.len == 0 {
		l.lastNode = l.firstNode
	} else {
		l.firstNode.Next.Prev = l.firstNode
	}

	l.len++
	return l.firstNode
}
func (l *list) PushBack(v interface{}) *ListItem {
	l.lastNode = &ListItem{Value: v, Next: nil, Prev: l.lastNode}
	if l.len == 0 {
		l.firstNode = l.lastNode
	} else {
		l.lastNode.Prev.Next = l.lastNode
	}

	l.len++
	return l.lastNode
}

func (l *list) Remove(i *ListItem) {
	if i.Prev == nil {
		l.firstNode = i.Next
	} else {
		i.Prev.Next = i.Next
	}
	if i.Next == nil {
		l.lastNode = i.Prev
	} else {
		i.Next.Prev = i.Prev
	}
	i.Next, i.Prev = nil, nil
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	switch {
	case i.Prev == nil:
		return
	case i.Next == nil:
		l.lastNode = i.Prev
	default:
		i.Next.Prev = i.Prev
	}
	l.firstNode.Prev, l.firstNode, i.Next, i.Prev.Next = i, i, l.firstNode, i.Next
}

func (l *list) PrintList() {
	node := l.firstNode
	for node != nil {
		fmt.Print(node.Value)
		node = node.Next
	}
}

func NewList() List {
	return new(list)
}
