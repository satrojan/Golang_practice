package main

import "fmt"

func main() {
    var float_value float32

	fmt.Println("Please enter decimal.")
    fmt.Scan(&float_value)
    var int_value = int(float_value)
    fmt.Println(int_value)
}
