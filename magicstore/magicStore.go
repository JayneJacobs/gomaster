package magicstore

import "fmt"

type magicStore struct {
	value interface{}
	name  string
}

func (ms *magicStore) SetValue(v interface{}) {
	ms.value = v
}

func (ms *magicStore) GetValue() interface{} {
	return ms.value
}

func NewMagicStore(nm string) *magicStore {
	return &magicStore{name: nm}
}

// MagicStore takes a string and an integer
func MagicStore(n string, v int) {
	mstore := NewMagicStore(n)
	mstore.SetValue(v)
	fmt.Println(mstore.GetValue())

	istore := NewMagicStore("Integer Store")
	istore.SetValue(v)
	if v, ok := istore.GetValue().(int); ok {
		v *= 4
		fmt.Println(v)

	}

	sstore := NewMagicStore("String Store")
	sstore.SetValue(n)
	if v, ok := sstore.GetValue().(string); ok {
		v += " World"
		fmt.Println(v)

	}
}
