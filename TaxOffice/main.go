package main

import "fmt"

func main() {
    var tax int
    var income int
    var newIncom int
    fmt.Scanf("%d",&income)

    newIncom = income - 100
    if newIncom <= 0{
        tax += (income * 5) / 100
        fmt.Println(tax)
        return
    }else{
        tax += 5
        income -= 100
    }
    newIncom = 0

    newIncom = income - 400
    if newIncom <= 0{
        tax += (income * 10) / 100
        fmt.Println(tax)
        return 
    }else{
        tax += 40
        income -= 400
    }
    
    newIncom = 0
    newIncom = income - 500
    if newIncom <= 0{
        tax += (income * 15) / 100
        fmt.Println(tax)
        return
    }else{
        tax += 75
        income -= 500
    }
    
    tax += (income * 20) / 100
    
    fmt.Println(tax)
}