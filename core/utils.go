package core

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

//CountLine Count the number of line in a wordlist to determine how much we will brutforce.
func CountLine(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

//ReadFromFile this will read the content of file if -proxyfile is provided.
func ReadFromFile(filePath string) {

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
		fmt.Println(scanner.Text())
	}

}

//Readln file line per line.
func Readln(r *bufio.Reader) (string, error) {

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

//Execute a commande basically used to excute grepproxylist.sh .
func Execute(pathExec string, args []string) (string, error) {

	path, err := exec.LookPath(pathExec)
	if err != nil {
		return "", err
	}
	cmd, err := exec.Command(path, args...).CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(cmd), nil

}
