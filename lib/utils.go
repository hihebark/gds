package lib

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

//CountLine Count the number of line in a wordlist to determine how much we will brutforce.
func CountLine(file string) string {
	count, err := Execute("/usr/bin/wc", []string{"-l", file})
	Printerr(err, "utils:Countline")
	return strings.Split(count, " ")[0]
}

//ReadFromFile this will read the content of file if -proxyfile is provided.
func ReadFromFile(filePath string) []string {

	var line []string
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err.Error() + `: ` + filePath)
		os.Exit(1)
	} else {
		defer file.Close()
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line = append(line, scanner.Text())
		//fmt.Println(scanner.Text())
	}
	return line

}

//ReturnStringFile return the content of file
func ReturnStringFile(filePath string) string {

	var line string
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err.Error() + `: ` + filePath)
		os.Exit(1)
	} else {
		defer file.Close()
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line += scanner.Text()
		//fmt.Println(scanner.Text())
	}
	return line

}

//WriteToFile to write to a file
func WriteToFile(filePath string, instring string) {

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println(filePath, err)
		return
	}
	defer file.Close()
	file.WriteString(instring + "\r\n")
}

//ReadLine file line per line.
func ReadLine(r *bufio.Reader) (string, error) {

	var (
		isPrefix = true
		err      error
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err

}

// Execute a shell command.
// return cmd (output) or emty string if error.
func Execute(pathExec string, args []string) (string, error) {

	path, err := exec.LookPath(pathExec)
	if err != nil {
		return "", err
	}
	cmd, err := exec.Command(path, args...).CombinedOutput()
	if err != nil {
		return string(cmd), err
	}
	return string(cmd), nil

}

//GetRandLine to return random line of file
func GetRandLine(file string) string {
	line, err := Execute("/usr/bin/shuf", []string{"-n 1", file})
	Printerr(err, fmt.Sprintf("utils:GetRandLine: f:%s", file))
	return line
}
