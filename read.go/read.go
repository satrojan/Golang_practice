package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    _"encoding/json"
)

type names struct{
    fname string
    lname string
}

func main() {

    //declare variables
    var f_name string
    var s_names[] names

    //get values
	fmt.Println("Please enter file name.")
    fmt.Scanln(&f_name)

    //open file and read each line
    file, _ := os.Open(f_name)
    scn := bufio.NewScanner(file)
    for scn.Scan() {
        line :=scn.Text()
        //split into first and last name based on whitespace
        splt_name:=strings.Fields(line)
        //appends struct into slice
        s_names = append(s_names, names{fname:splt_name[0],lname:splt_name[1]})

    }
    //print out names in file
    for _, name := range s_names {
        fmt.Println(name)
    }

}
