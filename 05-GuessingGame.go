//Author : Thomas Duffy
//Date : 25/9/17

//Adapted from : http://golangcookbook.blogspot.ie/2012/12/guess-number-game-v2.html

package main
import (
    "fmt"
    "math/rand"
    "time"
	//need the math random and time libraries 
)
	//function to set the random number off the current time so as to make it look random.
func xrand(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

func main() {
    myrand := xrand(1, 10) //get a random numnber between 1 and 10
    tries := 0 //variable to track the number of tries.
    var guess int //track the user guesses
	prevGuess := 0 //track the previous guess

    fmt.Println("Welcome to Guess My Number Game!")
    for guess != myrand {
        fmt.Println("Take a guess...")
        fmt.Scanf("%v", &guess)
        tries++
		
		if guess==prevGuess{
			fmt.Println("You already guessed that number!!")
			tries--;
		} else if guess > myrand {
            fmt.Println("Too high")
        } else if guess < myrand {
            fmt.Println("Too low")
        } else {
            fmt.Printf("Good job! You guessed it in %v tries", tries)
            break
        }
	prevGuess = guess; //assign previous guess the value for current guess
    }
}