package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("couldn't convert number and it is not a number : %v\n", err)
		return
	}
	if (n % 2) == 0 {
		fmt.Printf("%v is even", os.Args[1])
	} else {
		fmt.Printf("%v is odd", os.Args[1])
	}

}
