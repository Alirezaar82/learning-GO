package main

import "fmt"
import "strings"

func main() {
    var hope int
    var end_game int
    number := 1
    fmt.Scanf("%d",&hope)
    fmt.Scanf("%d",&end_game)
    
    for i := 1 ; i <= end_game; i++ {
        if i % hope == 0 {
            result := strings.Repeat("Hope ", number)
            fmt.Println(result)
            number++
        }else{
            fmt.Println(i)
        }
    }

}