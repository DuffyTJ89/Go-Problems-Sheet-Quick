//Author : Thomas Duffy
//Date : 4/10/17

//Adapted from : https://play.golang.org/p/Vq-TIHJCHE
//what is rune : https://stackoverflow.com/questions/19310700/what-is-a-rune				

package main

import "fmt"

func Reverse(s string) string {
	r := []rune(s) //rune maps character/string to their unicode codepoint.
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 { //loop over and swap characters around
	r[i], r[j] = r[j], r[i]
	}
	return string(r) //return the string reversed
}

func main() { //some test words to reverse 
	fmt.Println(Reverse("Hello, world"))
	fmt.Println(Reverse("test"))
	fmt.Println(Reverse("Thomas Duffy"))
}

