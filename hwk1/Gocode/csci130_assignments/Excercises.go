//1 & 2
package main

import "fmt"

func main(){

	fmt.Println("Hello World")
	fmt.Println("Hello my name is William")
}

//3
func main(){
        var name string

	fmt.Println("enter name:")

	fmt.Scan(&name)

	fmt.Println("Hello", name)

}


//4
package main

import "fmt"

func main(){
        var number_big int
	var number_small int

	fmt.Println("Enter a large number:") //asks for large number
	fmt.Scan(&number_big)                //scans large number in

	fmt.Println("Enter a small number:")
	fmt.Scan(&number_small)

	var solution int

	solution = (number_big/number_small)  //division operation here and sets value to solution
	fmt.Println("Solution is:",solution)  //prints value in solution

}
