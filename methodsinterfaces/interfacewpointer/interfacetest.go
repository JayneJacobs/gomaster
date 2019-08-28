package methodsinterfaces

import "fmt"

//Node interface declares methods
type Node interface {
	SetValue(v int)
	GetValue() int
}

// SLLNode Implements the Node Interface
type SLLNode struct {
	next  *SLLNode
	value int
}

// SetValue *SSLNode Method
func (sNode *SLLNode) SetValue(v int) {
	sNode.value = v
}

// GetValue *SSLNode Method
func (sNode *SLLNode) GetValue() int {
	return sNode.value
}

// NewSLLNode creates an instance of SSLNode and returns a pointer
func NewSLLNode() *SLLNode {
	return new(SLLNode) //new keyword allocates memory and returns a reference to the data .
}

// PowerNode Type
type PowerNode struct {
	next  *PowerNode
	value int
}

// SetValue method
func (sNode *PowerNode) SetValue(v int) {
	sNode.value = v * 10
}

// GetValue returns sNode as an int
func (sNode *PowerNode) GetValue() int {
	return sNode.value
}

// NewPowerNode Constructor
func NewPowerNode() *PowerNode {
	return new(PowerNode)
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

	if n, ok := node.(*PowerNode); ok {
		fmt.Println("This is PowerNode", n.value)
	}

}
