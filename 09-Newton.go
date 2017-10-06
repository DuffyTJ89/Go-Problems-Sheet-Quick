//Author : Thomas Duffy
//Date : 6/10/17

//Adapted from : In Class example
//

//Program goes into an infinite loop when the user enters numbers over 37 for root and any number for a guess. 
//Not sure what the problem is as it can handle 100 and 90 and a few more bigger numbers fine, so it isnt a size problem?

package main

import (

	"fmt"
	"math"
)

func z_next(z float64, x float64) float64 { //z float64, x float64...returns float64

	return z - ((z*z - x) / (2 * z))//Newton's formula 
}

func main() {

	//the number to find square root of
	//x := 20.0
	var x float64

	//the initial guess
	var z float64
	var startingGuess float64 = 1.0

	fmt.Print("Please enter the number to find square root of: ") //user input
	fmt.Scanf("%f", &x)

	//loop over until the next value is the same as your current guess.
	//When they are the same the problem is solved
	
	fmt.Print("Enter your guess : ")//user input for starting guess
	fmt.Scanf("%f", &startingGuess)

	for z = startingGuess; z != z_next(z, x); z = z_next(z, x) { //start at that guess and loop through until the two values are the same
		//print the guess on each loop
		fmt.Printf("Current guess: %13f\n", z)

	}
	//finally, z is a good appriximation of the square root of x
	fmt.Printf("The square root of %f is %f\n", x, z)

	//print out the math.Sqrt value
	fmt.Printf("math.Sqrt gives the value as %f\n", math.Sqrt(x))

}