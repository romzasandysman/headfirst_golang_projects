package main

import (
	"fmt"
	"log"
	"os"
)

func check(err error)  {
	if err != nil{
		log.Fatal(err)
	}
}

func main() {
	//file, err := os.OpenFile("aardvark.txt", os.O_RDONLY, os.FileMode(0600))
	file, err := os.OpenFile("aardvark.txt", os.O_WRONLY | os.O_APPEND | os.O_CREATE, os.FileMode(0600))
	check(err)
	_, err = file.Write([]byte("amazing\n"))
	check(err)

	//scanner := bufio.NewScanner(file)
	//for scanner.Scan(){
	//	fmt.Println(scanner.Text())
	//}

	err = file.Close()
	check(err)
	//check(scanner.Err())

	fmt.Printf("%08b\n", 170 & 15)
	fmt.Println(os.O_WRONLY)
}
