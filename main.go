package main

import "fmt"

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
		return nil, nil
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
		return nil, nil
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
}
