package main

import (
    "fmt"
    "encoding/json"
)

func main() {

    //declare variables
    var h_name, h_addr string
    m := make(map[string]string)

    //get values
	fmt.Println("Please enter your name.")
    fmt.Scanln(&h_name)
    fmt.Println("Please enter your address.")
    fmt.Scanln(&h_addr)

    //put values into map
    m["name"]=h_name
    m["address"]=h_addr

    //convert map to json
    j,_ := json.Marshal(m)
    fmt.Println(string(j))

}
