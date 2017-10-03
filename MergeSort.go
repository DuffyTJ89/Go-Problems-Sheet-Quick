//Author : Thomas Duffy
//Date : 3/10/17

//Adapted from : https://rosettacode.org/wiki/Sort_an_integer_array#Go
//	
// [1,4,6],[2,3,5] → [1,2,3,4,5,6].

package main

import "fmt"
import "sort"

func appendSlices() {
	myIntSlice := []int{1, 4, 6} //first list of numbers
	myIntSlice = append(myIntSlice, []int{2, 3, 5}...) //add on the 2nd list 
	fmt.Println("Unsorted : ", myIntSlice) //print out the list unsorted
	sort.Ints(myIntSlice) //use sort to reorder the list 
	fmt.Println("Sorted : ", myIntSlice) //print sorted list
	

}//end appendSlices
func main() {
	appendSlices()
}