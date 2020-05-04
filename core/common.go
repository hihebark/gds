package core

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func readFile(path string, lines chan string) {
	file, err := os.Open(path)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines <- strings.TrimRight(scanner.Text(), " ")
	}
	close(lines)
}

//WriteToFile to write to a file
func WriteToFile(filePath string, content string) error {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(content + "\r\n")
	return nil
}

func byteConverter(length int64) string {
	mbyte := []string{"bytes", "KB", "MB", "GB", "TB"}
	if length == -1 {
		return "0 byte"
	}
	for _, x := range mbyte {
		if length < 1024.0 {
			return fmt.Sprintf("%3.1d %s", length, x)
		}
		length = length / 1024.0
	}
	return "Error"
}

// RandomInt function
func RandomInt(count int) int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(count)
}
