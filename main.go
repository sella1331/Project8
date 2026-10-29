package main

import (
	"errors"
	"fmt"
)

type Stek interface {
	Push(item any)      //добавляет объект на вершину стека
	Pop() (any, error)  //удаляет объект с вершины стека
	Peek() (any, error) //Возвращает занчение объекта с вершины стек
}

type StekImpl struct {
	items []any
}

func NewSliceStack() *StekImpl {
	return &StekImpl{items: make([]any, 0)}
}
func (s *StekImpl) Push(item any) {
	s.items = append(s.items, item)
	fmt.Printf("Append item: %v \n", item)
}

func (s *StekImpl) Pop() (any, error) {
	if len(s.items) == 0 {
		return nil, errors.New("empty stack")
	}
	itemLastIndex := len(s.items) - 1
	item := s.items[itemLastIndex]
	s.items = s.items[:itemLastIndex]
	fmt.Printf("you remove itemLastIndex: %d, item: %v\n", itemLastIndex, item)

	return item, nil

}
func (s *StekImpl) Peek() (any, error) {
	fmt.Println("Run PEEK")
	fmt.Printf("all items: %v\n", len(s.items))
	if len(s.items) == 0 {
		return nil, errors.New("no items")
	}
	itemLastIndex := len(s.items) - 1
	item := s.items[itemLastIndex]
	fmt.Printf("you get item: %v\n", item)
	return item, nil
}

func main() {
	ctek := NewSliceStack()
	ctek.Push(1)
	ctek.Push(2)
	ctek.Push(3)
	ctek.Push(4)

	ctek.Peek()
	ctek.Pop()
	ctek.Peek()

	test := NewList()
	test.Len()
	test.Find(2)
	test.Find(3)

	_, err := test.Find(10)
	if errors.Is(err, ErrEmptyList) {
		fmt.Println("ErrEmptyList")
	}

	test.PushBack(1)
	test.PushBack(2)
	if test.Len() == 2 {
		fmt.Println("Длина списка равна 2")
	}

	countT := 2
	node, err := test.Find(countT)
	if err == nil && node.Value == countT {
		fmt.Printf(" Узел %d найден\n", countT)
	}

	err = test.Remove(node)
	if err == nil && test.Len() == countT-1 {
		fmt.Println(" Узел 2 удален, новая длина 1")
	}
	test.PushBack(3)
	fmt.Print("Текущий список: ")
	for it := test.head; it != nil; it = it.next {
		fmt.Printf("%d ", it.Value)
	}
	fmt.Println()
	err = test.Remove(node)
	if errors.Is(err, ErrNodeNotFound) {
		fmt.Println("ErrNodeNotFound")
	}

	err = test.Remove(nil)
	if errors.Is(err, ErrNilNode) {
		fmt.Println("ErrNilNode")
	}
}

//2 task

type Node struct {
	Value int
	prev  *Node
	next  *Node
}

type DoublyList struct {
	head *Node
	tail *Node
	len  int
}

func NewList() *DoublyList {
	return &DoublyList{}
}

func (d *DoublyList) Len() int {
	return d.len
}

func (d *DoublyList) PushBack(val int) {
	newNode := &Node{Value: val}

	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		newNode.prev = d.tail
		d.tail.next = newNode
		d.tail = newNode
	}
	d.len++
}

func (d *DoublyList) PushFront(val int) {
	newNode := &Node{Value: val}

	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		newNode.next = d.head
		d.head.prev = newNode
		d.head = newNode
	}
	d.len++
}

func (d *DoublyList) Remove(node *Node) error {
	if node == nil {
		return ErrNilNode
	}
	if d.len == 0 || d.head == nil {
		return ErrEmptyList
	}
	if !d.contains(node) {
		return ErrNodeNotFound
	}
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		d.head = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	} else {
		d.tail = node.prev
	}

	node.prev = nil
	node.next = nil
	d.len--

	return nil
}

func (d *DoublyList) Find(el int) (*Node, error) {
	if d.len == 0 || d.head == nil {
		return nil, ErrEmptyList
	}

	for item := d.head; item != nil; item = item.next {
		if item.Value == el {
			return item, nil
		}
	}
	return nil, ErrValueNotFound
}

func (d *DoublyList) contains(target *Node) bool {
	for item := d.head; item != nil; item = item.next {
		if item == target {
			return true
		}
	}
	return false
}
