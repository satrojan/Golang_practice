package main

import (
    "fmt"
    "strings"
    "bufio"
    "os"
)

func main() {
    in := bufio.NewReader(os.Stdin) //One Example contained spaces. bufio allows us to get input with spaces
    var str_value string
    fmt.Printf("please enter a string:")
    str_value, err := in.ReadString('\n') //Get the value
    _=err //handle the err return. this isnt used so I assign it to _ to remove the error on run.
    str_lower := strings.ToLower(str_value) //set the string to lowercase as we no not care about case. Uppercase would work too.
    str_split := strings.Split(str_lower,"") //split the string into an array containing each character.
    /*
    str_split[0]=="i" tests if the first character is "i"
    str_split[len(str_split)-2]=="n" tests if the last character is 'n'.
        we use len(str_split)-2 to get the second to last character as the last character is a carrage return
    strings.Contains(str_lower,"a") tests if there is an "a" in the string str_lower
    */
    if str_split[0]=="i" && str_split[len(str_split)-2]=="n" && strings.Contains(str_lower,"a"){
        //I used \r\n to do a carrage return in linux. This is different in other OS's
        fmt.Printf("Found!\r\n")
    } else {
        fmt.Printf("Not Found!\r\n")
    }


}
