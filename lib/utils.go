package lib

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

//ReadFromFile this will read the content of file if -proxyfile is provided.
func readFromFile(filePath string) []string {

	var line []string
	file, err := os.Open(filePath)
	if err != nil {
		Printerr(err, fmt.Sprintf("utils:ReadFromFile: filePath: %s", filePath))
		os.Exit(1)
	} else {
		defer file.Close()
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line = append(line, scanner.Text())
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

// Execute a shell command.
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

//RandomLine give you a randomly line from file
func RandomLine(f string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n, _ := lineCounter(f)
	n = r.Intn(n)
	file, err := os.Open(f)
	defer file.Close()
	if err != nil {
		fmt.Printf("utils:RandomLine:file = %s, error %v", f, err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	i, line := 1, ""
	for scanner.Scan() {
		if n == i {
			line = scanner.Text()
			break
		}
		i++
	}
	return line
}

func lineCounter(f string) (int, error) {
	file, err := os.Open(f)
	defer file.Close()
	if err != nil {
		fmt.Printf("utils:RandomLine:file = %s, error %v", f, err)
	}
	b, count := make([]byte, 32*1024), 0
	for {
		c, err := file.Read(b)
		count += bytes.Count(b[:c], []byte{'\n'})
		switch {
		case err == io.EOF:
			return count, nil
		case err != nil:
			return count, err
		}
	}
}

//GetListFile get the list of a file in a directory.
func GetListFile(dir string) []string {

	files, err := ioutil.ReadDir(dir)
	Printerr(err, fmt.Sprintf("utils:GetListFile: d:%s", dir))
	var listfiles []string
	for _, f := range files {
		listfiles = append(listfiles, f.Name())
	}
	return listfiles

}

//CountNumberFileinFolder <- WTF
func CountNumberFileinFolder(dir string) int {

	count := 0
	if Existe(dir) {
		err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
			if strings.HasSuffix(path, ".json") {
				count++
			}
			return nil
		})
		Printerr(err, fmt.Sprintf("utils:CountNumberFileinFolder: d:%s", dir))
	}
	return count
}

//Existe check if a folder or file existe
func Existe(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
