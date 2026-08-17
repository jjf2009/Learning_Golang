package main

import "fmt"

// const age = 30
// name :="Golang" // not allowed

func main(){
	// :=
	const name = "Golang"

	// name = "Java"

	// fmt.Println(name)

	const (
		port = 3000
		host = "localhost"
	)
    fmt.Println(port,host)

}
