//Package keyboard reads user imput from the keyboard
package keyboard

import (
	"bufio"
	"strings"
	"strconv"
	"os"
)

//GetFloat reads a floating-point number from the keyboard.
//It return the number read and any error encountered
func GetFloat() (float64, error){
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil{
		return 0, err
	}

	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)

	if err != nil{
		return 0, err
	}

	return number, nil
}
