package lib

import "fmt"

//All the const are defined for what they do!
const (
	RESET       = "\033[0m"
	START       = "\033[%sm"
	ORANGE      = "33"
	GREEN       = "32"
	LIGHTGREEN  = "92"
	GREY        = "37"
	BLACK       = "30"
	RED         = "31"
	LIGHTRED    = "91"
	CYAN        = "36"
	LIGHTCYAN   = "96"
	BLUE        = "34"
	LIGHTBLUE   = "94"
	PURPLE      = "35"
	YELLOW      = "93"
	WHITE       = "97"
	LIGHTPURPLE = "95"

	INFO = "[!] "
	QUE  = "[?] "
	BAD  = "[-] "
	GOOD = "[+] "
	RUN  = "[~] "

	BG     = ";7"
	BOLD   = ";1"
	ITALIC = "3"
	UNDER  = "4"
	STRIKE = "09"
)

//Say is a function that will output a message with the defined color.
func Say(color, message string) {
	fmt.Printf("%s%s%s\n", fmt.Sprintf(START, color), message, RESET)
}

//SayMe is a function that will return a message with the defined color.
func SayMe(color, message string) string {
	return fmt.Sprintf(START, color) + message + RESET
}

//Info is a function that will show output with orange color.
func Info(message string) {
	fmt.Printf("%s%s%s%s\n", fmt.Sprintf(START, YELLOW), INFO, message, RESET)
}

//Que is a function that will show output with blue color.
func Que(message string) {
	fmt.Printf("%s%s%s%s\n", fmt.Sprintf(START, LIGHTBLUE), QUE, message, RESET)
}

//Bad is a function that will show output with red color.
func Bad(message string) {
	fmt.Printf("%s%s%s%s\n", fmt.Sprintf(START, LIGHTRED), BAD, message, RESET)
}

//Good is a function that will show output with green color.
func Good(message string) {
	fmt.Printf("%s%s%s%s\n", fmt.Sprintf(START, LIGHTGREEN), GOOD, message, RESET)
}

//Run is a function that will show output with white color.
func Run(message string) {
	fmt.Printf("%s%s%s%s\n", fmt.Sprintf(START, WHITE), RUN, message, RESET)
}

//Printerr is a function that will print error message if found
func Printerr(err error, fromwhere string) {
	if err != nil {
		Bad(fmt.Sprintf("%s : %v", fromwhere, err))
	}
}
