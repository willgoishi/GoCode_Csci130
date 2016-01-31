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
//5
func main(){

	for x:=0; x<=100; x++{
            if x%2 == 0 {
		fmt.Println(x)
	    }

           }


}

//6

func main(){

	for x:=0; x<=100; x++{
            if x%3 == 0 && x%5 == 0 {
		fmt.Println("FizzBuzz")
	    }else if x%3 == 0 {
                fmt.Println("Fizz")
            }else if x%5 == 0 {
                fmt.Println("Buzz")
            }else {
                fmt.Println(x)
            }

           }


}
//7
func main(){

	var sum int
	for x:=0; x<1000; x++{
            if x%3 == 0 || x%5 == 0 {
		sum  = sum + x
	    }
           }
        fmt.Println(sum)

}

//Part 2, 1
func half(number int)(int ,bool){
	return number/2, number%2 == 0
}
func main(){
        var number int
	fmt.Println("Enter a number:")
	fmt.Scan(&number)
	fmt.Println(half(number))

}
//Part 2, 2
func main(){

        half := func(number int)(int, bool) {
		return number / 2, number % 2 == 0
	}
	fmt.Println(half(45))

}
//part2, 3
func maximum(numbers ... int) int{
    var largest_int int
	for _,variables := range numbers{
		if variables> largest_int{
			largest_int = variables
		}
	}
	return largest_int
}

func main(){

	fmt.Println( maximum(24,324,452,34534,626,54))

}

//part 2,4    answer = true
func main(){


	fmt.Println( (true && false) || (false && true) || !(false && false))

}
//part 2, 5
func foo(numbers ...int) {
	fmt.Println(numbers)
}
func main() {
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
}
//part 2, 6

/*Problem was to take a number and create a revolving pattern making a square format of numbers. Then add up the diagonals of numbers that make up the cross between those numbers.
*/
func main() {
	var count int
	var sum_of_diagonals int
	count = 1001 * 1001

	var spacer int
        spacer = 1                              //init spacer to odd 1
		//traverses 4 numbers so that you get the diagonal values (4 of them)
		for i := 1; i < count; i ++ {         //traverses through the count size
                        for k := 1; k<spacer; k = i {
				for j := 0; j < 4; j++ {
					if i % 2 != 0 {
						//checks if the value is odd
						sum_of_diagonals = sum_of_diagonals + i     //adds those values to the diagonal sum
						fmt.Println(sum_of_diagonals)
					}
				}

			}
			spacer = spacer + 2//increments spacer so that each side is incremented by 2 (odd incrementation)
			                   //skip 1 then 3 then 5 numbers

		}
                spacer = spacer + 2//increments spacer so that each side is incremented by 2 (odd incrementation)
                           //skip 1 then 3 then 5 numbers

	fmt.Println(sum_of_diagonals)
}
