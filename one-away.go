// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func isOneEdit(s1 string, s2 string) bool {
	if len(s1) == len(s2) {
		return oneEditReplace(s1, s2)
	} else if (len(s1) + 1) == len(s2) {
		return oneEditSubtractOrAdd(s1, s2)
	} else if (len(s1) - 1) == len(s2) {
		return oneEditSubtractOrAdd(s2, s1)
	} else {
		return false
	}
}

func oneEditSubtractOrAdd(s1 string, s2 string) bool {
	oneEdit := 0
	for i, _ := range s1 {
		if oneEdit > 1 {
			return false
		}
		if s1[i] != s2[i] {
			oneEdit += 1
		}
	}
	return true
}

func oneEditReplace(s1 string, s2 string) bool {
	oneEdit := 0
	for i, _ := range s1 {
		if oneEdit > 1 {
			return false
		}
		if s1[i] != s2[i] {
			oneEdit += 1
		}
	}
	return true
}

func main() {
	fmt.Println(isOneEdit("pale", "ple"))
	fmt.Println(isOneEdit("pales", "pale"))
	fmt.Println(isOneEdit("pale", "bale"))
	fmt.Println(isOneEdit("pale", "bake"))
}
