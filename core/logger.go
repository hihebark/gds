package core

import "fmt"

var (
    Green, Lightgreen, Grey, Black, Red, Lightred, Cyan, Lightcyan, Blue,
    Lightblue, Purple, Yellow, White, Lightpurple, Orange, reset, start string
    
    info, que, bad, good, run string
    
    bg, bold, italic, under, strike string
)

func init(){
    reset           = "\033[0m"
    start           = "\033[%sm"
    Orange          = "33"
    Green           = "32"
    Lightgreen      = "92"
    Grey            = "37"
    Black           = "30"
    Red             = "31"
    Lightred        = "91"
    Cyan            = "36"
    Lightcyan       = "96"
    Blue            = "34"
    Lightblue       = "94"
    Purple          = "35"
    Yellow          = "93"
    White           = "97"
    Lightpurple     = "95"
    
    info    = "[!] "
    que     = "[?] "
    bad     = "[-] "
    good    = "[+] "
    run     = "[~] "

    bg      = ";7"
    bold    = ";1"
    italic  = "3"
    under   = "4"
    strike  = "09"
}

func Say(color, message string) string{
    return fmt.Sprintf(start, color)+message+reset
}

func Info(message string) string{
    return fmt.Sprintf(start, Orange)+info+message+reset
}

func Que(message string) string{
    return fmt.Sprintf(start, Lightblue)+que+message+reset
}

func Bad(message string) string{
    return fmt.Sprintf(start, Lightred)+bad+message+reset
}

func Good(message string) string{
    return fmt.Sprintf(start, Lightgreen)+good+message+reset
}

func Run(message string) string{
    return fmt.Sprintf(start, White)+run+message+reset
}

