package customrand

import (
	"fmt"
	"math/rand"
	"time"
)

type customRand struct {
	*rand.Rand
	count int
}

//NewCustomRand takes an integer and returns a pointer to a Random Number

func NewCustomRand(i int64) *customRand {
	return &customRand{
		Rand:  rand.New(rand.NewSource(i)),
		count: 0,
	}

}

func (cr *customRand) GetCount() int {
	return cr.count
}

func (cr *customRand) RandRange(min, max int) int {
	cr.count++
	return cr.Rand.Intn(max-min) + min
}

func (cr *customRand) Intn(n int) int {
	fmt.Println("OuterIntn Called...")
	cr.count++
	return cr.Rand.Intn(n) + 1
}

//  CustomRandEx shows how to use time.Now to create a random number
func CustomRandEx() {
	cr := NewCustomRand(time.Now().UnixNano())
	fmt.Println(cr.RandRange(5, 30))
	fmt.Println(cr.Intn(10))
	fmt.Println(cr.GetCount())

}
