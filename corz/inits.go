package corz

import "fmt"

//Always come before init
var INIT = tierUp()

func init() {
	//Always come before main
	fmt.Println("Initialize cor here!")
}

func tierUp() int {
	fmt.Println("Tier up!")
	return 5
}
