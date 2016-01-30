
package main


import "fmt"

func main(){
	var int_var = 500
	var str_var = "hello world"

	id_number := 123453

	boolean_var,i,j,k := false,3,4,5
	fmt.Println(int_var, str_var)

        fmt.Print(id_number)
	fmt.Println(boolean_var,i,j,k)

	//printing types
	fmt.Printf("%T \n",id_number) //takes in %T which represents type

	fmt.Printf("%T \n",boolean_var)
        fmt.Printf("%T","hello")
}
