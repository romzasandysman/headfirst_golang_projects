package main

import (
	"keyboard"
	"fmt"
)

func main() {
	float, err := keyboard.GetFloat()
	if(err != nil){
		fmt.Println(err)
	}

	fmt.Println(float)
}
