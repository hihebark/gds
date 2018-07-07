package lib

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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

//GetRandLine to return random line of file
func GetRandLine(file string) string {
	line, err := Execute("/usr/bin/shuf", []string{"-n 1", file})
	Printerr(err, fmt.Sprintf("utils:GetRandLine: f:%s", file))
	return line
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
