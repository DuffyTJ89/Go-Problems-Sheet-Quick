//Author : Thomas Duffy
//Date : 21/9/17

//Adapted from : https://gist.github.com/4E71/3735193
package main

import ( 
	"fmt" 
)

func main() {

	fmt.Print("Enter integer: ") //ask user to enter a whole number
	var input int //create a variable called input 
	fmt.Scanf("%d", &input) //read the user input and assign it to input variable 
	
	for i := 1; i <= input; i++ { //for loop to loop around to the user submitted end number and call the fizzbuzz function
		fizzbuzz(i)	
	}
}//end main

func fizzbuzz(i int) { //fnction where the calculations will be done
	fizz := "fizz" // assign variables 
	buzz := "buzz"
	
	if i % 3 == 0 && i % 5 == 0 { //anytime a number is divisible by both 3 and 5 print this
		fmt.Println(i, fizz + buzz)
	} else if i % 3 == 0 { // anytime a number is only divisible by 3 print this
		fmt.Println(i, fizz)
	} else if i % 5 == 0 { // anytime a number is only divisible by 5 print this
		fmt.Println(i, buzz)
	} else { // if the above conditions don't match print this
		fmt.Println(i)
	}
} //end fizzbuzz