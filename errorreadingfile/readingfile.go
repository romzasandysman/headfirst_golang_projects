package main

import (
	"bufio"
	"os"
	"strconv"
)

//GetFloats read value float64 from each row of file
func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := os.Open(fileName)
	if err != nil{
		return nil, err
	}

	var number float64

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		number, err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil{
			return nil, err
		}
		numbers = append(numbers, number)
	}

	err = file.Close()
	if err != nil{
		return numbers, err
	}
	if scanner.Err() != nil{
		return nil, err
	}

	return numbers, nil
}
