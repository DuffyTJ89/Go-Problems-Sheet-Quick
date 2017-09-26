//Author : Thomas Duffy
//Date : 26/9/17

//Adapted from :
package main


import (
	"fmt"
	"math/big" //math/big used for big ints
)



func factorial(n int64) *big.Int { //get factorial of a number. can handle big numbers using int64
	if n < 0 {
		return big.NewInt(1) 
	}
	if n == 0 {
		return big.NewInt(1)
	}
	bigN := big.NewInt(n)
	return bigN.Mul(bigN, factorial(n-1))

}

//main calls factorial() and sums and prints the sum of digits
func main() {
	fmt.Println(factorial(100)) //print out the factorial of 100
	sum := big.NewInt(0) //set up sum as a big int 
	n := new (big.Int)//to store the factorial
	fact := (factorial(100)) 
	for fact.BitLen() > 0 {
		_, n := fact.DivMod(fact, new(big.Int).SetUint64(uint64(10)), n)
		sum = sum.Add(sum,n)
	} 
	fmt.Println(sum)
}