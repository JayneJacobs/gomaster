package methodsinterfaces

import (
	"errors"
	"fmt"
)

var ErrInvalidNode = errors.New("Node is not valid")

//Node interface declares methods
type Node interface {
	SetValue(v int) error
	GetValue() int
}

// SLLNode Implements the Node Interface
type SLLNode struct {
	next          *SLLNode
	value         int
	SNodeMesssage string
}

// SetValue *SSLNode Method
func (sNode *SLLNode) SetValue(v int) error {
	if sNode == nil {
		return ErrInvalidNode
	}
	sNode.value = v
	return nil
}

// GetValue *SSLNode Method
func (sNode *SLLNode) GetValue() int {
	return sNode.value
}

// NewSLLNode creates an instance of SSLNode and returns a pointer
func NewSLLNode() *SLLNode {
	return &SSLNode{SNodeMessage: "This is a message from the normal Node"}
}

// PowerNode Type
type PowerNode struct {
	next         *PowerNode
	value        int
	PNodeMessage string
}

// SetValue method
func (sNode *PowerNode) SetValue(v int) error {
	if sNode == nil {
		return ErrInvalidNode
	}
	sNode.value = v * 10
	return nil
}

// GetValue returns sNode as an int
func (sNode *PowerNode) GetValue() int {
	return sNode.value
}

// NewPowerNode Constructor
func NewPowerNode() *PowerNode {
	return &PowerNode{PNodeMessage: "this is a message from the power Node"}
}

// GetSLLNode Gets teh values of SSLNode
func GetSLLNode() {
	var node Node
	node = NewSLLNode() //NewSSLNode can be used because SSLNode Implements Node Interface
	node.SetValue(4)
	fmt.Println("Node is of value ", node.GetValue())
	if i, ok := node.(*SLLNode); ok {
		fmt.Println("This is SLLNode", i.value)
	}

	node = NewPowerNode() //NewPowerNode can be used because  PowerNode Implements Node Interface
	node.SetValue(5)
	fmt.Println("Node is of value ", node.GetValue())

	n = &SSLNode
	if n, ok := node.(*PowerNode); ok {
		fmt.Println("This is PowerNode", n.value)
	}
	var sllnode *SLLNode
	n := createNode(5)

	switch concreten := n.(type) {
	case *SSLNode:
		fmt.Println("Type is SSLNode, message", concreten.SNodeMessage)
	case *PowerNode:
		fmt.Println("type is PowerNode, message", concreten.PNodeMessage)
	}
	sNode := &SSLNode{value: 15}
	fmt.Println(sNode.GetValue())
}

func createNode(v int) Node {
	sn := NewSLLNode()
	sn.SetValue(v)
	return sn
}
