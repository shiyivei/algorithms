package string

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	s := "lrloseumgh"
	a := reverseLeftWords(s,6)
	fmt.Println(a)

	//s := "we are happy"
	//
	//fmt.Println(replaceSpace(s))

	//s := "lrloseumgh"

	//fmt.Println(reverseLeftWords(s, 6))

	//words := []string{"a", "b", "a"}
	//fmt.Println(removeAnagrams(words))

	//a := "a"
	//b := "b"
	//special := []int{12, 24, 38, 48}

	//fmt.Println(maxConsecutive(1, 50, special))


	fmt.Println(firstUniqChar(s))

	fmt.Println(ReplaceSpace(s))

	fmt.Println(ReverseLeftWords(s,3))
}
