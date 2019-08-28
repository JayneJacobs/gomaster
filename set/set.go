package set

import "fmt"

type Set map[string]struct{}

//Set Adds items to the Set
func Setadd() {
	s := make(Set)
	//add Items
	s["Item1"] = struct{}{}
	s["Item2"] = struct{}{}
	fmt.Println(GetSetValues(s))

}

func GetSetValues(s Set) []string {
	var retVal []string
	for k, _ := range s {
		retVal = append(retVal, k)
	}
	return retVal
}
