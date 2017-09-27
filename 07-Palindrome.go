//Author : Thomas Duffy
//Date : 27/9/17

//Adapted from : https://play.golang.org/p/5FUOzjBa-o

package main

import "fmt"

func isPalindrome(input string) bool { 
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}

func main() {

	var userInput
	
	fmt.Println("Palindrome check : ")
    fmt.Scanf("%s", userInput)


	fmt.Println(isPalindrome(userInput))
	//fmt.Println(isPalindrome("bob"))
}



