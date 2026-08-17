package main

import "fmt"

// for -> only construct in GO for looping
func main(){

	// while loop
	i:=1
	for i<= 3 {
		fmt.Println(i)
		i=i+1;
	}

	// infinite loop 
	// for{
	// 	fmt.Println("1")
	// }

	//claassic for loop
	for i:=0 ; i<=4;i++ {
		// break
		// continue
		fmt.Println(i)
	}

	// range
	for i:= range 3 {
		fmt.Println(i)
	}
}
