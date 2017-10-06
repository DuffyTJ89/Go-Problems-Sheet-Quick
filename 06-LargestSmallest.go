//Author : Thomas Duffy
//Date : 25/9/17

//Adapted from : https://play.golang.org/p/vhHmjhOMEo

package main

import "fmt"

func main() { //function to create a list of numbers
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	smallest, biggest := x[0], x[0] //set largest and smallest to 0
	for _, v := range x { //loop through the list
		if v > biggest { //if the number is bigger than the last sorted largest overwrite it 
			biggest = v
		}
		if v < smallest { //if the number is smaller than the last sorted smallest overwrite it 
			smallest = v
		}
	}
	//print out largest and smallest 
	fmt.Println("List of numbers used : ", x)
	fmt.Println("The biggest number is ", biggest)
	fmt.Println("The smallest number is ", smallest)
}