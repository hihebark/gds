package core

import "fmt"

const (
	RESET       = "\033[0m" //all the const are defined for what thay do!
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

//Say: will output a message with the defined color
func Say(color, message string) string {
	return fmt.Sprintf(START, color) + message + RESET
}

//Info: to show output with orange color
func Info(message string) string {
	return fmt.Sprintf(START, ORANGE) + INFO + message + RESET
}
//Que: to show output with blue color
func Que(message string) string {
	return fmt.Sprintf(START, LIGHTBLUE) + QUE + message + RESET
}
//Bad: to show output with red color
func Bad(message string) string {
	return fmt.Sprintf(START, LIGHTRED) + BAD + message + RESET
}
//Info: to show output with green color
func Good(message string) string {
	return fmt.Sprintf(START, LIGHTGREEN) + GOOD + message + RESET
}
//Run: to show output with white color
func Run(message string) string {
	return fmt.Sprintf(START, WHITE) + RUN + message + RESET
}
