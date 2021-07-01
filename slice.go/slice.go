package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	num_slice := make([]int, 3) //declare slice of size 3

	//set up bufio scanner
	scn := bufio.NewScanner(os.Stdin)
	str := "Input a number, enter x to quit:"
	i := 0

	for fmt.Println(str); scn.Scan(); fmt.Println(str) {
		line := scn.Text()
		//handle exit case first
		if line == "x" {
			break
		}

		//convert string into number
		num_, err := strconv.Atoi(line)

		//error handeling
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}
		//Add first three values to the slice
		if i < 3 {
			num_slice[i] = num_
			i++
			fmt.Println(i)
			//Add additional values to the slice with append
		} else {
			num_slice = append(num_slice, num_)
			fmt.Println("ok")
		}
		disp_slice := make([]int, len(num_slice)) //create temp slice for display
		copy(disp_slice, num_slice)               //copy slice

		sort.Ints(disp_slice)   //sort slice
		fmt.Println(disp_slice) //display slice

	}
}
