
package main

import (
	"fmt"
)

// func main(){

// 	//ways of declaring variables in go
// 	//---1st method
// 	//var i int
// 	//i = 100
// 	//-------------
// 	// var i int ---2nd method
// 	// i = 100
// 	//------------the first to method are valuable since they give one control in deciding the type and use of each variable

// 	// ---3rd method
// 	i := "Kenya"
// 	fmt.Println(i)
// }

//----declaring variables outside the function-----//
// var i int = 100
// var k float32 = 100
// func main(){
// 	fmt.Printf("%v, %T", k, k)
// }


// //-----------variable block(s)-----------//

// var (
// 	actorName string = "Elizabeth Sladen"
// 	companion string = "Sarah Smith"
// 	doctorNumber int = 3
// 	season int = 11
// )

// //--or--//

// var (
// 	counter int = 0
// )

// //then function
// func main() {
// 	//code goes here
// }

//---how variables work when you try to redeclare them
// var i int = 100
// func main() {
// 	var i int = 101
// 	// i := 103 //this brings an error since there is no new variable
// 	i = 105 //you can reassign the value to a variable but can not redeclare it in the same scope
// 	fmt.Println(i)
// } 

// //---Shadowing-------//
// //-->this means that the variable in the innermost scope takes precedence 
// var i int = 106
// func main() {
// 	var i int = 105 //when your run this is the variable that will be compiled -->shadowing
// 	fmt.Println(i)
// }

// //----Declared Variables in Go have to be used----//
// func main() {
// 	var i int = 105
// 	j :="Kenya" //this brings are runtime error since j is declared and not used
// 	fmt.Println(i)
// }

