//Author : Thomas Duffy
//Date : 27/9/17

//Adapted from : https://play.golang.org/p/5FUOzjBa-o
//				 

package main

import "fmt"

func isPalindrome(input string) bool { //function that takes a user inout which is passed from the main 
	for i := 0; i < len(input)/2; i++ { //loop over the individual characters in the word for the length of the word. Divide it in 2 for check
		if input[i] != input[len(input)-i-1] { //if there are any differences between the characters when the word is split in 2 then return false
			return false
		}
	}
	return true //otherwise the test returns true
}

func main() {

	var userInput string
	
		fmt.Println("Palindrome check : ")
		fmt.Scanf("%s ", &userInput)  //user input from console 
		
		fmt.Println(isPalindrome(userInput)) //call the function where the checking takes place and pss it the word  the user entered.
	
}//end main



