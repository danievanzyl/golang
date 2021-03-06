//Main

package main

import (
    "fmt"
    "time"
    "strconv"
)

func pinger(c chan string) {
    for i := 0; ; i++ {
        c <- "ping" + strconv.Itoa(i)
    }
}
func printer(c chan string) {
    for {
        msg := <- c
        fmt.Println(time.Now())
        fmt.Println(msg)
        time.Sleep(time.Second * 1)
    }
}
func main() {
    var c chan string = make(chan string)
    
    go pinger(c)
    go printer(c)
    
    var input string
    fmt.Scanln(&input)
}
