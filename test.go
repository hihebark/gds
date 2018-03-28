package main

import (
    "fmt"
    "github.com/hihebark/godirsearch/core"
)

func main(){
    fmt.Println(core.Say(core.Grey, "Test"))
    fmt.Println(core.Info("info"))
    fmt.Println(core.Que("Que"))
    fmt.Println(core.Bad("Bad"))
    fmt.Println(core.Good("Good"))
    fmt.Println(core.Run("Run"))
    
}
