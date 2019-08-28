package methodsinterfaces

import "fmt"

//type SLLNode
type SLLNode1 struct {
	next  *SLLNode1
	value int
}

func (sNode1 *SLLNode1) SetValue1(v int) {
	sNode1.value = v
}

func (sNode1 *SLLNode1) GetValue1() int {
	return sNode1.value
}

func NewSLLNode1() *SLLNode1 {
	return new(SLLNode1)
}

//type linked list
type SingleLinkedList1 struct {
	head *SLLNode1
	tail *SLLNode1
}

func newSingleLinkedList1() *SingleLinkedList1 {
	return new(SingleLinkedList1)
}

func (list *SingleLinkedList1) Add(v int) {
	newNode1 := &SLLNode1{value: v}
	if list.head == nil {
		list.head = newNode1
	} else if list.tail == list.head {
		list.head.next = newNode1
	} else if list.tail != nil {
		list.tail.next = newNode1
	}
	list.tail = newNode1
}

func (list *SingleLinkedList1) String() string {
	s := ""
	for n := list.head; n != nil; n = n.next {
		s += fmt.Sprintf(" {%d} ", n.GetValue1())
	}
	return s
}

func SingleLinkedList() {
	list := newSingleLinkedList1()
	list.Add(3)
	list.Add(4)
	list.Add(5)
	list.Add(6)
	fmt.Println("Hello, playground", list)
}
